package ebird

import (
	"context"
	"fmt"
)

type HotspotInfo struct {
	LocId            string  `json:"locId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Latitude         float64 `json:"latitude,omitempty"`
	Longitude        float64 `json:"longitude,omitempty"`
	CountryCode      string  `json:"countryCode,omitempty"`
	CountryName      string  `json:"countryName,omitempty"`
	Subnational1Name string  `json:"subnational1Name,omitempty"`
	Subnational1Code string  `json:"subnational1Code,omitempty"`
	Subnational2Code string  `json:"subnational2Code,omitempty"`
	Subnational2Name string  `json:"subnational2Name,omitempty"`
	IsHotspot        bool    `json:"isHotspot,omitempty"`
	LocID            string  `json:"locID,omitempty"`
	LocName          string  `json:"locName,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
}

type NearbyHotspot struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Lat               float64 `json:"lat"`
	Lng               float64 `json:"lng"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime int     `json:"numSpeciesAllTime,omitempty"`
}

type HotspotInRegion struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Subnational2Code  string  `json:"subnational2Code,omitempty"`
	Lat               float64 `json:"lat,omitempty"`
	Lng               float64 `json:"lng,omitempty"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime int     `json:"numSpeciesAllTime,omitempty"`
}

func (c *Client) HotspotsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]HotspotInRegion, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.HotspotsInRegion, regionCode)
	params := processOptions(opts...)

	var hotspots []HotspotInRegion
	err := c.get(ctx, endpoint, params.URLParams, &hotspots)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotspots in region: %w", err)
	}

	return hotspots, nil
}

func (c *Client) NearbyHotspots(ctx context.Context, opts ...RequestOption) ([]NearbyHotspot, error) {
	params := processOptions(opts...)

	params.URLParams.Set("fmt", "json")

	if params.URLParams.Get("lat") == "" || params.URLParams.Get("lng") == "" {
		return nil, fmt.Errorf("must provide Lat and Lng parameters")
	}

	var hotspots []NearbyHotspot
	err := c.get(ctx, APIEndpoints.NearbyHotspots, params.URLParams, &hotspots)
	if err != nil {
		return nil, fmt.Errorf("failed to get nearby hotspots: %w", err)
	}

	return hotspots, nil
}

func (c *Client) HotspotInfo(ctx context.Context, locId string, opts ...RequestOption) (*HotspotInfo, error) {
	if locId == "" {
		return nil, fmt.Errorf("locId cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.HotspotInfo, locId)
	params := processOptions(opts...)

	var hotspot HotspotInfo
	err := c.get(ctx, endpoint, params.URLParams, &hotspot)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotspot info: %w", err)
	}

	return &hotspot, nil
}
