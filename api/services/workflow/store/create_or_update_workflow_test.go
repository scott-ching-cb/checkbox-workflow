package store_test

import (
	"context"
	"fmt"
	"testing"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/pkg/db"
	"workflow-code-test/api/services/workflow/store"
	"workflow-code-test/api/services/workflow/testdata"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
)

func TestCreateOrUpdateWorkflow(t *testing.T) {
	ctx := context.Background()
	sampleWorkflowId := uuid.New()
	sampleWorkflowEdges := testdata.SampleWorkflowEdges
	sampleWorkflowNodes := testdata.SampleWorkflowNodes

	type CreateOrUpdateWorkflowParams struct {
		Ctx           context.Context
		Error         bool
		ErrorMessage  string
		GetDB         func() *pgx.Conn
		WorkflowEdges *workflow.Edges
		WorkflowId    uuid.UUID
		WorkflowNodes *workflow.Nodes
	}
	testcases := map[string]CreateOrUpdateWorkflowParams{
		"Should persist new workflow definition": {
			Ctx: ctx,
			GetDB: func() *pgx.Conn {
				return DBConnection
			},
			WorkflowEdges: sampleWorkflowEdges,
			WorkflowId:    sampleWorkflowId,
			WorkflowNodes: sampleWorkflowNodes,
		},
		"Should throw correct error if database input (workflow id) is invalid": {
			Ctx:          ctx,
			Error:        true,
			ErrorMessage: "store : failed to save or update workflow",
			GetDB: func() *pgx.Conn {
				pgxConn, err := db.GetPool().Acquire(context.Background())
				if err != nil {
					t.Fatalf("failed to acquire database connection")
				}
				newConnection := pgxConn.Conn()
				if err := newConnection.Close(ctx); err != nil {
					t.Fatalf("failed to close database connection")
				}
				return newConnection
			},
			WorkflowEdges: sampleWorkflowEdges,
			WorkflowId:    uuid.Nil,
			WorkflowNodes: sampleWorkflowNodes,
		},
	}

	index := 1
	for description, testcase := range testcases {
		testcaseDescription := fmt.Sprintf("%d %s", index, description)
		t.Run(testcaseDescription, func(t *testing.T) {
			testStore := store.NewStore(testcase.GetDB())
			err := testStore.CreateOrUpdateWorkflow(&store.CreateOrUpdateWorkflowParams{
				Ctx:           testcase.Ctx,
				DB:            testcase.GetDB(),
				WorkflowEdges: testcase.WorkflowEdges,
				WorkflowId:    testcase.WorkflowId,
				WorkflowNodes: testcase.WorkflowNodes,
			})
			if testcase.Error {
				assert.EqualError(t, err, testcase.ErrorMessage)
			} else {
				assert.NoError(t, err, "Error in workflow creation")
			}
		})
		index++
	}
}
