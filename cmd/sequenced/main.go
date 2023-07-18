package main

import (
	"net/http"
	"github.com/kzrl/sequence"
	"os"
	"fmt"
	"time"
	"log"
)


func main() {

	port := os.Getenv("SEQUENCED_PORT")
	if port == "" {
		port = ":8080"
	}
	

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// The "/" pattern matches everything, so we need to check
		// that we're at the root here.
		// snippted and comment shamelessly copypasted from https://pkg.go.dev/net/http#ServeMux
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}

		fmt.Fprintf(w, fmt.Sprintf("%d", sequence.Next()))	
	})

	s := &http.Server{
	Addr:           port,
	Handler:        mux,
	ReadTimeout:    10 * time.Second,
	WriteTimeout:   10 * time.Second,
	MaxHeaderBytes: 1 << 20,
	}
	log.Printf("Listening on port %s\n", port)
	log.Fatal(s.ListenAndServe())

}
