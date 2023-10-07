package ebird

import (
	"context"
	"fmt"
	"strings"
)

type AdjacentRegions struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) AdjacentRegions(ctx context.Context, regionCode string, opts ...RequestOption) ([]AdjacentRegions, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndpointAdjacentRegions, regionCode)
	if !strings.Contains(ebirdURL, "ref/adjacent") {
		return nil, fmt.Errorf("invalid API endpoint for AdjacentRegions")
	}

	var t []AdjacentRegions
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, fmt.Errorf("failed to make GET request: %v", err)
	}
	return t, nil
}
