package schema

type ChatGeneration struct {
	Text    string
	Message Message
}

type ChatUsage struct {
	PromptTokens     int
	CompletionTokens int
	TotalTokens      int
}

type ChatResult struct {
	Generations []ChatGeneration
	Usage       ChatUsage
}
