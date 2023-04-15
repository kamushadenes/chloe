package schema

type CostUnit string

const (
	Token  CostUnit = "token"
	Image  CostUnit = "image"
	Minute CostUnit = "minute"
)

type CostObject struct {
	Price    float64
	Unit     CostUnit
	UnitSize int
}
