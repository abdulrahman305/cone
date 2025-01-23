// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// The ListWorkflowExecutionsResponse message.
type ListWorkflowExecutionsResponse struct {
	// The nextPageToken field.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	// The workflowExecutions field.
	WorkflowExecutions []WorkflowExecution `json:"workflowExecutions,omitempty"`
}

func (o *ListWorkflowExecutionsResponse) GetNextPageToken() *string {
	if o == nil {
		return nil
	}
	return o.NextPageToken
}

func (o *ListWorkflowExecutionsResponse) GetWorkflowExecutions() []WorkflowExecution {
	if o == nil {
		return nil
	}
	return o.WorkflowExecutions
}
