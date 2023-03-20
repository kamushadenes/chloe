package react

import (
	"context"
	"fmt"
	"github.com/Knetic/govaluate"
	"github.com/kamushadenes/chloe/structs"
	"github.com/rs/zerolog"
	"strings"
)

func Calculate(ctx context.Context, request *structs.CalculationRequest) error {
	logger := zerolog.Ctx(ctx)

	expr := strings.ReplaceAll(request.Content, ",", "")

	logger.Info().Str("expression", expr).Msg("evaluating expression")

	expression, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		return NotifyError(request, err)
	}

	StartAndWait(request)

	result, err := expression.Evaluate(make(map[string]interface{}))
	if err != nil {
		return NotifyError(request, err)
	}

	if _, err := request.Writer.Write([]byte(fmt.Sprintf("%v", result))); err != nil {
		return NotifyError(request, err)
	}

	if !request.SkipClose {
		err := request.Writer.Close()
		return NotifyError(request, err)
	}

	return NotifyError(request, nil)
}
