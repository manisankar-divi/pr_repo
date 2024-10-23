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
		// You can add more test cases here if needed
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

			// Check the status code
			if status := rr.Code; status != tt.status {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.status)
			}

			// Check the response body
			if rr.Body.String() != tt.expected {
				t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), tt.expected)
			}
		})
	}
}
