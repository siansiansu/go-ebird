package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type RecentObservationsResponse struct {
	SpeciesCode     string  `json:"speciesCode,omitempty"`
	ComName         string  `json:"comName,omitempty"`
	SciName         string  `json:"sciName,omitempty"`
	LocId           string  `json:"locId,omitempty"`
	LocName         string  `json:"locName,omitempty"`
	ObsDt           string  `json:"obsDt,omitempty"`
	HowMany         uint32  `json:"howMany,omitempty"`
	Lat             float32 `json:"lat,omitempty"`
	Lng             float32 `json:"lng,omitempty"`
	ObsValid        bool    `json:"obsValid,omitempty"`
	ObsReviewed     bool    `json:"obsReviewed,omitempty"`
	LocationPrivate bool    `json:"locationPrivate,omitempty"`
	SubId           string  `json:"subId,omitempty"`
}

type RecentObservationsOptions struct {
	Back               uint32 `url:"back"`
	Cat                string `url:"cat"`
	Hotspot            bool   `url:"hotspot"`
	IncludeProvisional bool   `url:"includeProvisional"`
	MaxResults         uint32 `url:"maxResults"`
	R                  string `url:"r"`
	SppLocale          string `url:"sppLocale"`
}

type RecentNearbyObservationsOptions struct {
	Back               uint32  `url:"back"`
	Cat                string  `url:"cat"`
	Dist               uint32  `url:"dist"`
	Hotspot            bool    `url:"hotspot"`
	IncludeProvisional bool    `url:"includeProvisional"`
	Lat                float32 `json:"lat" validate:"required"`
	Lng                float32 `json:"lng" validate:"required"`
	MaxResults         uint32  `url:"maxResults"`
	R                  string  `url:"r"`
	SppLocale          string  `url:"sppLocale"`
}

type Loc struct {
	LocId            string  `json:"locId,omitempty"`
	Name             string  `json:"name,omitempty"`
	Latitude         float32 `json:"latitude,omitempty"`
	Longitude        float32 `json:"longitude,omitempty"`
	CountryCode      string  `json:"countryCode,omitempty"`
	CountryName      string  `json:"countryName,omitempty"`
	Subnational1Name string  `json:"subnational1Name,omitempty"`
	Subnational1Code string  `json:"subnational1Code,omitempty"`
	IsHotspot        bool    `json:"isHotspot,omitempty"`
	LocName          string  `json:"locName,omitempty"`
	Lat              float32 `json:"lat,omitempty"`
	Lng              float32 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
	LocID            string  `json:"locID,omitempty"`
}
type RecentChecklistsFeedResponse struct {
	LocId           string `json:"locId,omitempty"`
	SubId           string `json:"subId,omitempty"`
	UserDisplayName string `json:"userDisplayName,omitempty"`
	NumSpecies      uint32 `json:"numSpecies,omitempty"`
	ObsDt           string `json:"obsDt,omitempty"`
	ObsTime         string `json:"obsTime,omitempty"`
	IsoObsDate      string `json:"isoObsDate,omitempty"`
	SubID           string `json:"subID,omitempty"`
	Loc             Loc    `json:"loc,omitempty"`
}

type RecentChecklistsFeedOptions struct {
	MaxResults uint32 `url:"maxResults"`
}

func (c *Client) HistoricObservationsOnDate(ctx context.Context, regionCode string, y, m, d uint32, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointHistoricObservationsOnDate, regionCode, y, m, d)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentChecklistsFeed(ctx context.Context, regionCode string, query interface{}) ([]RecentChecklistsFeedResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRecentChecklistsFeed, regionCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentChecklistsFeed(req)
}

func (c *Client) RecentNearbyNotableObservations(ctx context.Context, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := APIEndpointRecentNearbyNotableObservations
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) NearestObservationsOfSpecies(ctx context.Context, speciesCode string, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointNearestObservationsOfSpecies, speciesCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentNearbyObservationsOfSpecies(ctx context.Context, speciesCode string, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRecentNearbyObservationsOfSpecies, speciesCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentNearbyObservations(ctx context.Context, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := APIEndpointRecentNearbyObservations
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentObservationsOfSpeciesInRegion(ctx context.Context, regionCode, speciesCode string, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRecentObservationsOfSpeciesInRegion, regionCode, speciesCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentNotableObservationsInRegion(ctx context.Context, regionCode string, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRecentNotableObservationsInRegion, regionCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func (c *Client) RecentObservationsInRegion(ctx context.Context, regionCode string, query interface{}) ([]RecentObservationsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRecentObservationsInRegion, regionCode)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRecentObservations(req)
}

func decodeToRecentObservations(res *http.Response) ([]RecentObservationsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []RecentObservationsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToRecentChecklistsFeed(res *http.Response) ([]RecentChecklistsFeedResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []RecentChecklistsFeedResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}
