package main

import (
	"flag"
	"log"
	"net/http"
	"path/filepath"
)

func main() {
	dir := flag.String("dir", ".", "the directory of which to serve files")
	port := flag.String("port", "8080", "on which port to listen to")
	flag.Parse()

	directory, err := filepath.Abs(*dir)
	if err != nil {
		log.Fatalf("could not get absolute path from %q: %v", *dir, err)
	}

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(directory)))

	if err := http.ListenAndServe(":"+*port, mux); err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}
