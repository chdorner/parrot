package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusHandlers(t *testing.T) {
	handler := StatusHandler{200}
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/_/200", nil)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/plain", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "200 - OK\n", w.Body.String())

	handler = StatusHandler{404}
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/_/404.json", nil)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "application/json", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "{\"code\":404,\"text\":\"Not Found\"}\n", w.Body.String())

	handler = StatusHandler{503}
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/_/503.xml", nil)
	handler.ServeHTTP(w, req)

	assert.Equal(t, 503, w.Code)
	assert.Equal(t, "application/xml", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "<status><code>503</code><text>Service Unavailable</text></status>\n", w.Body.String())
}
