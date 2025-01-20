package cfg

import (
	"fmt"

	"github.com/reiver/ld-base/env"
)

func WebServerTCPAddress() string {
	return fmt.Sprintf(":%s", env.TcpPort)
}
