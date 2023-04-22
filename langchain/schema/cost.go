package schema

type CostUnit string
type ContextUnit string

const (
	CostUnitToken  CostUnit = "token"
	CostUnitImage  CostUnit = "image"
	CostUnitMinute CostUnit = "minute"

	ContextUnitToken ContextUnit = "token"
	ContextUnitByte  ContextUnit = "byte"
	ContextUnitChar  ContextUnit = "char"
)

type CostObject struct {
	Price    float64
	Unit     CostUnit
	UnitSize int
}
