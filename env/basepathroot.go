package env

import (
	"os"
)

const EnvVarNameBase string = "BASE"

var BasePathRoot string = basePathRoot()

func basePathRoot() string {
	base := os.Getenv(EnvVarNameBase)
	if "" == base {
		base = "."
	}

	return base
}
