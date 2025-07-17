package testdata

import (
	"github.com/jackc/pgx/v5"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"
)

type TestStore struct {
	ApiEndpoint string
	DB          *pgx.Conn
}

func (s *TestStore) GetWorkflow(params *store.GetWorkflowParams) (*workflow.WorkflowResponse, error) {
	currentNodes := SampleWorkflowNodes.Nodes
	for _, node := range currentNodes {
		if node.Id == "weather-api" {
			node.Data.Metadata.ApiEndpoint = GetStringPointer(s.ApiEndpoint)
		}
	}
	return &workflow.WorkflowResponse{
		Nodes: currentNodes,
		Edges: SampleWorkflowEdges.Edges,
		Id:    SampleWorkflowId.String(),
	}, nil
}

func (s *TestStore) CreateOrUpdateWorkflow(params *store.CreateOrUpdateWorkflowParams) error {
	return nil
}

func NewTestStore(DB *pgx.Conn, apiEndpoint string) store.Store {
	return &TestStore{
		ApiEndpoint: apiEndpoint,
		DB:          DB,
	}
}
