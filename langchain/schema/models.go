package schema

type Model struct {
	Name             string
	ContextSize      int
	ContextUnit      ContextUnit
	TokensPerMessage int
	TokensPerName    int
	UsageCost        *CostObject
	PromptCost       *CostObject
	CompletionCost   *CostObject
	Tokenizer        string
}
