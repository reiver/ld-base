package verboten

import (
	"io"
	"io/fs"
	"net/http"

	"github.com/reiver/go-erorr"

	"github.com/reiver/ld-base/srv/drive"
	"github.com/reiver/ld-base/srv/log"
)

func serveGET(responsewriter http.ResponseWriter, basepath string) {
	log := logsrv.Prefix("www("+path+").GET").Begin()
	defer log.End()

	if nil == responsewriter {
		log.Error("nil response-writer")
		return
	}

	var file fs.File
	{
		var err error
		file, err = drivesrv.Open(basepath)
		if nil != err {
			switch {
			case erorr.Is(err, fs.ErrNotExist):
				const code int = http.StatusNotFound
				http.Error(responsewriter, http.StatusText(code), code)
				log.Errorf("problem opening file from %q: %s", basepath, err)
				return
			default:
				const code int = http.StatusInternalServerError
				http.Error(responsewriter, http.StatusText(code), code)
				log.Errorf("problem opening file from %q: %s", basepath, err)
				return
			}
		}
	}
	defer func() {
		err := file.Close()
		if nil != err {
			log.Errorf("problem closing file from %q: %s", basepath, err)
		}
	}()

	{
		_, err := io.Copy(responsewriter, file)
		if nil != err {
			log.Errorf("problem sending file from %q to client: %s", basepath, err)
		}
	}
}
