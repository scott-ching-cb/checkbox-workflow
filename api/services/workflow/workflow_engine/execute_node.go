package workflow_engine

import (
	"fmt"
	"strings"
	"time"
	"workflow-code-test/api/binding/workflow"

	"google.golang.org/protobuf/types/known/wrapperspb"
)

type ExecuteNodeParams struct {
	FormData                    *workflow.WorkflowFormData
	Node                        *workflow.Node
	Nodes                       []*workflow.Node
	PreviousExecutionStepResult *workflow.ExecutionStep
	StepNumber                  int32
	StoredOutputVariables       map[string]string
}

// ExecuteNode executes the required actions for each node (based on node type) and returns
// an ExecutionStep (result) object
func ExecuteNode(params *ExecuteNodeParams) (*workflow.ExecutionStep, map[string]string) {
	var outputVariables map[string]string
	startTime := time.Now()

	// Modify node description based on stored output variables
	validatedDescription := params.Node.Data.GetDescription()
	for key, value := range params.StoredOutputVariables {
		formattedKey := fmt.Sprintf("{{%s}}", key)
		validatedDescription = strings.ReplaceAll(validatedDescription, formattedKey, value)
	}
	executionStepResults := &workflow.ExecutionStep{
		Description: validatedDescription,
		Label:       params.Node.Data.GetLabel(),
		NodeId:      params.Node.Id,
		NodeType:    params.Node.Type,
		Output:      &workflow.ExecutionStep_Output{},
		StepNumber:  params.StepNumber,
		Timestamp:   time.Now().Format(time.RFC3339),
	}

	// Execute nodes based on type
	switch params.Node.Type {
	case "start":
		executionStepResults.Status = "completed"
	case "form":
		formData, stepOutputVariables, err := ExecuteFormNode(&ExecuteFormNodeParams{
			FormData: params.FormData,
			Node:     params.Node,
			Nodes:    params.Nodes,
		})
		outputVariables = stepOutputVariables
		executionStepResults.Output.FormData = formData
		if err != nil {
			executionStepResults.Status = "failed"
			executionStepResults.Output.Message = err.Error()
		} else {
			executionStepResults.Status = "completed"
		}
	case "integration":
		apiResponse, stepOutputVariables, err := ExecuteIntegrationNode(&ExecuteIntegrationNodeParams{
			Node:                  params.Node,
			StoredOutputVariables: params.StoredOutputVariables,
		})
		outputVariables = stepOutputVariables
		executionStepResults.Output.ApiResponse = apiResponse
		if err != nil {
			executionStepResults.Status = "failed"
			executionStepResults.Output.Message = err.Error()
		} else {
			executionStepResults.Status = "completed"
		}
	case "condition":
		conditionResult := ExecuteConditionNode(&ExecuteConditionNodeParams{
			FormData:              params.FormData,
			PreviousExecutionStep: params.PreviousExecutionStepResult,
		})
		executionStepResults.Output.ConditionResult = conditionResult
		executionStepResults.Status = "completed"
	case "email":
		emailContent := ExecuteEmailNode(&ExecuteEmailNodeParams{
			FormData:              params.FormData,
			Node:                  params.Node,
			StoredOutputVariables: params.StoredOutputVariables,
		})
		executionStepResults.Output.EmailContent = emailContent
		executionStepResults.Status = "completed"
	case "end":
		executionStepResults.Status = "completed"
	}
	endTime := time.Now()
	executionStepResults.Duration = wrapperspb.Int64(endTime.Sub(startTime).Milliseconds())
	return executionStepResults, outputVariables
}
