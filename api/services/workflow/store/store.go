package store

import (
	"github.com/jackc/pgx/v5"
	"workflow-code-test/api/binding/workflow"
)

type Store interface {
	GetWorkflow(params *GetWorkflowParams) (*workflow.WorkflowResponse, error)
	CreateOrUpdateWorkflow(params *CreateOrUpdateWorkflowParams) error
}

type DatabaseStore struct {
	DB *pgx.Conn
}

func NewStore(DB *pgx.Conn) Store {
	return &DatabaseStore{
		DB: DB,
	}
}
