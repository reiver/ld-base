package drivesrv

import (
	gofs "io/fs"
)

func Open(name string) (gofs.File, error) {
	if nil == fs {
		return nil, errNilFileSystem
	}

	return fs.Open(name)
}
