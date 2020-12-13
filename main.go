package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

func main() {
	help := flag.Bool("help", false, "prints this help")
	dir := flag.StringP("directory", "d", ".", "the directory of which to serve files")
	port := flag.StringP("port", "p", "8080", "on which port to listen to")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

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
