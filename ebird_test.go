package ebird

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	t.Run("Valid API Key", func(t *testing.T) {
		client, err := NewClient("valid_api_key")
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, "valid_api_key", client.apikey)
		assert.Equal(t, APIEndpointBase, client.baseURL.String())
	})

	t.Run("Empty API Key", func(t *testing.T) {
		client, err := NewClient("")
		assert.Error(t, err)
		assert.Nil(t, client)
		assert.Contains(t, err.Error(), "eBird API key is missing")
	})

	t.Run("Custom Base URL", func(t *testing.T) {
		customURL := "https://custom.ebird.api/"
		client, err := NewClient("valid_api_key", WithBaseURL(customURL))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, customURL, client.baseURL.String())
	})

	t.Run("Custom HTTP Client", func(t *testing.T) {
		customClient := &http.Client{Timeout: 5 * time.Second}
		client, err := NewClient("valid_api_key", WithHTTPClient(customClient))
		assert.NoError(t, err)
		assert.NotNil(t, client)
		assert.Equal(t, customClient, client.httpClient)
	})
}

func TestClientGet(t *testing.T) {
	t.Run("Successful Request", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			assert.Equal(t, "GET", r.Method)
			assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
			assert.Equal(t, "test_api_key", r.Header.Get("X-eBirdApiToken"))
			json.NewEncoder(w).Encode(map[string]string{"key": "value"})
		}))
		defer server.Close()

		client, err := NewClient("test_api_key", WithBaseURL(server.URL+"/"))
		require.NoError(t, err)

		var result map[string]string
		err = client.get(context.Background(), "test", nil, &result)
		assert.NoError(t, err)
		assert.Equal(t, map[string]string{"key": "value"}, result)
	})

	t.Run("Error Response", func(t *testing.T) {
		server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]interface{}{
				"error": map[string]interface{}{
					"message": "Bad Request",
					"status":  400,
				},
			})
		}))
		defer server.Close()

		client, err := NewClient("test_api_key", WithBaseURL(server.URL+"/"))
		require.NoError(t, err)

		var result map[string]string
		err = client.get(context.Background(), "test", nil, &result)
		assert.Error(t, err)
		apiErr, ok := err.(Error)
		assert.True(t, ok)
		assert.Equal(t, "Bad Request", apiErr.Message)
		assert.Equal(t, 400, apiErr.Status)
	})
}

func TestConvertToJsonFormat(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "URL without parameters",
			input:    "https://api.ebird.org/v2/data/obs/US-NY/recent",
			expected: "https://api.ebird.org/v2/data/obs/US-NY/recent?fmt=json",
		},
		{
			name:     "URL with existing parameters",
			input:    "https://api.ebird.org/v2/data/obs/US-NY/recent?back=30",
			expected: "https://api.ebird.org/v2/data/obs/US-NY/recent?back=30&fmt=json",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := convertToJsonFormat(tt.input)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestClientWithAcceptLanguage(t *testing.T) {
	client, err := NewClient("test_api_key", WithAcceptLanguage("es"))
	require.NoError(t, err)
	assert.Equal(t, "es", client.acceptLanguage)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "es", r.Header.Get("Accept-Language"))
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Hola, mundo!"})
	}))
	defer server.Close()

	client.baseURL, _ = client.baseURL.Parse(server.URL + "/")

	var result map[string]string
	err = client.get(context.Background(), "test", nil, &result)
	assert.NoError(t, err)
	assert.Equal(t, "Hola, mundo!", result["message"])
}
