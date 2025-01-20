package cfg

import (
	"github.com/reiver/go-path"

	"github.com/reiver/ld-base/env"
)

func BasePathRoot() string {
	return path.Canonical(env.BasePathRoot)
}
