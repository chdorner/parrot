package main

import (
	"encoding/json"
	"encoding/xml"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

var (
	Version = "bleeding-edge"

	version = flag.Bool("version", false, "print version and exit")
	addr    = flag.String("a", ":4242", "Address to bind to")
)

type Parrot struct {
	XMLName int    `json:"-" xml:"parrot"`
	URL     string `json:"url" xml:"url"`
}

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

	http.HandleFunc("/", parrotHandler)

	log.Println("Parrot server listening on", *addr)
	http.ListenAndServe(*addr, nil)
}

func parrotHandler(w http.ResponseWriter, r *http.Request) {
	resp := Parrot{URL: r.URL.String()}
	switch {
	case strings.HasSuffix(r.URL.Path, ".json"):
		w.Header().Set("Content-Type", "application/json")
		b, _ := json.Marshal(resp)
		fmt.Fprintf(w, "%s\n", b)
	case strings.HasSuffix(r.URL.Path, ".xml"):
		w.Header().Set("Content-Type", "application/xml")
		b, _ := xml.Marshal(resp)
		fmt.Fprintf(w, "%s\n", b)
	default:
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "%s\n", r.URL.String())
	}
}
