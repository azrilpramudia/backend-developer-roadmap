package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleRoot(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello", nil)
	w := httptest.NewRecorder()

	handleRoot(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got %v\nbody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Welcome to our Server\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleGoodbye(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/goodbye", nil)
	w := httptest.NewRecorder()

	handleGoodbye(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Goodbye\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterized(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?user=TestMan", nil) 
	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad responese code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, TestMan!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterizedNoParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello/", nil) 
	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad responese code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloParameterizedWrongParam(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?foo=bar", nil) 
	w := httptest.NewRecorder()

	handleHelloParameterized(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad responese code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, User!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleUserResponsesHello(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/responses/TestMan/hello/", nil)
	req.SetPathValue("user", "TestMan")
	
	w := httptest.NewRecorder()

	handleUserResponsesHello(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got: %v\nbody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, TestMan!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.Bytes(), expectedMessage)
	}
}

func TestHandleHelloHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/hello/", nil)
	req.Header.Set("user", "Test Man")

	w := httptest.NewRecorder()

	handleHelloHeader(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad responese code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, Test Man!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleHelloNoHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/user/hello/", nil)

	w := httptest.NewRecorder()

	handleHelloHeader(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad responese code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("invalid username provided\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleJSON(t *testing.T) {
	testRequest := UserData{
		Name: "Test Man",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatalf("error marshalling test data: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(marshalledRequestBody))

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusOK
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("Hello, Test Man!\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleJSONEmptyBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/json", nil)

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("bad request body\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}

func TestHandleJSONEmptyNameField(t *testing.T) {
	testRequest := UserData{
		Name: "",
	}

	marshalledRequestBody, err := json.Marshal(testRequest)
	if err != nil {
		t.Fatalf("error marshaling test data: %v", err)
	}
	
	req := httptest.NewRequest(http.MethodPost, "/json", bytes.NewBuffer(marshalledRequestBody))

	w := httptest.NewRecorder()

	handleJSON(w, req)

	desiredCode := http.StatusBadRequest
	if w.Code != desiredCode {
		t.Errorf("bad response code, expected: %v but got %v\nBody: %s\n",
	desiredCode, w.Code, w.Body.String())
	}

	expectedMessage := []byte("invalid username provided\n")
	if !bytes.Equal(expectedMessage, w.Body.Bytes()) {
		t.Errorf("bad return, got: %q, expected: %q", w.Body.String(), expectedMessage)
	}
}