package workflow_engine

import (
	"fmt"
	"workflow-code-test/api/binding/workflow"
)

type ExecuteConditionNodeParams struct {
	FormData              *workflow.WorkflowFormData
	PreviousExecutionStep *workflow.ExecutionStep
}

// ExecuteConditionNode evaluates the condition provided in the form input (threshold) against the integration API response
func ExecuteConditionNode(params *ExecuteConditionNodeParams) *workflow.ExecutionStep_Output_ConditionResult {
	currentTemperature := params.PreviousExecutionStep.Output.ApiResponse.Data.GetNumberValue()
	operatorFunction := ConditionOperatorToFunctionMap[params.FormData.Operator]
	isThresholdReached := operatorFunction(currentTemperature, params.FormData.Threshold)
	operatorAsString := ConditionOperatorToStringMap[params.FormData.Operator]
	return &workflow.ExecutionStep_Output_ConditionResult{
		Expression:  fmt.Sprintf("%.2f %s %.2f", currentTemperature, operatorAsString, params.FormData.Threshold),
		Operator:    params.FormData.Operator,
		Result:      isThresholdReached,
		Temperature: params.PreviousExecutionStep.Output.ApiResponse.Data.GetNumberValue(),
		Threshold:   params.FormData.Threshold,
	}
}
