package store

import (
	"context"
	"workflow-code-test/api/binding/workflow"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type CreateOrUpdateWorkflowParams struct {
	Ctx           context.Context
	DB            *pgx.Conn
	WorkflowEdges *workflow.Edges
	WorkflowId    uuid.UUID
	WorkflowNodes *workflow.Nodes
}

// CreateOrUpdateWorkflow creates a workflow with the given id if not exists, else it updates the nodes and edges
func (s *DatabaseStore) CreateOrUpdateWorkflow(params *CreateOrUpdateWorkflowParams) error {
	return nil
}
