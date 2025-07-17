package store

import (
	"context"
	"workflow-code-test/api/binding/workflow"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type GetWorkflowParams struct {
	Ctx        context.Context
	DB         *pgx.Conn
	WorkflowId uuid.UUID
}

// GetWorkflow retrieves the workflow nodes and edges for a give Workflow ID
func (s *DatabaseStore) GetWorkflow(params *GetWorkflowParams) (*workflow.WorkflowResponse, error) {
	return nil, nil
}
