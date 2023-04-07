package config

import (
	"os"
	"path"
	"time"
)

func init() {
	if err := os.MkdirAll(Misc.TempDir, 0755); err != nil {
		panic(err)
	}
}

type MiscConfig struct {
	TempDir              string
	CostTrackingInterval time.Duration
}

var Misc = &MiscConfig{
	TempDir:              envOrDefault("CHLOE_TEMP_DIR", path.Join(os.TempDir(), "chloe")),
	CostTrackingInterval: envOrDefaultDuration("CHLOE_COST_TRACKING_INTERVAL", 5*time.Minute),
}
