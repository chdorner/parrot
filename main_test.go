package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParrotHandler(t *testing.T) {
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/foobar", nil)
	parrotHandler(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "/foobar\n", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/foobar.json", nil)
	parrotHandler(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "{\"url\":\"/foobar.json\"}\n", w.Body.String())

	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/foobar.xml", nil)
	parrotHandler(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "<parrot><url>/foobar.xml</url></parrot>\n", w.Body.String())
}

func TestStatusHandlers(t *testing.T) {
	w := httptest.NewRecorder()
	path, handler := statusHandler(200, "", "text/plain")
	req, _ := http.NewRequest("GET", path, nil)
	handler(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "text/plain", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "200 - OK\n", w.Body.String())

	w = httptest.NewRecorder()
	path, handler = statusHandler(404, ".json", "application/json")
	req, _ = http.NewRequest("GET", path, nil)
	handler(w, req)

	assert.Equal(t, 404, w.Code)
	assert.Equal(t, "application/json", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "{\"code\":404,\"text\":\"Not Found\"}\n", w.Body.String())

	w = httptest.NewRecorder()
	path, handler = statusHandler(503, ".xml", "application/xml")
	req, _ = http.NewRequest("GET", path, nil)
	handler(w, req)

	assert.Equal(t, 503, w.Code)
	assert.Equal(t, "application/xml", w.HeaderMap["Content-Type"][0])
	assert.Equal(t, "<status><code>503</code><text>Service Unavailable</text></status>\n", w.Body.String())
}
