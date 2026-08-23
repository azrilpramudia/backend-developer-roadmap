package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	handleHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\nbody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, World!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleGoodbye(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/goodbye", nil)
	w := httptest.NewRecorder()

	handleGoodbye(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Goodbye\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected %q", w.Body.String(), expectedMessage)
	}
}

