package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type HotspotInfoResponse struct {
	LocId            string  `json:"locId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Latitude         float32 `json:"latitude,omitempty"`
	Longitude        float32 `json:"longitude,omitempty"`
	CountryCode      string  `json:"countryCode,omitempty"`
	CountryName      string  `json:"countryName,omitempty"`
	Subnational1Name string  `json:"subnational1Name,omitempty"`
	Subnational1Code string  `json:"subnational1Code,omitempty"`
	IsHotspot        bool    `json:"isHotspot,omitempty"`
	LocID            string  `json:"locID,omitempty"`
	LocName          string  `json:"locName,omitempty"`
	Lat              float32 `json:"lat,omitempty"`
	Lng              float32 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
}

type NearbyHotspotsResponse struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Lat               float32 `json:"lat" validate:"required"`
	Lng               float32 `json:"lng" validate:"required"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime uint32  `json:"numSpeciesAllTime,omitempty"`
}

type HotspotsInRegionResponse struct {
	LocId             string  `json:"locId,omitempty"`
	LocName           string  `json:"locName,omitempty"`
	CountryCode       string  `json:"countryCode,omitempty"`
	Subnational1Code  string  `json:"subnational1Code,omitempty"`
	Lat               float32 `json:"lat,omitempty"`
	Lng               float32 `json:"lng,omitempty"`
	LatestObsDt       string  `json:"latestObsDt,omitempty"`
	NumSpeciesAllTime uint32  `json:"numSpeciesAllTime,omitempty"`
}

type HotspotsInRegionOptions struct {
	Back uint16 `url:"back"`
	Fmt  string `url:"fmt"`
}

func (c *Client) HotspotInfo(ctx context.Context, locId string) (*HotspotInfoResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointHotspotInfo, locId)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	return decodeToHotspotInfo(req)
}

func (c *Client) NearbyHotspots(ctx context.Context, query interface{}) ([]NearbyHotspotsResponse, error) {
	endpoint := APIEndpointNearbyHotspots
	q := addOptions(query)
	q.Add("Fmt", "json")
	options := RequestOptions{
		Query:   q,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToNearbyHotspots(req)
}

func (c *Client) HotspotsInRegion(ctx context.Context, regionCode string, query interface{}) ([]HotspotsInRegionResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointHotspotsInRegion, regionCode)
	q := addOptions(query)
	q.Add("Fmt", "json")
	options := RequestOptions{
		Query:   q,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToHotspotsInRegion(req)
}

func decodeToHotspotsInRegion(res *http.Response) ([]HotspotsInRegionResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []HotspotsInRegionResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToNearbyHotspots(res *http.Response) ([]NearbyHotspotsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []NearbyHotspotsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToHotspotInfo(res *http.Response) (*HotspotInfoResponse, error) {
	decoder := json.NewDecoder(res.Body)
	result := &HotspotInfoResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}
