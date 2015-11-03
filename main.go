package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var (
	Version = "bleeding-edge"

	version = flag.Bool("version", false, "print version and exit")
	addr    = flag.String("a", ":4242", "Address to bind to")
)

func init() {
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%v - ", time.Now().Format(time.RFC3339)))
}

func main() {
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	log.Println("Parrot server listening on", *addr)
	http.ListenAndServe(*addr, nil)
}
