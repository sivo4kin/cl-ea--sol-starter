package http

import (
	"log"
	"net/http"
)

const addr = "localhost:12345"

func main() {
	mux := http.NewServeMux()
	handler := &MyHandler{}
	mux.Handle("/favicon.ico", http.NotFoundHandler())
	mux.Handle("/", handler)
	log.Printf("Now listening on %s...\n", addr)
	server := http.Server{Handler: mux, Addr: addr}
	log.Fatal(server.ListenAndServe())
}
