package workflow_engine

import (
	"fmt"
	"strings"
	"time"
	"workflow-code-test/api/binding/workflow"
)

type ExecuteEmailNodeParams struct {
	FormData              *workflow.WorkflowFormData
	Node                  *workflow.Node
	StoredOutputVariables map[string]string
}

// ExecuteEmailNode prepares an email payload with a body, subject, timestamp and recipient (to)
func ExecuteEmailNode(params *ExecuteEmailNodeParams) *workflow.ExecutionStep_Output_EmailContent {
	validatedEmailBody := params.Node.Data.Metadata.EmailTemplate.Body
	for key, value := range params.StoredOutputVariables {
		formattedKey := fmt.Sprintf("{{%s}}", key)
		validatedEmailBody = strings.ReplaceAll(validatedEmailBody, formattedKey, value)
	}
	timestamp := time.Now().Format(time.RFC3339)
	return &workflow.ExecutionStep_Output_EmailContent{
		Body:      validatedEmailBody,
		Subject:   params.Node.Data.Metadata.EmailTemplate.Subject,
		Timestamp: &timestamp,
		To:        params.FormData.Email,
	}
}
