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

	statuses = [...]int{
		200, 201, 202, 203, 204, 205, 206,
		300, 301, 302, 303, 304, 305, 307,
		400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418,
		500, 501, 502, 503, 504, 505,
	}
)

type Parrot struct {
	XMLName int    `json:"-" xml:"parrot"`
	URL     string `json:"url" xml:"url"`
}

type Status struct {
	XMLName int    `json:"-" xml:"status"`
	Code    int    `json:"code" xml:"code"`
	Text    string `json:"text" xml:"text"`
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
	for _, s := range statuses {
		http.HandleFunc(statusHandler(s, "", "text/plain"))
		http.HandleFunc(statusHandler(s, ".json", "application/json"))
		http.HandleFunc(statusHandler(s, ".xml", "application/xml"))
	}

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

func statusHandler(status int, ext string, ctype string) (string, func(http.ResponseWriter, *http.Request)) {
	pattern := fmt.Sprintf("/_/%d%v", status, ext)
	return pattern, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(status)
		w.Header().Set("Content-Type", ctype)

		resp := Status{Code: status, Text: http.StatusText(status)}
		switch ext {
		case "":
			fmt.Fprintf(w, "%d - %v\n", resp.Code, resp.Text)
		case ".json":
			b, _ := json.Marshal(resp)
			fmt.Fprintf(w, "%s\n", b)
		case ".xml":
			b, err := xml.Marshal(resp)
			if err != nil {
				log.Println(err)
			}
			fmt.Fprintf(w, "%s\n", b)
		}
	}
}
