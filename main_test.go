package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestVisitURL is a test function that checks if the visitURL function works as expected
func TestVisitURL(t *testing.T) {
	// Create a fake HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Return a dummy response for testing
		_, _ = w.Write([]byte("Hello, World!"))
	}))
	defer ts.Close()

	// Call the visitURL function with the fake HTTP server URL
	resp, err := visitURL(ts.URL)
	if err != nil {
		t.Fatalf("error visiting URL: %v", err)
	}

	// Check if the response body size is as expected
	if resp.BodyLen != 13 {
		t.Fatalf("unexpected response body size: %d", resp.BodyLen)
	}
}

func TestProcessURLs(t *testing.T) {
	// Setup fake HTTP server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	}))
	defer ts.Close()

	urls := []string{ts.URL, ts.URL}
	responses, err := processURLs(urls)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := []Response{{URL: ts.URL, BodyLen: 11}, {URL: ts.URL, BodyLen: 11}}
	if len(responses) != len(expected) {
		t.Fatalf("expected %d responses, got %d", len(expected), len(responses))
	}
	for i, resp := range responses {
		if resp != expected[i] {
			t.Fatalf("expected %v, got %v", expected[i], resp)
		}
	}
}
