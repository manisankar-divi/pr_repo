package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestHelloHandler tests the helloHandler function.
func TestHelloHandler(t *testing.T) {
	tests := []struct {
		name     string
		method   string
		expected string
		status   int
	}{
		{"Valid Request", "GET", "Hello, World!", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest(tt.method, "/", nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(helloHandler)

			handler.ServeHTTP(rr, req)

			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.status)
			}

			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
			}
		})
	}
}

// TestStartServer tests the startServer function.
func TestStartServer(t *testing.T) {
	// Set up a test server
	ts := httptest.NewServer(http.HandlerFunc(helloHandler))
	defer ts.Close()

	res, err := http.Get(ts.URL)
	if err != nil {
		t.Fatalf("could not make GET request: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", res.StatusCode)
	}
}
