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
