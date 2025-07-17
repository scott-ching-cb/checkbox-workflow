package workflow_engine_test

import (
	"fmt"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/workflow_engine"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestExecuteEmailNode(t *testing.T) {
	type ExecuteEmailNodeParams struct {
		ExpectedResult        *workflow.ExecutionStep_Output_EmailContent
		FormData              *workflow.WorkflowFormData
		Node                  *workflow.Node
		StoredOutputVariables map[string]string
	}

	testcases := map[string]ExecuteEmailNodeParams{
		"Should return the correct email template output": {
			ExpectedResult: &workflow.ExecutionStep_Output_EmailContent{
				Body:    "Weather alert for Sydney! Temperature is 1.00°C!",
				Subject: "Weather Alert",
				To:      "sample@sample.com",
			},
			FormData: &workflow.WorkflowFormData{
				Email: "sample@sample.com",
			},
			Node: &workflow.Node{
				Id:   "email",
				Type: "email",
				Position: &workflow.Node_Position{
					X: 1096,
					Y: 88,
				},
				Data: &workflow.Node_Data{
					Label:       "Send Alert",
					Description: "Email weather alert notification",
					Metadata: &workflow.MetaData{
						HasHandles: &workflow.MetaData_HasHandles{
							Source: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: true,
								},
							},
							Target: &structpb.Value{
								Kind: &structpb.Value_BoolValue{
									BoolValue: true,
								},
							},
						},
						InputVariables: []string{"name", "city", "temperature"},
						EmailTemplate: &workflow.MetaData_EmailTemplate{
							Body:    "Weather alert for {{city}}! Temperature is {{temperature}}°C!",
							Subject: "Weather Alert",
						},
						OutputVariables: []string{"emailSent"},
					},
				},
			},
			StoredOutputVariables: map[string]string{
				"city":        "Sydney",
				"temperature": "1.00",
			},
		},
	}

	count := 1
	for description, testcase := range testcases {
		testDescription := fmt.Sprintf("%d %s", count, description)
		t.Run(testDescription, func(t *testing.T) {
			emailResult := workflow_engine.ExecuteEmailNode(&workflow_engine.ExecuteEmailNodeParams{
				FormData:              testcase.FormData,
				Node:                  testcase.Node,
				StoredOutputVariables: testcase.StoredOutputVariables,
			})
			assert.Equal(t, testcase.ExpectedResult.To, emailResult.To)
			assert.Equal(t, testcase.ExpectedResult.Subject, emailResult.Subject)
			assert.Equal(t, testcase.ExpectedResult.Body, emailResult.Body)
			assert.NotNil(t, emailResult.Timestamp)
		})
		count++
	}
}
