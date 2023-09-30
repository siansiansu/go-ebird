package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

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

type RegionInfoOptions struct {
	RegionNameFormat string `url:"regionNameFormat"`
	Delim            string `url:"delim"`
}

func (c *Client) RegionInfo(ctx context.Context, region string, query interface{}) (*RegionInfoResponse, error) {
	endpoint := fmt.Sprintf(APIEndointRegionInfo, region)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	return decodeToRegionInfoResponse(req)
}

func (c *Client) SubRegionList(ctx context.Context, regionType, parentRegionCode string) ([]SubRegionListResponse, error) {
	endpoint := fmt.Sprintf(APIEndointSubRegionInfo, regionType, parentRegionCode)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	return decodeToSubRegionListResponse(req)
}

func decodeToRegionInfoResponse(res *http.Response) (*RegionInfoResponse, error) {
	decoder := json.NewDecoder(res.Body)
	result := &RegionInfoResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, err
	}
	return result, nil
}

func decodeToSubRegionListResponse(res *http.Response) ([]SubRegionListResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []SubRegionListResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}
