package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/chdorner/parrot/server"
)

var (
	Version = "bleeding-edge"

	version   = flag.Bool("version", false, "print version and exit")
	addr      = flag.String("a", ":4242", "Address to bind to")
	directory = flag.String("dir", "", "Serve a local directory (disables other features)")
)

func init() {
	flag.Parse()

	log.SetFlags(0)
	log.SetPrefix(fmt.Sprintf("%v - ", time.Now().Format(time.RFC3339)))
}

func main() {
	sigch := make(chan os.Signal)
	go handleSignals(sigch)
	signal.Notify(sigch)

	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if *directory == "" {
		http.Handle("/", &server.ParrotHandler{})
		for _, s := range server.Statuses {
			pattern := fmt.Sprintf("/_/%d", s)
			h := &server.StatusHandler{s}

			http.Handle(pattern, h)
			http.Handle(fmt.Sprintf("%s.json", pattern), h)
			http.Handle(fmt.Sprintf("%s.xml", pattern), h)
		}
	} else {
		http.Handle("/", http.FileServer(http.Dir(*directory)))
	}

	log.Println("Parrot server listening on", *addr)
	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handleSignals(sigch chan os.Signal) {
	for sig := range sigch {
		switch sig {
		case syscall.SIGTERM, syscall.SIGINT:
			os.Exit(0)
		}
	}
}
