package validator

import (
	"context"
	"workflow-code-test/api/binding/workflow"
)

type ValidateWorkflowParams struct {
	Ctx           context.Context
	WorkflowEdges []*workflow.Edge
	WorkflowNodes []*workflow.Node
}

// ValidateWorkflow validates the nodes and edges of a given workflow
func ValidateWorkflow(params *ValidateWorkflowParams) error {
	return nil
}
