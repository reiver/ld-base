package env

import (
	"os"
)

var BasePathRoot string = basePathRoot()

func basePathRoot() string {
	base := os.Getenv("BASE")
	if "" == base {
		base = "."
	}

	return base
}
