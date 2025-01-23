// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"github.com/conductorone/conductorone-sdk-go/pkg/utils"
)

// TaskActionsServiceProcessNowResponseExpanded - Contains an arbitrary serialized message along with a @type that describes the type of the serialized message.
type TaskActionsServiceProcessNowResponseExpanded struct {
	// The type of the serialized message.
	AtType               *string        `json:"@type,omitempty"`
	AdditionalProperties map[string]any `additionalProperties:"true" json:"-"`
}

func (t TaskActionsServiceProcessNowResponseExpanded) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(t, "", false)
}

func (t *TaskActionsServiceProcessNowResponseExpanded) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &t, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *TaskActionsServiceProcessNowResponseExpanded) GetAtType() *string {
	if o == nil {
		return nil
	}
	return o.AtType
}

func (o *TaskActionsServiceProcessNowResponseExpanded) GetAdditionalProperties() map[string]any {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

// The TaskActionsServiceProcessNowResponse message.
type TaskActionsServiceProcessNowResponse struct {
	// Contains a task and JSONPATH expressions that describe where in the expanded array related objects are located. This view can be used to display a fully-detailed dashboard of task information.
	TaskView *TaskView `json:"taskView,omitempty"`
	// The expanded field.
	Expanded []TaskActionsServiceProcessNowResponseExpanded `json:"expanded,omitempty"`
}

func (o *TaskActionsServiceProcessNowResponse) GetTaskView() *TaskView {
	if o == nil {
		return nil
	}
	return o.TaskView
}

func (o *TaskActionsServiceProcessNowResponse) GetExpanded() []TaskActionsServiceProcessNowResponseExpanded {
	if o == nil {
		return nil
	}
	return o.Expanded
}
