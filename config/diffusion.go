package config

type DiffusionProvider string

const (
	OpenAIDiffusion DiffusionProvider = "openai"
)

type DiffusionConfig struct {
	Provider DiffusionProvider
}

var Diffusion = &DiffusionConfig{
	Provider: DiffusionProvider(envOrDefaultWithOptions("CHLOE_DIFFUSION_PROVIDER", string(OpenAIDiffusion),
		[]string{string(OpenAIDiffusion)})),
}
