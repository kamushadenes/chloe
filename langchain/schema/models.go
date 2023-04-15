package schema

type Model struct {
	Name             string
	ContextSize      int
	TokensPerMessage int
	TokensPerName    int
	UsageCost        *CostObject
	PromptCost       *CostObject
	CompletionCost   *CostObject
}
