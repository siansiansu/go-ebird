package ebird

import (
	"context"
	"fmt"
	"strings"
)

type HotspotInfo struct {
	LocId            string  `json:"locId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Latitude         float32 `json:"latitude,omitempty"`
	Longitude        float32 `json:"longitude,omitempty"`
	CountryCode      string  `json:"countryCode,omitempty"`
	CountryName      string  `json:"countryName,omitempty"`
	Subnational1Name string  `json:"subnational1Name,omitempty"`
	Subnational1Code string  `json:"subnational1Code,omitempty"`
	Subnational2Code string  `json:"Subnational2Code,omitempty"`
	Subnational2Name string  `json:"Subnational2Name,omitempty"`
	IsHotspot        bool    `json:"isHotspot,omitempty"`
	LocID            string  `json:"locID,omitempty"`
	LocName          string  `json:"locName,omitempty"`
	Lat              float32 `json:"lat,omitempty"`
	Lng              float32 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
}

type NearbyHotspots struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Lat               float32 `json:"lat" validate:"required"`
	Lng               float32 `json:"lng" validate:"required"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime int32   `json:"numSpeciesAllTime,omitempty"`
}

type HotspotsInRegion struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Subnational2Code  string  `json:"subnational2Code,omitempty"`
	Lat               float32 `json:"lat,omitempty"`
	Lng               float32 `json:"lng,omitempty"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime int32   `json:"numSpeciesAllTime,omitempty"`
}

func (c *Client) HotspotsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]HotspotsInRegion, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}
	ebirdURL := fmt.Sprintf(APIEndpointHotspotsInRegion, regionCode)

	var t []HotspotsInRegion
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	ebirdURL = convertToJsonFormat(ebirdURL)

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) NearbyHotspots(ctx context.Context, opts ...RequestOption) ([]NearbyHotspots, error) {
	ebirdURL := APIEndpointNearbyHotspots

	var t []NearbyHotspots
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	if !strings.Contains(ebirdURL, "lat") || !strings.Contains(ebirdURL, "lng") {
		return nil, fmt.Errorf("must provide Lat and Lng parameters")
	}

	ebirdURL = convertToJsonFormat(ebirdURL)

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (c *Client) HotspotInfo(ctx context.Context, locId string, opts ...RequestOption) (*HotspotInfo, error) {
	if locId == "" {
		return nil, fmt.Errorf("locId cannot be empty")
	}
	ebirdURL := fmt.Sprintf(APIEndpointHotspotInfo, locId)

	var t HotspotInfo
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}
	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
