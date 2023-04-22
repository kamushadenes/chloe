package common

import "time"

type DiffusionOptions interface {
	GetRequest() interface{}
	WithPrompt(string) DiffusionOptions
	WithModel(*DiffusionModel) DiffusionOptions
	GetModel() *DiffusionModel
	WithCount(int) DiffusionOptions
	WithFormat(string) DiffusionOptions
	WithTimeout(time.Duration) DiffusionOptions
	GetTimeout() time.Duration
}
