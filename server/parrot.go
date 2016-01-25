package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

type ParrotHandler struct{}

type ParrotResponse struct {
	XMLName int    `json:"-" xml:"parrot"`
	URL     string `json:"url" xml:"url"`
}

func (h *ParrotHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := ParrotResponse{URL: r.URL.String()}
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
