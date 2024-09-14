package ebird

import (
	"context"
	"fmt"
)

type Bounds struct {
	MinX float64 `json:"minX,omitempty"`
	MaxX float64 `json:"maxX,omitempty"`
	MinY float64 `json:"minY,omitempty"`
	MaxY float64 `json:"maxY,omitempty"`
}

type RegionInfo struct {
	Bounds    Bounds  `json:"bounds"`
	Result    string  `json:"result,omitempty"`
	Code      string  `json:"code,omitempty"`
	Type      string  `json:"type,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type SubRegion struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) RegionInfo(ctx context.Context, regionCode string, opts ...RequestOption) (*RegionInfo, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.RegionInfo, regionCode)
	params := processOptions(opts...)

	var info RegionInfo
	err := c.get(ctx, endpoint, params.URLParams, &info)
	if err != nil {
		return nil, fmt.Errorf("failed to get region info: %w", err)
	}

	return &info, nil
}

func (c *Client) SubRegionList(ctx context.Context, regionType, parentRegionCode string, opts ...RequestOption) ([]SubRegion, error) {
	if regionType == "" {
		return nil, fmt.Errorf("regionType cannot be empty")
	}
	if parentRegionCode == "" {
		return nil, fmt.Errorf("parentRegionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.SubRegionInfo, regionType, parentRegionCode)
	params := processOptions(opts...)

	var subRegions []SubRegion
	err := c.get(ctx, endpoint, params.URLParams, &subRegions)
	if err != nil {
		return nil, fmt.Errorf("failed to get subregion list: %w", err)
	}

	return subRegions, nil
}
