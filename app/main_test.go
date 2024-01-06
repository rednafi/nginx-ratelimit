package main

import (
    "encoding/json"
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHelloWorldHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/greetings", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(helloWorldHandler)

    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    var response HelloWorldResponse
    err = json.Unmarshal(rr.Body.Bytes(), &response)
    if err != nil {
        t.Fatal(err)
    }

    expected := "Hello World"
    if response.Message != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", response.Message, expected)
    }

    ct := rr.Header().Get("Content-Type")
    if ct != "application/json" {
        t.Errorf("content type header does not match: got %v want %v", ct, "application/json")
    }
}
