package whispercpp

import (
	"github.com/kamushadenes/chloe/langchain/audio_models/common"
	"github.com/kamushadenes/chloe/langchain/schema"
)

var (
	Base = &common.ASRModel{
		Name:        "base",
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0, Unit: schema.CostUnitMinute, UnitSize: 1},
	}

	Tiny = &common.ASRModel{
		Name:        "tiny",
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0, Unit: schema.CostUnitMinute, UnitSize: 1},
	}

	Small = &common.ASRModel{
		Name:        "small",
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0, Unit: schema.CostUnitMinute, UnitSize: 1},
	}

	Medium = &common.ASRModel{
		Name:        "medium",
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0, Unit: schema.CostUnitMinute, UnitSize: 1},
	}

	Large = &common.ASRModel{
		Name:        "large",
		ContextSize: 26214400, // 25MB
		ContextUnit: schema.ContextUnitByte,
		UsageCost:   &schema.CostObject{Price: 0, Unit: schema.CostUnitMinute, UnitSize: 1},
	}
)
