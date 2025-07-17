package validator_test

import (
	"fmt"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/testdata"
	"workflow-code-test/api/services/workflow/validator"

	"github.com/stretchr/testify/assert"
)

func TestValidateEdges(t *testing.T) {

	type ValidateEdgesParams struct {
		Edges         []*workflow.Edge
		ExpectedError string
		Nodes         []*workflow.Node
	}

	testcases := map[string]ValidateEdgesParams{
		"Should return correct error if there are duplicate edge ids": {
			Edges: append(testdata.SampleWorkflowEdges.Edges, &workflow.Edge{
				Id:     "e1",
				Source: "condition",
				Target: "end",
			}),
			ExpectedError: "duplicate edges with id e1",
			Nodes:         testdata.SampleWorkflowNodes.Nodes,
		},
		"Should return correct error if the source or target nodes are not in the workflow": {
			Edges: append(testdata.SampleWorkflowEdges.Edges, &workflow.Edge{
				Id:     "e7",
				Source: "invalid",
				Target: "also-invalid",
			}),
			ExpectedError: "invalid target or source node for edge e7",
			Nodes:         testdata.SampleWorkflowNodes.Nodes,
		},
		"Should return no error for a valid workflow": {
			Edges:         testdata.SampleWorkflowEdges.Edges,
			ExpectedError: "",
			Nodes:         testdata.SampleWorkflowNodes.Nodes,
		},
	}

	count := 1
	for description, testcase := range testcases {
		testDescription := fmt.Sprintf("%d %s", count, description)
		t.Run(testDescription, func(t *testing.T) {
			err := validator.ValidateEdges(testcase.Edges, testcase.Nodes)
			if testcase.ExpectedError != "" {
				assert.EqualError(t, err, testcase.ExpectedError)
			} else {
				assert.NoError(t, err)
			}
		})
		count++
	}
}
