package flags

import (
	_ "embed"
)

var InteractiveCLI bool

//go:generate bash ../.github/scripts/get_version.sh
//go:embed version.txt
var Version string
