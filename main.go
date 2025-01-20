package main

import (
	"github.com/reiver/ld-base/srv/log"
)

func main() {
	log := logsrv.Prefix("main")

	log.Inform("ld-base ⚡")
	blur()

	//log.Inform("Here we go…")
	//webserve()
}
