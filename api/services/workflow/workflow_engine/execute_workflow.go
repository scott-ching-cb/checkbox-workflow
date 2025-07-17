package workflow_engine

import (
	"context"
	"log/slog"
	"os"
	"time"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ExecuteWorkflowParams struct {
	Ctx        context.Context
	DB         *pgx.Conn
	FormData   *workflow.WorkflowFormData
	StartTime  string
	Store      *store.Store
	WorkflowId uuid.UUID
}

type ComputeNextNodeParams struct {
	CurrentNode           *workflow.Node
	ExecutionStepResults  *workflow.ExecutionStep
	NodeIdToNodeMap       map[string]*workflow.Node
	SourceNodeIdToEdgeMap map[string][]*workflow.Edge
}

// ComputeNextNode returns the next node to be executed by the workflow engine
func ComputeNextNode(params *ComputeNextNodeParams) *workflow.Node {
	// If the current node is not a condition node, there is only one edge pointing to the next node
	currentNode := params.CurrentNode
	edges := params.SourceNodeIdToEdgeMap[currentNode.Id]
	if currentNode.Type != "condition" {
		return params.NodeIdToNodeMap[edges[0].Target]
	}

	// If the current node is of type condition and alert is triggered return email node, else return next node
	isAlertTriggered := params.ExecutionStepResults.Output.ConditionResult.Result
	var nextNode *workflow.Node
	for _, edge := range edges {
		isEdgeTargetEmailNode := params.NodeIdToNodeMap[edge.Target].Type == "email"
		if isEdgeTargetEmailNode && isAlertTriggered {
			nextNode = params.NodeIdToNodeMap[edge.Target]
			break
		} else if !isEdgeTargetEmailNode && !isAlertTriggered {
			nextNode = params.NodeIdToNodeMap[edge.Target]
			break
		}
	}
	return nextNode
}

// ExecuteWorkflow retrieves the persisted workflow from the database and executes the workflow
func ExecuteWorkflow(params *ExecuteWorkflowParams) (*workflow.ExecutionResults, error) {
	// Retrieve the persisted workflow from the database
	slog.Debug("workflow_engine : initiating workflow execution", "id", params.WorkflowId)
	databaseStore := *params.Store
	workflowResponse, err := databaseStore.GetWorkflow(&store.GetWorkflowParams{
		Ctx:        params.Ctx,
		DB:         params.DB,
		WorkflowId: params.WorkflowId,
	})
	if err != nil {
		return nil, err
	}

	// Create a map for faster node look-up of nodes from edge target
	var startNode *workflow.Node
	nodeIdToNodeMap := make(map[string]*workflow.Node)
	for _, node := range workflowResponse.Nodes {
		if node.Type == "start" {
			startNode = node
		}
		nodeIdToNodeMap[node.Id] = node
	}

	// Create a map for faster lookup from source node to edge
	sourceNodeIdToEdgeMap := make(map[string][]*workflow.Edge)
	for _, edge := range workflowResponse.Edges {
		if _, ok := nodeIdToNodeMap[edge.Source]; ok {
			sourceNodeIdToEdgeMap[edge.Source] = append(sourceNodeIdToEdgeMap[edge.Source], edge)
		} else {
			sourceNodeIdToEdgeMap[edge.Source] = []*workflow.Edge{edge}
		}
	}

	// Populate execution result metadata (environment and triggered by)
	environment := os.Getenv("ENVIRONMENT")
	executionResults := &workflow.ExecutionResults{
		ExecutionId: uuid.New().String(),
		Metadata: &workflow.ExecutionResults_ExecutionMetadata{
			Environment: &environment,
			TriggeredBy: &params.FormData.Name,
		},
		StartTime:     params.StartTime,
		Steps:         []*workflow.ExecutionStep{},
		TotalDuration: wrapperspb.Int64(time.Duration(0).Milliseconds()),
	}

	// Execute workflow nodes in order (beginning from "start" type node)
	currentNode := startNode
	stepNumber := int32(1)
	var previousExecutionStepResults *workflow.ExecutionStep
	outputVariables := make(map[string]string)
	for currentNode != nil {
		executionStepResults, stepOutputVariables := ExecuteNode(&ExecuteNodeParams{
			FormData:                    params.FormData,
			Node:                        currentNode,
			Nodes:                       workflowResponse.Nodes,
			PreviousExecutionStepResult: previousExecutionStepResults,
			StepNumber:                  stepNumber,
			StoredOutputVariables:       outputVariables,
		})
		executionResults.Steps = append(executionResults.Steps, executionStepResults)

		// Update the output variables
		if stepOutputVariables != nil {
			for key, value := range stepOutputVariables {
				outputVariables[key] = value
			}
		}

		// Update the execution result's total duration
		currentTotalDuration := executionResults.TotalDuration.Value
		newTotalDuration := currentTotalDuration + executionStepResults.Duration.Value
		executionResults.TotalDuration = wrapperspb.Int64(newTotalDuration)

		// Get the next node in the graph to execute, or break if current node failed or current node is the end node
		if executionStepResults.Status != "completed" || currentNode.Type == "end" {
			break
		}
		currentNode = ComputeNextNode(&ComputeNextNodeParams{
			CurrentNode:           currentNode,
			ExecutionStepResults:  executionStepResults,
			SourceNodeIdToEdgeMap: sourceNodeIdToEdgeMap,
			NodeIdToNodeMap:       nodeIdToNodeMap,
		})
		previousExecutionStepResults = executionStepResults
		stepNumber++
	}

	// Set the execution result's status and end time
	if executionResults.Steps[len(executionResults.Steps)-1].Status != "completed" {
		executionResults.Status = "failed"
	} else {
		executionResults.Status = "completed"
	}
	executionResults.EndTime = time.Now().Format(time.RFC3339)
	return executionResults, nil
}
