package store

import (
	"context"
	"fmt"
	"log/slog"
	"workflow-code-test/api/binding/workflow"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"google.golang.org/protobuf/encoding/protojson"
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
	// Marshal nodes to JSON to be persisted as jsonb
	nodes, err := protojson.Marshal(params.WorkflowNodes)
	if err != nil {
		slog.Error("Failed to save workflow nodes", "error", err)
		return err
	}

	// Marshal edges to JSON to be persisted as jsonb
	edges, err := protojson.Marshal(params.WorkflowEdges)
	if err != nil {
		slog.Error("Failed to save workflow edges", "error", err)
		return err
	}

	// Insert the workflow or update nodes and edges if it already exists
	query := `
		INSERT INTO workflows (id, nodes, edges)
		VALUES ($1, $2, $3)
		ON CONFLICT (id) DO UPDATE
		    SET nodes = excluded.nodes,
      			edges = excluded.edges
	`
	if _, err = params.DB.Exec(params.Ctx, query, params.WorkflowId, nodes, edges); err != nil {
		slog.Error("Failed to save workflow workflow", "error", err)
		return fmt.Errorf("store : failed to save or update workflow")
	}
	return nil
}
