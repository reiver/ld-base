package verboten

import (
	"net/http"
	"strings"

	libpath "github.com/reiver/go-path"

	"github.com/reiver/ld-base/srv/http"
	"github.com/reiver/ld-base/srv/log"
)

const path string = "/data"

func init() {
	err := httpsrv.Mux.HandleDirectory(http.HandlerFunc(serveHTTP), path)
	if nil != err {
		panic(err)
	}
}

func serveHTTP(responsewriter http.ResponseWriter, request *http.Request) {
	log := logsrv.Prefix("www("+path+")").Begin()
	defer log.End()

	if nil == responsewriter {
		log.Error("nil response-writer")
		return
	}
	if nil == request {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil request")
		return
	}
	if nil == request.URL {
		const code int = http.StatusInternalServerError
		http.Error(responsewriter, http.StatusText(code), code)
		log.Error("nil request-url")
		return
	}

	var httpRequestPath string = libpath.Canonical(request.URL.Path)
	log.Debugf("http-request-path = %q", httpRequestPath)

	var basePath string
	{
		var prefix string = path
		if '/' != prefix[len(prefix)-1] {
			prefix = prefix + "/"
		}
		log.Debugf("path-prefix: %q", prefix)

		if !strings.HasPrefix(httpRequestPath, prefix) {
			const code int = http.StatusInternalServerError
			http.Error(responsewriter, http.StatusText(code), code)
			log.Errorf("http-request-path (%q) not prefixed by %q", httpRequestPath, prefix)
			return
		}

		basePath = httpRequestPath[len(prefix):]
	}
	log.Debugf("base-path = %q", basePath)

	var method string = request.Method
	log.Debugf("http-request-method = %q", method)

	switch method {
	case http.MethodGet:
		serveGET(responsewriter, basePath)
		return
//	case http.MethodDelete:
		
//	case http.MethodPatch:
		
//	case http.MethodPost:
		
	case http.MethodPut:
		servePUT(responsewriter, basePath, request.Body)
		return
	default:
		const code int = http.StatusMethodNotAllowed
		http.Error(responsewriter, http.StatusText(code), code)
		log.Debugf("method not supported: %q", method)
		return
	}
}
