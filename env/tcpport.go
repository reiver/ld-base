package env

import (
	"os"
)

const EnvVarNamePort string = "PORT"

var TcpPort string = tcpPort()

// tcpPort returns the TCP-port that the HTTP server should use.
//
// It defaults to TCP-port 8080.
//
// But that can be overridden by a value set in the "PORT" environment variable.
func tcpPort() string {
	tcpPort := os.Getenv(EnvVarNamePort)
	if "" == tcpPort {
		tcpPort = "8080"
	}

	return tcpPort
}
