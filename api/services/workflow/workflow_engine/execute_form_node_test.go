package workflow_engine_test

import (
	"fmt"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/testdata"
	"workflow-code-test/api/services/workflow/workflow_engine"

	"github.com/stretchr/testify/assert"
)

func TestExecuteFormNode(t *testing.T) {
	type ExecuteFormNodeParams struct {
		ExpectedError           string
		ExpectedOutputVariables map[string]string
		ExpectedFormData        *workflow.WorkflowFormData
		FormData                *workflow.WorkflowFormData
		Node                    *workflow.Node
		Nodes                   []*workflow.Node
	}

	var formNode *workflow.Node
	for _, node := range testdata.SampleWorkflowNodes.Nodes {
		if node.Type == "form" {
			formNode = node
			break
		}
	}

	testcases := map[string]ExecuteFormNodeParams{
		"Should return correct error if operator is invalid": {
			ExpectedError: "invalid form operator",
			ExpectedFormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "invalid-operator",
				Threshold: 2,
			},
			ExpectedOutputVariables: map[string]string{
				"city":  "Sydney",
				"email": "sample@sample.com",
				"name":  "sample-user",
			},
			FormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "invalid-operator",
				Threshold: 2,
			},
			Node:  formNode,
			Nodes: testdata.SampleWorkflowNodes.Nodes,
		},
		"Should return correct error if form data contains invalid email address": {
			ExpectedError: "invalid form data",
			ExpectedFormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "invalid-email",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			ExpectedOutputVariables: map[string]string{
				"city":  "Sydney",
				"email": "invalid-email",
				"name":  "sample-user",
			},
			FormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "invalid-email",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			Node:  formNode,
			Nodes: testdata.SampleWorkflowNodes.Nodes,
		},
		"Should return correct error if form data contains invalid location": {
			ExpectedError: "invalid location option in form data",
			ExpectedFormData: &workflow.WorkflowFormData{
				City:      "New York",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			ExpectedOutputVariables: map[string]string{
				"city":  "New York",
				"email": "sample@sample.com",
				"name":  "sample-user",
			},
			FormData: &workflow.WorkflowFormData{
				City:      "New York",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			Node:  formNode,
			Nodes: testdata.SampleWorkflowNodes.Nodes,
		},
		"Should not return an error if successful": {
			ExpectedError: "",
			ExpectedFormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			ExpectedOutputVariables: map[string]string{
				"city":  "Sydney",
				"email": "sample@sample.com",
				"name":  "sample-user",
			},
			FormData: &workflow.WorkflowFormData{
				City:      "Sydney",
				Email:     "sample@sample.com",
				Name:      "sample-user",
				Operator:  "greater_than",
				Threshold: 2,
			},
			Node:  formNode,
			Nodes: testdata.SampleWorkflowNodes.Nodes,
		},
	}

	count := 1
	for description, testcase := range testcases {
		testDescription := fmt.Sprintf("%d %s", count, description)
		t.Run(testDescription, func(t *testing.T) {
			formResult, stepOutputVariables, err := workflow_engine.ExecuteFormNode(&workflow_engine.ExecuteFormNodeParams{
				FormData: testcase.FormData,
				Node:     testcase.Node,
				Nodes:    testcase.Nodes,
			})

			// Test the expected output
			assert.Equal(t, testcase.ExpectedFormData.City, formResult.City)
			assert.Equal(t, testcase.ExpectedFormData.Threshold, formResult.Threshold)
			assert.Equal(t, testcase.ExpectedFormData.Operator, formResult.Operator)
			assert.Equal(t, testcase.ExpectedFormData.Name, formResult.Name)
			assert.Equal(t, testcase.ExpectedFormData.Email, formResult.Email)

			// Test the output variables
			for key, value := range testcase.ExpectedOutputVariables {
				outputVariable, ok := stepOutputVariables[key]
				assert.True(t, ok)
				assert.Equal(t, outputVariable, value)
			}

			if testcase.ExpectedError != "" {
				assert.EqualError(t, err, testcase.ExpectedError)
			} else {
				assert.NoError(t, err)
			}
		})
		count++
	}
}
