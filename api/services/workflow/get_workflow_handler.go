package workflow

import (
	"log/slog"
	"net/http"
	"workflow-code-test/api/services/workflow/store"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
)

// HandleGetWorkflow retrieves a workflow (id, nodes and edges) from the database
func (s *Service) HandleGetWorkflow(w http.ResponseWriter, r *http.Request) {
	// Get and parse workflow id as UUID
	id := mux.Vars(r)["id"]
	slog.Debug("Returning workflow definition for id", "id", id)
	workflowId, err := uuid.Parse(id)
	if err != nil {
		slog.Error("Failed to parse workflow id", id, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Retrieve workflow nodes and edges (workflow configuration)
	workflowResponse, err := s.Store.GetWorkflow(&store.GetWorkflowParams{
		Ctx:        r.Context(),
		DB:         s.DB,
		WorkflowId: workflowId,
	})
	if err != nil {
		slog.Error("Failed to retrieve stored workflow definition", "err", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Marshal workflow to JSON to be returned
	workflowResponseJSON, err := protojson.Marshal(workflowResponse)
	if err != nil {
		slog.Error("Failed to serialize stored workflow definition to JSON", "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Write request header and body for server response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(workflowResponseJSON); err != nil {
		slog.Error("Failed to write response body", "err", err)
	}
}
