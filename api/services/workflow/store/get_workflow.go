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

type GetWorkflowParams struct {
	Ctx        context.Context
	DB         *pgx.Conn
	WorkflowId uuid.UUID
}

// GetWorkflow retrieves the workflow nodes and edges for a give Workflow ID
func (s *DatabaseStore) GetWorkflow(params *GetWorkflowParams) (*workflow.WorkflowResponse, error) {
	// Check database connection state and defer de-allocation of prepared statements due to single db connection
	if params.DB.PgConn().IsBusy() {
		slog.Error("store : database connection busy")
		return nil, fmt.Errorf("database connection busy")
	}

	defer func(DB *pgx.Conn, ctx context.Context) {
		if err := DB.DeallocateAll(ctx); err != nil {
			slog.Error("Failed to deallocate prepared statements", "error", err)
		}
	}(params.DB, params.Ctx)

	// Query database for workflow nodes and edges
	query := `
		SELECT nodes, edges
		FROM workflows
		WHERE id = @id
	`
	rows, err := params.DB.Query(params.Ctx, query, pgx.NamedArgs{"id": params.WorkflowId})
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Scan nodes and edges to memory
	var nodesAsBytes []byte
	var edgesAsBytes []byte
	if rows.Next() {
		if err := rows.Scan(&nodesAsBytes, &edgesAsBytes); err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("no workflow found with id %s", params.WorkflowId.String())
	}

	// Unmarshal the nodes and edges
	var nodes workflow.Nodes
	if err := protojson.Unmarshal(nodesAsBytes, &nodes); err != nil {
		return nil, err
	}

	var edges workflow.Edges
	if err := protojson.Unmarshal(edgesAsBytes, &edges); err != nil {
		return nil, err
	}

	// Scan the workflow nodes and edges into GO Structs (from bytes)
	return &workflow.WorkflowResponse{
		Edges: edges.Edges,
		Id:    params.WorkflowId.String(),
		Nodes: nodes.Nodes,
	}, nil
}
