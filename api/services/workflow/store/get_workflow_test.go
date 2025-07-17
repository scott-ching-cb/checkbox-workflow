package store_test

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"
	"workflow-code-test/api/services/workflow/testdata"
)

func validateNodes(t *testing.T, testcaseNodes []*workflow.Node, nodes []*workflow.Node) {
	assert.NotNil(t, nodes)
	assert.Equal(t, len(testcaseNodes), len(nodes))
	for index, node := range testcaseNodes {

		assert.Equal(t, node.Id, nodes[index].Id)
		assert.Equal(t, node.Type, nodes[index].Type)
		assert.Equal(t, node.Position.X, nodes[index].Position.X)
		assert.Equal(t, node.Position.Y, nodes[index].Position.Y)
		assert.Equal(t, node.Data.Label, nodes[index].Data.Label)
		assert.Equal(t, node.Data.Description, nodes[index].Data.Description)
		if node.Type != "condition" {
			assert.Equal(t, node.Data.Metadata.HasHandles.Source.GetBoolValue(), nodes[index].Data.Metadata.HasHandles.Source.GetBoolValue())
			assert.Equal(t, node.Data.Metadata.HasHandles.Target.GetBoolValue(), nodes[index].Data.Metadata.HasHandles.Target.GetBoolValue())
		}

		if node.Type == "form" {
			assert.Equal(t, node.Data.Metadata.InputFields, nodes[index].Data.Metadata.InputFields)
			assert.Equal(t, node.Data.Metadata.OutputVariables, nodes[index].Data.Metadata.OutputVariables)
		} else if node.Type == "integration" {
			assert.Equal(t, node.Data.Metadata.InputVariables, nodes[index].Data.Metadata.InputVariables)
			assert.Equal(t, node.Data.Metadata.ApiEndpoint, nodes[index].Data.Metadata.ApiEndpoint)
			assert.Equal(t, node.Data.Metadata.OutputVariables, nodes[index].Data.Metadata.OutputVariables)
			for optionIndex, option := range node.Data.Metadata.Options {
				assert.Equal(t, option.City, nodes[index].Data.Metadata.Options[optionIndex].City)
				assert.Equal(t, option.Lat, nodes[index].Data.Metadata.Options[optionIndex].Lat)
				assert.Equal(t, option.Lon, nodes[index].Data.Metadata.Options[optionIndex].Lon)
			}
		} else if node.Type == "condition" {
			assert.Equal(t, node.Data.Metadata.ConditionExpression, nodes[index].Data.Metadata.ConditionExpression)
		} else if node.Type == "email" {
			assert.Equal(t, node.Data.Metadata.InputVariables, nodes[index].Data.Metadata.InputVariables)
			assert.Equal(t, node.Data.Metadata.EmailTemplate.Subject, nodes[index].Data.Metadata.EmailTemplate.Subject)
			assert.Equal(t, node.Data.Metadata.EmailTemplate.Body, nodes[index].Data.Metadata.EmailTemplate.Body)
		}
	}
}

func validateEdges(t *testing.T, testcaseEdges []*workflow.Edge, edges []*workflow.Edge) {
	assert.NotNil(t, edges)
	assert.Equal(t, len(testcaseEdges), len(edges))
	for index, edge := range testcaseEdges {
		assert.Equal(t, edge.Id, edges[index].Id)
		assert.Equal(t, edge.Source, edges[index].Source)
		assert.Equal(t, edge.Target, edges[index].Target)
		assert.Equal(t, edge.Type, edges[index].Type)
		assert.Equal(t, edge.Animated, edges[index].Animated)
		assert.Equal(t, edge.Style.Stroke, edges[index].Style.Stroke)
		assert.Equal(t, edge.Style.StrokeWidth, edges[index].Style.StrokeWidth)
		assert.Equal(t, *edge.Label, *edges[index].Label)
		if edge.LabelStyle != nil {
			assert.Equal(t, edge.LabelStyle.FontWeight, edges[index].LabelStyle.FontWeight)
			assert.Equal(t, edge.LabelStyle.Fill, edges[index].LabelStyle.Fill)
		}
	}
}

func TestGetWorkflow(t *testing.T) {
	ctx := context.Background()
	testWorkflowId, err := uuid.Parse("550e8400-e29b-41d4-a716-446655440000")
	if err != nil {
		t.Error(err)
	}

	type GetWorkflowParams struct {
		Ctx                    context.Context
		ExpectedError          string
		ExpectedWorkflowConfig *workflow.WorkflowResponse
		GetDB                  func() *pgx.Conn
		Seed                   string
		WorkflowId             uuid.UUID
	}

	testcases := map[string]GetWorkflowParams{
		"Should retrieve the stored workflow configuration": {
			Ctx: ctx,
			ExpectedWorkflowConfig: &workflow.WorkflowResponse{
				Id:    testWorkflowId.String(),
				Nodes: testdata.SampleWorkflowNodes.Nodes,
				Edges: testdata.SampleWorkflowEdges.Edges,
			},
			GetDB: func() *pgx.Conn {
				return DBConnection
			},
			Seed:       testdata.SampleWorkflowInsert,
			WorkflowId: testWorkflowId,
		},
		"Should return correct error if workflow with corresponding id does not exist": {
			Ctx:           ctx,
			ExpectedError: fmt.Sprintf("no workflow found with id %s", uuid.Nil),
			GetDB: func() *pgx.Conn {
				return DBConnection
			},
			Seed:       testdata.SampleWorkflowInsert,
			WorkflowId: uuid.Nil,
		},
	}

	//
	for desc, testcase := range testcases {
		t.Run(desc, func(t *testing.T) {
			dbConnection := testcase.GetDB()
			_, err := dbConnection.Exec(testcase.Ctx, "DELETE FROM workflows")
			if err != nil {
				t.Fatalf("failed to delete workflows: %v", err)
			}
			_, err = dbConnection.Exec(testcase.Ctx, testcase.Seed)
			if err != nil {
				t.Fatalf("failed to insert seed data: %v", err)
			}
			testStore := store.NewStore(testcase.GetDB())
			workflowConfig, err := testStore.GetWorkflow(&store.GetWorkflowParams{
				Ctx:        testcase.Ctx,
				DB:         testcase.GetDB(),
				WorkflowId: testcase.WorkflowId,
			})
			if testcase.ExpectedError != "" {
				assert.EqualError(t, err, testcase.ExpectedError)
			} else {
				assert.Nil(t, err)
				assert.NotNil(t, workflowConfig)
				validateNodes(t, testcase.ExpectedWorkflowConfig.Nodes, workflowConfig.Nodes)
				validateEdges(t, testcase.ExpectedWorkflowConfig.Edges, workflowConfig.Edges)
			}
		})
	}
}
