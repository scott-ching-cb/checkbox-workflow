package workflow_engine

import (
	"context"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type ExecuteWorkflowParams struct {
	Ctx        context.Context
	DB         *pgx.Conn
	FormData   *workflow.WorkflowFormData
	StartTime  string
	Store      *store.Store
	WorkflowId uuid.UUID
}

// ExecuteWorkflow retrieves the persisted workflow from the database and executes the workflow
func ExecuteWorkflow(params *ExecuteWorkflowParams) (*workflow.ExecutionResults, error) {
	return nil, nil
}
