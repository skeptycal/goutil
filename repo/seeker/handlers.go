package seeker

import (
	"fmt"
	"net/http"
	"os"
)

// http.HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request))

// simple argument server
func ArgServer(w http.ResponseWriter, req *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, " ")
	}
}
