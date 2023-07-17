// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// TaskRevokeSourceExpired - The TaskRevokeSourceExpired message.
type TaskRevokeSourceExpired struct {
	ExpiredAt *time.Time `json:"expiredAt,omitempty"`
}

func (o *TaskRevokeSourceExpired) GetExpiredAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.ExpiredAt
}
