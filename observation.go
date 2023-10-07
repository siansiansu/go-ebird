package ebird

import (
	"context"
	"fmt"
	"strings"
)

type Observations struct {
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
	ExoticCategory  string  `json:"exoticCategory,omitempty"`
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
	Subnational2Code string  `json:"Subnational2Code,omitempty"`
	Subnational2Name string  `json:"Subnational2Name,omitempty"`
	IsHotspot        bool    `json:"isHotspot,omitempty"`
	LocName          string  `json:"locName,omitempty"`
	Lat              float32 `json:"lat,omitempty"`
	Lng              float32 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
	LocID            string  `json:"locID,omitempty"`
}

type RecentChecklistsFeed struct {
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

func (c *Client) RecentObservationsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointRecentObservationsInRegion, regionCode)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentNotableObservationsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointRecentNotableObservationsInRegion, regionCode)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentObservationsOfSpeciesInRegion(ctx context.Context, regionCode, speciesCode string, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointRecentObservationsOfSpeciesInRegion, regionCode, speciesCode)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentNearbyObservations(ctx context.Context, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := APIEndpointRecentNearbyObservations

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	if !strings.Contains(ebirdURL, "lat") || !strings.Contains(ebirdURL, "lng") {
		return nil, fmt.Errorf("must provide Lat and Lng parameters")
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentNearbyObservationsOfSpecies(ctx context.Context, speciesCode string, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointRecentNearbyObservationsOfSpecies, speciesCode)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) NearestObservationsOfSpecies(ctx context.Context, speciesCode string, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointNearestObservationsOfSpecies, speciesCode)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentNearbyNotableObservations(ctx context.Context, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := APIEndpointRecentNearbyNotableObservations

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RecentChecklistsFeed(ctx context.Context, regionCode string, opts ...RequestOption) ([]RecentChecklistsFeed, error) {
	ebirdURL := fmt.Sprintf(APIEndpointRecentChecklistsFeed, regionCode)

	var t []RecentChecklistsFeed
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) HistoricObservationsOnDate(ctx context.Context, regionCode string, y, m, d int, opts ...RequestOption) ([]Observations, error) {
	ebirdURL := fmt.Sprintf(APIEndpointHistoricObservationsOnDate, regionCode, y, m, d)

	var t []Observations
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
