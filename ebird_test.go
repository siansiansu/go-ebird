package ebird

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestNewClient(t *testing.T) {
	client, err := NewClient("ebird_api_key")
	if err != nil {
		t.Fatalf("Error creating client: %v", err)
	}

	if client == nil {
		t.Fatal("Expected non-nil client, but got nil")
	}

	_, err = NewClient("")
	if err == nil {
		t.Fatal("Expected an error for empty API key, but got nil")
	}
}

func testClient(code int, body io.Reader) (*Client, *httptest.Server) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(code)
		_, _ = io.Copy(w, body)
	}))
	u, err := url.ParseRequestURI(server.URL + "/")
	if err != nil {
		panic(err)
	}
	client := &Client{
		httpClient: http.DefaultClient,
		baseURL:    u,
	}
	return client, server
}
