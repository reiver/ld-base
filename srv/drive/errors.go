package drivesrv

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyDestination = erorr.Error("empty destination")
	errNilFileSystem    = erorr.Error("nil file-system")
	errNilSource        = erorr.Error("nil source")
)
