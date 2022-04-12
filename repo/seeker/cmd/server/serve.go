package main

import (
	"fmt"
	http "net/http"

	"
)

type handler struct {
	endpoint    string
	handlerfunc http.HandlerFunc
}

var portString = seeker.PortString

var handlers []handler

func index_handler(w http.ResponseWriter, r *http.Request) {
	fake_test_page(w, r)
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fake_about_page(w, r)
}

func search_handler(w http.ResponseWriter, r *http.Request) {
	fake_search_page(w, r)
}

func init() {

	// config

	// setup end points

	// handlers :=
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/about/", about_handler)

}

func main() {
	cfg := seeker.Config(portString, handlers)
	seeker.Seek(portString)
}

func fake_test_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whoa, that's a BIG mushroom! (on port %v)\n", portString)
}

func fake_about_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seeker minimal web browser (on port %v)\n", portString)
}

func fake_search_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Seeker minimal web browser (on port %v)\n", portString)
}
