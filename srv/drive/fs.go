package drivesrv

import (
	gofs "io/fs"
	"os"

	"github.com/reiver/ld-base/cfg"
)

var fs gofs.FS

func init () {
	fs = os.DirFS(OvertRoot())
	if nil == fs {
		panic("nil fs.FS for root of base")
	}
}

func OvertRoot() string {
	return cfg.BasePathRoot()
}
