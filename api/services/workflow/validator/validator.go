package validator

import (
	"context"
	"golang.org/x/sync/errgroup"
	"workflow-code-test/api/binding/workflow"
)

type ValidateWorkflowParams struct {
	Ctx           context.Context
	WorkflowEdges []*workflow.Edge
	WorkflowNodes []*workflow.Node
}

// ValidateWorkflow validates the nodes and edges of a given workflow
func ValidateWorkflow(params *ValidateWorkflowParams) error {

	// Validate the nodes and edges in a sub-routine
	errorGroup, _ := errgroup.WithContext(context.Background())
	errorGroup.Go(func() error {
		return ValidateNodes(params.WorkflowNodes)
	})
	errorGroup.Go(func() error {
		return ValidateEdges(params.WorkflowEdges, params.WorkflowNodes)
	})

	// Check for validation errors, else return nil
	if validationError := errorGroup.Wait(); validationError != nil {
		return validationError
	}
	return nil
}
