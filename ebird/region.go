package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegionInfoOptions struct {
	RegionNameFormat string `url:"regionNameFormat"`
	Delim            string `url:"delim"`
}

type SubRegionListOptions struct {
	Fmt string `url:"fmt"`
}

type Bounds struct {
	MinX float64 `json:"minX,omitempty"`
	MaxX float64 `json:"maxX,omitempty"`
	MinY float64 `json:"minY,omitempty"`
	MaxY float64 `json:"maxY,omitempty"`
}

type RegionInfoResponse struct {
	Bounds    Bounds
	Result    string  `json:"result,omitempty"`
	Code      string  `json:"code,omitempty"`
	Type      string  `json:"type,omitempty"`
	Longitude float64 `json:"longitude,omitempty"`
	Latitude  float64 `json:"latitude,omitempty"`
}

type SubRegionListResponse struct {
	Code string `json:"code,omitempty"`
	Name string `json:"name,omitempty"`
}

// Region Info
func decodeToRegionInfoResponse(res *http.Response) (*RegionInfoResponse, error) {
	decoder := json.NewDecoder(res.Body)
	result := &RegionInfoResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

// Sub Region List
func decodeToSubRegionListResponse(res *http.Response) ([]SubRegionListResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []SubRegionListResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
func (c *Client) SubRegionList(ctx context.Context, regionType, parentRegionCode string, options interface{}) ([]SubRegionListResponse, error) {
	endpoint := fmt.Sprintf(APIEndointSubRegionInfo, regionType, parentRegionCode)
	opts := addOptions(options)
	req, err := c.get(ctx, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return decodeToSubRegionListResponse(req)
}

func (c *Client) RegionInfo(ctx context.Context, region string, options interface{}) (*RegionInfoResponse, error) {
	endpoint := fmt.Sprintf(APIEndointRegionInfo, region)
	opts := addOptions(options)
	req, err := c.get(ctx, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return decodeToRegionInfoResponse(req)
}
