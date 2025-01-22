package drivesrv

import (
	"fmt"
	"io"
	"os"
	"time"

	"github.com/reiver/go-erorr"

	libpath "github.com/reiver/go-path"

	"github.com/reiver/ld-base/srv/log"
)

func WriteFrom(dst string, src io.ReadCloser) error {
	log := logsrv.Prefix("drivesrv.WriteFrom").Begin()
	defer log.End()

	if "" == dst {
		return errEmptyDestination
	}
	if nil == src {
		return errNilSource
	}

	{
		log.Debugf("temp-file should be created in directory: %q", os.TempDir())
	}

	var file *os.File
	{
		now := time.Now()

		var pattern string = fmt.Sprintf("ld-base_%d-%02d-%02d_*.tmp", now.Year(), now.Month(), now.Day())
		log.Debugf("temp-file pattern: %q", pattern)

		var err error
		file, err = os.CreateTemp("", pattern)
		if nil != err {
			log.Errorf("problem creating temp-file: %s", err)
			return err
		}
		defer func() {
			err := file.Sync()
			if nil != err {
				log.Errorf("problem (deferred) syncing temp-file from %q: %s", dst, err)
				return
			}

			err := file.Close()
			if nil != err && !erorr.Is(err, os.ErrClosed) {
				log.Errorf("problem (deferred) closing temp-file from %q: %s", dst, err)
			}
		}()
	}

	var temppath string = file.Name()
	log.Debugf("temp-file file-name: %q", temppath)

	{
		n, err := io.Copy(file, src)
		if nil != err {
			log.Errorf("problem copying src (probably a request-body) to temp-file: %s", err)
			return err
		}

		log.Debugf("%d bytes copied from src (probably a request-body) to temp-file %q", n, temppath)
	}

	{
		err := file.Sync()
		if nil != err {
			log.Errorf("problem syncing temp-file from %q: %s", dst, err)
			return
		}

		err := file.Close()
		if nil != err {
			log.Errorf("problem closing temp-file from %q: %s", dst, err)
			return err
		}
	}

	var newpath string
	{
		newpath = libpath.Join(OvertRoot(), dst)
		newpath = libpath.Canonical(newpath)

		log.Debugf("destination-path: %q", newpath)
	}

	{
		err := os.Rename(temppath, newpath)
		if nil != err {
			log.Errorf("problem moving (i.e., renameing) the temp-file %q to %q: %s", temppath, newpath, err)
			return err
		}
	}

	return nil
}
