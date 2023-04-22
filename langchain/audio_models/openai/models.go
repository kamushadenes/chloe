package openai

import (
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/kamushadenes/chloe/langchain/schema"
	"github.com/sashabaranov/go-openai"
)

var (
	Whisper1 = &common.ASRModel{
		Name:        openai.Whisper1,
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0.005, Unit: schema.CostUnitMinute, UnitSize: 1},
	}
)
