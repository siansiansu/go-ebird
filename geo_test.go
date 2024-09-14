package ebird

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAdjacentRegions(t *testing.T) {
	tests := []struct {
		name       string
		regionCode string
		response   []AdjacentRegion
		format     string
		expectErr  bool
	}{
		{
			name:       "Valid request without format",
			regionCode: "US",
			response: []AdjacentRegion{
				{Code: "US", Name: "United States"},
				{Code: "CA", Name: "Canada"},
			},
			expectErr: false,
		},
		{
			name:       "Valid request with format",
			regionCode: "US",
			response: []AdjacentRegion{
				{Code: "US", Name: "United States"},
				{Code: "CA", Name: "Canada"},
			},
			format:    "json",
			expectErr: false,
		},
		{
			name:       "Empty region code",
			regionCode: "",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Equal(t, tt.format, r.URL.Query().Get("fmt"))

				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(tt.response)
			}))
			defer server.Close()

			client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
			require.NoError(t, err)

			ctx := context.Background()
			var opts []AdjacentRegionsOption
			if tt.format != "" {
				opts = append(opts, WithFormat(tt.format))
			}

			regions, err := client.AdjacentRegions(ctx, tt.regionCode, opts...)

			if tt.expectErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.response, regions)
			}
		})
	}
}
