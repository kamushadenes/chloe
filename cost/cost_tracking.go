package cost

var sessionCost = map[string]float64{
	"completion":       0,
	"chain_of_thought": 0,
	"transcription":    0,
	"moderation":       0,
	"summarization":    0,
	"image_generation": 0,
	"image_edit":       0,
	"image_variation":  0,
}

func GetCategoryCost(category string) float64 {
	if _, ok := sessionCost[category]; !ok {
		return 0
	}
	return sessionCost[category]
}

func GetCategories() []string {
	var categories []string

	for k := range sessionCost {
		categories = append(categories, k)
	}

	return categories
}

func GetTotalSessionCost() float64 {
	var total float64
	for _, cost := range sessionCost {
		total += cost
	}
	return total
}

func AddCategoryCost(category string, cost float64) {
	if _, ok := sessionCost[category]; !ok {
		sessionCost[category] = 0
	}
	sessionCost[category] += cost
}

func ResetSessionCost() {
	sessionCost = make(map[string]float64)
}

func SubtractCategoryCost(category string, cost float64) {
	if _, ok := sessionCost[category]; !ok {
		sessionCost[category] = 0
	}
	sessionCost[category] -= cost
}
