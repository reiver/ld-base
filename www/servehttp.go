package verboten

import (
	"io"
	"net/http"

	"github.com/reiver/ld-base/srv/http"
	"github.com/reiver/ld-base/srv/log"
)

const path string = "/"

func init() {
	err := httpsrv.Mux.HandlePath(http.HandlerFunc(serveHTTP), path)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	log := logsrv.Prefix("www("+path+")").Begin()
	defer log.End()

	if nil == responsewriter {
		return
	}
	if nil == request {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		return
	}

	io.WriteString(responsewriter, "<h1>ld-base</h1><p>A back-end with an embedded database for linked-data.</p>")
}
