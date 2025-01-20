package drivesrv

import (
	gofs "io/fs"
	"os"

	"github.com/reiver/ld-base/cfg"
)

var fs gofs.FS

func init () {
	fs = os.DirFS(Path())
	if nil == fs {
		panic("nil fs.FS for root of base")
	}
}

func FS() gofs.FS {
	return fs
}

func Path() string {
	return cfg.BasePathRoot()
}
