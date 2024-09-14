package ebird

import (
	"context"
	"fmt"
	"net/url"
)

type AdjacentRegion struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

type AdjacentRegionsOptions struct {
	fmt string
}

type AdjacentRegionsOption func(*AdjacentRegionsOptions)

func WithFormat(format string) AdjacentRegionsOption {
	return func(o *AdjacentRegionsOptions) {
		o.fmt = format
	}
}

func (c *Client) AdjacentRegions(ctx context.Context, regionCode string, opts ...AdjacentRegionsOption) ([]AdjacentRegion, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	options := &AdjacentRegionsOptions{}
	for _, opt := range opts {
		opt(options)
	}

	params := url.Values{}
	if options.fmt != "" {
		params.Set("fmt", options.fmt)
	}

	endpoint := fmt.Sprintf(APIEndpoints.AdjacentRegions, regionCode)

	var regions []AdjacentRegion
	err := c.get(ctx, endpoint, params, &regions)
	if err != nil {
		return nil, fmt.Errorf("failed to get adjacent regions: %w", err)
	}

	return regions, nil
}
