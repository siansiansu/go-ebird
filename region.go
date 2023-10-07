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
	Bounds    Bounds
	Result    string  `json:"result,omitempty"`
	Code      string  `json:"code,omitempty"`
	Type      string  `json:"type,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type SubRegionList struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

func (c *Client) RegionInfo(ctx context.Context, regionCode string, opts ...RequestOption) (*RegionInfo, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndointRegionInfo, regionCode)

	var t RegionInfo
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}
	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Client) SubRegionList(ctx context.Context, regionType, parentRegionCode string, opts ...RequestOption) ([]SubRegionList, error) {
	if regionType == "" {
		return nil, fmt.Errorf("regionType cannot be empty")
	}
	if parentRegionCode == "" {
		return nil, fmt.Errorf("parentRegionCode cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndointSubRegionInfo, regionType, parentRegionCode)

	var t []SubRegionList
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
