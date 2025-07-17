package workflow_engine

import (
	"fmt"
	"net/mail"
	"workflow-code-test/api/binding/workflow"
)

type ExecuteFormNodeParams struct {
	FormData *workflow.WorkflowFormData
	Node     *workflow.Node
	Nodes    []*workflow.Node
}

// ExecuteFormNode performs validation on input provided to a workflow form node
func ExecuteFormNode(params *ExecuteFormNodeParams) (*workflow.WorkflowFormData, map[string]string, error) {
	// Initialise the execution step's result variable
	formData := params.FormData
	formDataResponse := &workflow.WorkflowFormData{
		City:      formData.GetCity(),
		Email:     formData.GetEmail(),
		Name:      formData.GetName(),
		Operator:  formData.GetOperator(),
		Threshold: formData.GetThreshold(),
	}

	// Get output variables from the node
	outputVariables := make(map[string]string)
	reflectedFormDataResponse := formDataResponse.ProtoReflect()
	messageDescriptor := reflectedFormDataResponse.Descriptor()
	messageDescriptorFields := messageDescriptor.Fields()
	for i := 0; i < messageDescriptorFields.Len(); i++ {
		fieldDescriptor := messageDescriptorFields.Get(i)
		fieldName := fieldDescriptor.JSONName()
		for _, outputVariable := range params.Node.Data.Metadata.OutputVariables {
			if outputVariable == fieldName {
				outputVariables[outputVariable] = reflectedFormDataResponse.Get(fieldDescriptor).String()
			}
		}
	}

	// Perform validation to determine if operator is valid
	if _, ok := ConditionOperatorToFunctionMap[formData.Operator]; !ok {
		return formDataResponse, outputVariables, fmt.Errorf("invalid form operator")
	}

	// Perform high-level validation on form data and condition operator
	_, emailValidationError := mail.ParseAddress(formData.GetEmail())
	hasInvalidFormData := formData.Name == "" || emailValidationError != nil || formData.City == ""
	if hasInvalidFormData {
		return formDataResponse, outputVariables, fmt.Errorf("invalid form data")
	}

	// Retrieve integration node from list of workflow nodes
	var integrationNode *workflow.Node
	for _, workflowNode := range params.Nodes {
		if workflowNode.Type == "integration" {
			integrationNode = workflowNode
		}
	}

	// Determine whether the form input has a location within the integration node options
	hasValidLocation := false
	if integrationNode != nil {
		for _, option := range integrationNode.Data.Metadata.Options {
			if option.City == formData.City {
				hasValidLocation = true
			}
		}
	}
	if !hasValidLocation {
		return formDataResponse, outputVariables, fmt.Errorf("invalid location option in form data")
	}
	return formDataResponse, outputVariables, nil
}
