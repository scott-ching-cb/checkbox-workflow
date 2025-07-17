package workflow_engine_test

import (
	"fmt"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/workflow_engine"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/structpb"
)

func TestExecuteConditionNode(t *testing.T) {
	type ExecuteConditionNodeParams struct {
		ExpectedResult        *workflow.ExecutionStep_Output_ConditionResult
		FormData              *workflow.WorkflowFormData
		PreviousExecutionStep *workflow.ExecutionStep
	}

	testcases := map[string]ExecuteConditionNodeParams{
		"Should return result true if threshold is reached": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "1.00 < 2.00",
				Operator:    "less_than",
				Result:      true,
				Temperature: 1.00,
				Threshold:   2.00,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "less_than",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 1,
							},
						},
					},
				},
			},
		},
		"Should return correct result if greater than operator is used": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "3.00 > 2.00",
				Operator:    "greater_than",
				Result:      true,
				Temperature: 3.00,
				Threshold:   2.00,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "greater_than",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 3,
							},
						},
					},
				},
			},
		},
		"Should return correct result if greater or equal to operator is used": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "3.00 >= 2.00",
				Operator:    "greater_than_or_equal",
				Result:      true,
				Temperature: 3.00,
				Threshold:   2.00,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "greater_than_or_equal",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 3,
							},
						},
					},
				},
			},
		},
		"Should return correct result if less than or equal to operator is used": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "2.00 <= 2.00",
				Operator:    "less_than_or_equal",
				Result:      true,
				Temperature: 2.00,
				Threshold:   2.00,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "less_than_or_equal",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 2,
							},
						},
					},
				},
			},
		},
		"Should return correct result if equal operator is used": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "2.00 == 2.00",
				Operator:    "equals",
				Result:      true,
				Temperature: 2,
				Threshold:   2,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "equals",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 2,
							},
						},
					},
				},
			},
		},
		"Should return result false if threshold is not reached": {
			ExpectedResult: &workflow.ExecutionStep_Output_ConditionResult{
				Expression:  "1.00 == 2.00",
				Operator:    "equals",
				Result:      false,
				Temperature: 1,
				Threshold:   2,
			},
			FormData: &workflow.WorkflowFormData{
				Operator:  "equals",
				Threshold: 2,
			},
			PreviousExecutionStep: &workflow.ExecutionStep{
				Output: &workflow.ExecutionStep_Output{
					ApiResponse: &workflow.ExecutionStep_Output_ApiResponse{
						Data: &structpb.Value{
							Kind: &structpb.Value_NumberValue{
								NumberValue: 1,
							},
						},
					},
				},
			},
		},
	}

	count := 1
	for description, testcase := range testcases {
		testDescription := fmt.Sprintf("%d %s", count, description)
		t.Run(testDescription, func(t *testing.T) {
			conditionResult := workflow_engine.ExecuteConditionNode(&workflow_engine.ExecuteConditionNodeParams{
				FormData:              testcase.FormData,
				PreviousExecutionStep: testcase.PreviousExecutionStep,
			})
			assert.Equal(t, testcase.ExpectedResult.Result, conditionResult.Result)
			assert.Equal(t, testcase.ExpectedResult.Threshold, conditionResult.Threshold)
			assert.Equal(t, testcase.ExpectedResult.Temperature, conditionResult.Temperature)
			assert.Equal(t, testcase.ExpectedResult.Expression, conditionResult.Expression)
			assert.Equal(t, testcase.ExpectedResult.Operator, conditionResult.Operator)
		})
		count++
	}
}
