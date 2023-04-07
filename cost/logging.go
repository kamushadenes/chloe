package cost

import "github.com/kamushadenes/chloe/logging"

func logCost(msg string) {
	logger := logging.GetLogger()

	l := logger.Info()

	for _, k := range GetCategories() {
		l = l.Float64(k, GetCategoryCost(k))
	}

	l = l.Float64("total", GetTotalSessionCost())

	l.Msg(msg)
}
