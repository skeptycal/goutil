package seeker

import (
	"flag"
	"fmt"
	"net/http"
)

type (
	handler struct {
		endpoint string

		// http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))
		handlefunc func(http.ResponseWriter, *http.Request)
	}
)

var (
	port    = flag.Int("port", 8000, "port used by server")
	logFile = flag.String("log", "", "log file for server")
	outFile = flag.String("out", "", "output file")
)

var (
	PortString string = fmt.Sprintf("%v:%d", defaultServer, *port)
)
