package common

import "context"

type Diffusion interface {
	Generate(DiffusionMessage) (DiffusionResult, error)
	GenerateWithContext(context.Context, DiffusionMessage) (DiffusionResult, error)
	GenerateWithOptions(context.Context, DiffusionOptions) (DiffusionResult, error)
}
