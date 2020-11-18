package constants

import (
	"fmt"
)

const (
	// httpPort is HTTP port
	httpPort = "3636"
	httpHost = "localhost"
)

// GetHTTPPort will return HTTP port with prefix ":"
func GetHTTPPort() string {
	return fmt.Sprintf(":%s", httpPort)
}

// GetAPIAddress will return HTTP port with prefix ":"
func GetAPIAddress() string {
	return fmt.Sprintf("http://%s:%s", httpHost, httpPort)
}
