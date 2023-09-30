package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type AdjacentRegionsResponse struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) AdjacentRegions(ctx context.Context, regionCode string) ([]AdjacentRegionsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointAdjacentRegions, regionCode)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	return decodeToAdjacentRegions(req)
}

func decodeToAdjacentRegions(res *http.Response) ([]AdjacentRegionsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []AdjacentRegionsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
