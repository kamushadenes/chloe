package config

import (
	"os"
	"path/filepath"
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
	WorkspaceDir         string
}

var Misc = &MiscConfig{
	TempDir:              envOrDefault("CHLOE_TEMP_DIR", filepath.Join(os.TempDir(), "chloe")),
	CostTrackingInterval: envOrDefaultDuration("CHLOE_COST_TRACKING_INTERVAL", 5*time.Minute),
	WorkspaceDir: envOrDefault("CHLOE_WORKSPACE_DIR", filepath.Join(func() string {
		if wd, err := os.Getwd(); err == nil {
			return wd
		}
		return os.TempDir()
	}(), "..", "..", "workspace")),
}
