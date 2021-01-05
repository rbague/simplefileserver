package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	help := flag.Bool("help", false, "prints this help")
	dir := flag.StringP("directory", "d", ".", "the directory of which to serve files")
	host := flag.StringP("host", "h", "", "which host to bind the server to")
	port := flag.StringP("port", "p", "8080", "on which port to listen to")
	cert := flag.String("cert-file", "", "path to the SSL certificate")
	key := flag.String("key-file", "", "path to the SSL certificate's private key")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	address := fmt.Sprintf("%s:%s", *host, *port)
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(*dir)))

	var err error
	if *cert != "" || *key != "" {
		if *cert == "" || *key == "" {
			log.Fatal("both 'cert-file' and 'key-file' flags must present to enable HTTPS")
		}

		err = http.ListenAndServeTLS(address, *cert, *key, mux)
	} else {
		err = http.ListenAndServe(address, mux)
	}

	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("server closed unexpectedly: %v", err)
	}
}
