package workflow

import (
	"io"
	"log/slog"
	"net/http"
	"time"
	"workflow-code-test/api/binding/workflow"
	"workflow-code-test/api/services/workflow/store"
	"workflow-code-test/api/services/workflow/validator"
	"workflow-code-test/api/services/workflow/workflow_engine"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

// parseRequestBody parses the request object and returns an unmarshalled workflow config and form data
func parseRequestBody(request *http.Request) (*workflow.ExecuteWorkflowRequest, error) {

	// Ensure the request body is closed after reading
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			slog.Error("failed to close request body on workflow execution", "error", err)
		}
	}(request.Body)

	// Read request body as bytes and unmarshal with proto-json
	requestBodyReader, err := io.ReadAll(request.Body)
	if err != nil {
		return nil, err
	}
	requestBody := &workflow.ExecuteWorkflowRequest{}
	invalidRequestBodyError := protojson.Unmarshal(requestBodyReader, requestBody)
	if invalidRequestBodyError != nil {
		return nil, invalidRequestBodyError
	}
	return requestBody, nil
}

func getInvalidRequestError() ([]byte, error) {
	rawResponseError := &workflow.ExecuteError{
		Message: "Invalid request body",
	}
	responseError, err := proto.Marshal(rawResponseError)
	if err != nil {
		slog.Error("api : error marshaling response error")
		return nil, err
	}
	return responseError, nil
}

// HandleExecuteWorkflow validates, saves, and executes a workflow and returns a summary of the workflow's execution
func (s *Service) HandleExecuteWorkflow(w http.ResponseWriter, r *http.Request) {
	// Fetch id from path params and generate current timestamp
	id := mux.Vars(r)["id"]
	slog.Debug("api : handling workflow execution for id", "id", id)
	currentTime := time.Now().Format(time.RFC3339)
	workflowId, err := uuid.Parse(id)
	if err != nil {
		slog.Error("api : error parsing id in path parameters", id, err)
		responseError, err := getInvalidRequestError()
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			slog.Error("api : error marshaling response error", "id", id)
			return
		}
		if _, err = w.Write(responseError); err != nil {
			slog.Error("api : error returning response for HandleExecuteWorkflow", "id", id)
		}
		return
	}

	// Parse the request body and marshal into GoLang struct
	requestBody, err := parseRequestBody(r)
	if err != nil {
		slog.Error("api : error validating request body", "id", id)
		w.WriteHeader(http.StatusBadRequest)
		rawResponseError := &workflow.ExecuteError{
			Message: "Invalid request body",
		}
		responseError, err := proto.Marshal(rawResponseError)
		if err != nil {
			slog.Error("api : error marshaling response error", "id", id)
			return
		}
		if _, err = w.Write(responseError); err != nil {
			slog.Error("api : error returning response", "id", id)
		}
		return
	}

	// Execute validation on workflow (edges and nodes)
	validationError := validator.ValidateWorkflow(&validator.ValidateWorkflowParams{
		Ctx:           r.Context(),
		WorkflowEdges: requestBody.WorkflowEdges,
		WorkflowNodes: requestBody.WorkflowNodes,
	})
	if validationError != nil {
		slog.Error("api : error validating request body", id, validationError)
		responseError, err := getInvalidRequestError()
		w.WriteHeader(http.StatusBadRequest)
		if err != nil {
			slog.Error("api : error marshaling response error", "id", id)
			return
		}
		if _, err = w.Write(responseError); err != nil {
			slog.Error("api : error returning response for HandleExecuteWorkflow", "id", id)
		}
		return
	}

	// Update workflow nodes and edges stored in the database
	if err = s.Store.CreateOrUpdateWorkflow(&store.CreateOrUpdateWorkflowParams{
		Ctx:        r.Context(),
		DB:         s.DB,
		WorkflowId: workflowId,
		WorkflowEdges: &workflow.Edges{
			Edges: requestBody.WorkflowEdges,
		},
		WorkflowNodes: &workflow.Nodes{
			Nodes: requestBody.WorkflowNodes,
		},
	}); err != nil {
		slog.Error("api : failed to persist workflow", "id", id, "error", err)
		w.WriteHeader(http.StatusBadRequest)
		rawResponseError := &workflow.ExecuteError{
			Message: "Invalid workflow id or request body",
		}
		responseError, err := proto.Marshal(rawResponseError)
		if err != nil {
			slog.Error("api : error marshaling response error", "id", id)
			return
		}
		if _, err = w.Write(responseError); err != nil {
			slog.Error("api : error returning response", "id", id)
		}
		return
	}
	slog.Debug("api : successfully persisted workflow", "id", id)

	// Execute the workflow
	executionResults, err := workflow_engine.ExecuteWorkflow(&workflow_engine.ExecuteWorkflowParams{
		Ctx:        r.Context(),
		DB:         s.DB,
		StartTime:  currentTime,
		Store:      &s.Store,
		FormData:   requestBody.FormData,
		WorkflowId: workflowId,
	})
	if err != nil {
		slog.Error("api : failed to execute workflow", "id", id, "err", err)
		w.WriteHeader(http.StatusBadRequest)
		rawResponseError := &workflow.ExecuteError{
			Message: "invalid workflow id or request body",
		}
		responseError, err := proto.Marshal(rawResponseError)
		if err != nil {
			slog.Error("api : error marshaling response error", "id", id)
			return
		}
		if _, err = w.Write(responseError); err != nil {
			slog.Error("api : error returning response", "id", id)
		}
		return
	}
	slog.Debug("api : successfully executed workflow", "id", id)

	// Return summary of execution as response
	executionResultsJSON, err := protojson.Marshal(executionResults)
	if err != nil {
		slog.Error("api : failed to marshal response for workflow execution results", "id", id, "err", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write(executionResultsJSON); err != nil {
		slog.Error("api : failed to write body for workflow execution response", "err", err)
	}
}
