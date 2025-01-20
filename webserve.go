package main

import (
	"net/http"

	"github.com/reiver/ld-base/cfg"
	"github.com/reiver/ld-base/srv/http"
	"github.com/reiver/ld-base/srv/log"
	_ "github.com/reiver/ld-base/www" // This import enables all the HTTP handlers.
)

func webserve() {
	log := logsrv.Prefix("webserve")

	var tcpaddr string = cfg.WebServerTCPAddress()
	log.Informf("serving HTTP on TCP address: %q", tcpaddr)
	log.Informf("NOTE that TCP port of TCP address can be overridden using %q environment variable", cfg.EnvVarNamePort)

	err := http.ListenAndServe(tcpaddr, &httpsrv.Mux)
	if nil != err {
		log.Errorf("ERROR: problem with serving HTTP on TCP address %q: %s", tcpaddr, err)
		panic(err)
	}
}
