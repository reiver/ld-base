package verboten

import (
	"io"
	"net/http"

	"github.com/reiver/ld-base/srv/drive"
	"github.com/reiver/ld-base/srv/log"
)

func servePUT(responsewriter http.ResponseWriter, basepath string, body io.ReadCloser) {
	log := logsrv.Prefix("www(" + path + ").PUT").Begin()
	defer log.End()

	if nil == responsewriter {
		log.Error("nil response-writer")
		return
	}
	if nil == body {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil request-body")
		return
	}

	{
		err := drivesrv.WriteFrom(basepath, body)
		if nil != err {
			const code int = http.StatusInternalServerError
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("problem writing http-request-body to base: %s", err)
			return
		}
	}

	{
		const code int = http.StatusCreated
		responsewriter.WriteHeader(code)
		_, err := io.WriteString(responsewriter, "true")
		if nil != err {
			log.Errorf("problem writing content to client: %s", err)
		}
		return
	}
}
