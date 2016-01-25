package server

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"net/http"
	"strings"
)

var (
	Statuses = [...]int{
		200, 201, 202, 203, 204, 205, 206,
		300, 301, 302, 303, 304, 305, 307,
		400, 401, 402, 403, 404, 405, 406, 407, 408, 409, 410, 411, 412, 413, 414, 415, 416, 417, 418,
		500, 501, 502, 503, 504, 505,
	}
)

type StatusHandler struct {
	Status int
}

type StatusResponse struct {
	XMLName int    `json:"-" xml:"status"`
	Code    int    `json:"code" xml:"code"`
	Text    string `json:"text" xml:"text"`
}

func (h *StatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(h.Status)
	resp := StatusResponse{Code: h.Status, Text: http.StatusText(h.Status)}
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
		fmt.Fprintf(w, "%d - %s\n", resp.Code, resp.Text)
	}
}
