package config

import (
	"os"
	"path"
)

func init() {
	if err := os.MkdirAll(Misc.TempDir, 0755); err != nil {
		panic(err)
	}
}

type MiscConfig struct {
	TempDir string
}

var Misc = &MiscConfig{
	TempDir: envOrDefault("CHLOE_TEMP_DIR", path.Join(os.TempDir(), "chloe")),
}
