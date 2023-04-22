package structs

import "github.com/kamushadenes/chloe/langchain/memory"

type Action interface {
	// GetName Get the name of the action
	GetName() string

	// GetNotification  Get the notification message that will be sent when the action is executed
	GetNotification() string

	// SetParam Set the parameters of the action
	SetParam(string, string)

	// GetParam Get the value of a parameter
	GetParam(string) (string, error)

	// GetParams Get all parameters
	GetParams() map[string]string

	// CheckRequiredParams Check if all required parameters are set
	CheckRequiredParams() error

	// SetMessage Set the message that triggered the action
	SetMessage(*memory.Message)

	// Execute The actual implementation of the action
	Execute(*ActionRequest) ([]*ResponseObject, error)

	// RunPreActions Actions that will be executed before the action is executed
	RunPreActions(*ActionRequest) error

	// RunPostActions Actions that will be executed after the action is executed
	RunPostActions(*ActionRequest) error
}
