package workflow_engine_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"
	"workflow-code-test/api/services/workflow/testdata"
	"workflow-code-test/api/services/workflow/workflow_engine"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestExecuteWorkflow(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(`{
			"latitude": -33.75,
			"longitude": 151.125,
			"generationtime_ms": 0.04291534423828125,
			"utc_offset_seconds": 0,
			"timezone": "GMT",
			"timezone_abbreviation": "GMT",
			"elevation": 86.0,
			"current_weather_units": {
				"time": "iso8601",
				"interval": "seconds",
				"temperature": "°C",
				"windspeed": "km/h",
				"winddirection": "°",
				"is_day": "",
				"weathercode": "wmo code"
			},
			"current_weather": {
				"time": "2025-07-16T16:45",
				"interval": 900,
				"temperature": 2.4,
				"windspeed": 1.8,
				"winddirection": 349,
				"is_day": 0,
				"weathercode": 0
			}
		}`))
		if err != nil {
			t.Fatal(err)
		}
	}))
	defer ts.Close()
	testStore := testdata.NewTestStore(nil, ts.URL)

	type ExecuteWorkflowParams struct {
		Ctx                      context.Context
		DB                       *pgx.Conn
		ExpectedError            error
		ExpectedExecutionResults *workflow.ExecutionResults
		FormData                 *workflow.WorkflowFormData
		StartTime                string
		Store                    *store.Store
		WorkflowId               uuid.UUID
	}

	testcases := map[string]ExecuteWorkflowParams{
		"Should return the correct execution results": {
			Ctx:           context.Background(),
			DB:            nil,
			ExpectedError: nil,
			ExpectedExecutionResults: &workflow.ExecutionResults{
				Status: "completed",
				Steps: []*workflow.ExecutionStep{
					{
						Description: "Begin weather check workflow",
						Label:       "Start",
						NodeId:      "start",
						NodeType:    "start",
						Status:      "completed",
						StepNumber:  int32(1),
					},
					{
						Description: "Process collected data - name, email, location",
						Label:       "User Input",
						NodeId:      "form",
						NodeType:    "form",
						Status:      "completed",
						StepNumber:  int32(2),
						Output: &workflow.ExecutionStep_Output{
							FormData: &workflow.WorkflowFormData{
								City:      "Sydney",
								Email:     "sample@sample.com",
								Name:      "sample",
								Operator:  "greater_than",
								Threshold: 2,
							},
						},
					},
					{
						Description: "Fetch current temperature for Sydney",
						Label:       "Weather API",
						NodeId:      "weather-api",
						NodeType:    "integration",
						Output: &workflow.ExecutionStep_Output{
							ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
								Data: &structpb.Value{
									Kind: &structpb.Value_NumberValue{
										NumberValue: 2.4,
									},
								},
								Endpoint:   fmt.Sprintf("%s?latitude=-33.8688&longitude=151.2093", ts.URL),
								Method:     "GET",
								StatusCode: 200,
							},
						},
						Status:     "completed",
						StepNumber: int32(3),
					},
					{
						Description: "Evaluate temperature threshold",
						Label:       "Check Condition",
						NodeId:      "condition",
						NodeType:    "condition",
						Output: &workflow.ExecutionStep_Output{
							ConditionResult: &workflow.ExecutionStep_Output_ConditionResult{
								Expression:  "2.40 > 2.00",
								Operator:    "greater_than",
								Result:      true,
								Temperature: 2.4,
								Threshold:   2,
							},
						},
						Status:     "completed",
						StepNumber: int32(4),
					},
					{
						Description: "Email weather alert notification",
						Label:       "Send Alert",
						NodeId:      "email",
						NodeType:    "email",
						Output: &workflow.ExecutionStep_Output{
							EmailContent: &workflow.ExecutionStep_Output_EmailContent{
								Body:    "Weather alert for Sydney! Temperature is 2.40°C!",
								Subject: "Weather Alert",
								To:      "sample@sample.com",
							},
						},
						Status:     "completed",
						StepNumber: int32(5),
					},
					{
						Description: "Workflow execution finished",
						Label:       "Complete",
						NodeId:      "end",
						NodeType:    "end",
						Status:      "completed",
						StepNumber:  int32(6),
					},
				},
				Metadata: &workflow.ExecutionResults_ExecutionMetadata{
					Environment: testdata.GetStringPointer(""),
					TriggeredBy: testdata.GetStringPointer("sample"),
				},
			},
			FormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "sample@sample.com",
				Name:      "sample",
				Operator:  "greater_than",
				Threshold: 2,
			},
			StartTime: time.Now().Format(time.RFC3339),
			Store:     &testStore,
		},
	}

	for description, testcase := range testcases {
		t.Run(description, func(t *testing.T) {
			executionResults, err := workflow_engine.ExecuteWorkflow(&workflow_engine.ExecuteWorkflowParams{
				Ctx:        testcase.Ctx,
				DB:         testcase.DB,
				FormData:   testcase.FormData,
				StartTime:  testcase.StartTime,
				Store:      testcase.Store,
				WorkflowId: testcase.WorkflowId,
			})
			assert.Equal(t, testcase.ExpectedExecutionResults.Status, executionResults.Status)
			assert.Equal(t, testcase.ExpectedExecutionResults.Metadata.TriggeredBy, executionResults.Metadata.TriggeredBy)
			assert.Equal(t, testcase.ExpectedExecutionResults.Metadata.Environment, executionResults.Metadata.Environment)
			if err == nil {
				assert.Equal(t, len(testcase.ExpectedExecutionResults.Steps), len(executionResults.Steps))
				for index, step := range testcase.ExpectedExecutionResults.Steps {
					assert.Equal(t, step.Description, executionResults.Steps[index].Description)
					assert.Equal(t, step.Label, executionResults.Steps[index].Label)
					assert.Equal(t, step.NodeId, executionResults.Steps[index].NodeId)
					assert.Equal(t, step.NodeType, executionResults.Steps[index].NodeType)
					assert.Equal(t, step.Status, executionResults.Steps[index].Status)
					assert.Equal(t, step.StepNumber, executionResults.Steps[index].StepNumber)

					if step.NodeType == "form" {
						assert.Equal(t, step.Output.FormData.City, executionResults.Steps[index].Output.FormData.City)
						assert.Equal(t, step.Output.FormData.Email, executionResults.Steps[index].Output.FormData.Email)
						assert.Equal(t, step.Output.FormData.Name, executionResults.Steps[index].Output.FormData.Name)
						assert.Equal(t, step.Output.FormData.Operator, executionResults.Steps[index].Output.FormData.Operator)
						assert.Equal(t, step.Output.FormData.Threshold, executionResults.Steps[index].Output.FormData.Threshold)
					} else if step.NodeType == "integration" {
						assert.Equal(t, step.Output.ApiResponse.Data.GetNumberValue(), executionResults.Steps[index].Output.ApiResponse.Data.GetNumberValue())
						assert.Equal(t, step.Output.ApiResponse.Endpoint, executionResults.Steps[index].Output.ApiResponse.Endpoint)
						assert.Equal(t, step.Output.ApiResponse.Method, executionResults.Steps[index].Output.ApiResponse.Method)
						assert.Equal(t, step.Output.ApiResponse.StatusCode, executionResults.Steps[index].Output.ApiResponse.StatusCode)
					} else if step.NodeType == "condition" {
						assert.Equal(t, step.Output.ConditionResult.Result, executionResults.Steps[index].Output.ConditionResult.Result)
						assert.Equal(t, step.Output.ConditionResult.Expression, executionResults.Steps[index].Output.ConditionResult.Expression)
						assert.Equal(t, step.Output.ConditionResult.Operator, executionResults.Steps[index].Output.ConditionResult.Operator)
						assert.Equal(t, step.Output.ConditionResult.Temperature, executionResults.Steps[index].Output.ConditionResult.Temperature)
						assert.Equal(t, step.Output.ConditionResult.Threshold, executionResults.Steps[index].Output.ConditionResult.Threshold)
					} else if step.NodeType == "email" {
						assert.Equal(t, step.Output.EmailContent.Body, executionResults.Steps[index].Output.EmailContent.Body)
						assert.Equal(t, step.Output.EmailContent.Subject, executionResults.Steps[index].Output.EmailContent.Subject)
						assert.Equal(t, step.Output.EmailContent.To, executionResults.Steps[index].Output.EmailContent.To)
					}
				}
			}
		})
	}
}
