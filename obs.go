package ebird

import (
	"context"
	"fmt"
	"time"
)

type Observation struct {
	SpeciesCode     string  `json:"speciesCode,omitempty"`
	ComName         string  `json:"comName,omitempty"`
	SciName         string  `json:"sciName,omitempty"`
	LocId           string  `json:"locId,omitempty"`
	LocName         string  `json:"locName,omitempty"`
	ObsDt           string  `json:"obsDt,omitempty"`
	HowMany         int     `json:"howMany,omitempty"`
	Lat             float64 `json:"lat,omitempty"`
	Lng             float64 `json:"lng,omitempty"`
	ObsValid        bool    `json:"obsValid,omitempty"`
	ObsReviewed     bool    `json:"obsReviewed,omitempty"`
	LocationPrivate bool    `json:"locationPrivate,omitempty"`
	SubId           string  `json:"subId,omitempty"`
	ExoticCategory  string  `json:"exoticCategory,omitempty"`
}

type Location struct {
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
	LocName          string  `json:"locName,omitempty"`
	Lat              float64 `json:"lat,omitempty"`
	Lng              float64 `json:"lng,omitempty"`
	HierarchicalName string  `json:"hierarchicalName,omitempty"`
	LocID            string  `json:"locID,omitempty"`
}

type RecentChecklistFeed struct {
	LocId           string   `json:"locId,omitempty"`
	SubId           string   `json:"subId,omitempty"`
	UserDisplayName string   `json:"userDisplayName,omitempty"`
	NumSpecies      int      `json:"numSpecies,omitempty"`
	ObsDt           string   `json:"obsDt,omitempty"`
	ObsTime         string   `json:"obsTime,omitempty"`
	IsoObsDate      string   `json:"isoObsDate,omitempty"`
	SubID           string   `json:"subID,omitempty"`
	Loc             Location `json:"loc,omitempty"`
}

func (c *Client) RecentObservationsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]Observation, error) {
	return c.getObservations(ctx, fmt.Sprintf(APIEndpoints.RecentObservationsInRegion, regionCode), opts...)
}

func (c *Client) RecentNotableObservationsInRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]Observation, error) {
	return c.getObservations(ctx, fmt.Sprintf(APIEndpoints.RecentNotableObservationsInRegion, regionCode), opts...)
}

func (c *Client) RecentObservationsOfSpeciesInRegion(ctx context.Context, regionCode, speciesCode string, opts ...RequestOption) ([]Observation, error) {
	if speciesCode == "" {
		return nil, fmt.Errorf("speciesCode cannot be empty")
	}
	return c.getObservations(ctx, fmt.Sprintf(APIEndpoints.RecentObservationsOfSpeciesInRegion, regionCode, speciesCode), opts...)
}

func (c *Client) RecentNearbyObservations(ctx context.Context, opts ...RequestOption) ([]Observation, error) {
	params := processOptions(opts...)
	if params.URLParams.Get("lat") == "" || params.URLParams.Get("lng") == "" {
		return nil, fmt.Errorf("must provide Lat and Lng parameters")
	}
	return c.getObservations(ctx, APIEndpoints.RecentNearbyObservations, opts...)
}

func (c *Client) RecentNearbyObservationsOfSpecies(ctx context.Context, speciesCode string, opts ...RequestOption) ([]Observation, error) {
	if speciesCode == "" {
		return nil, fmt.Errorf("speciesCode cannot be empty")
	}
	return c.getObservations(ctx, fmt.Sprintf(APIEndpoints.RecentNearbyObservationsOfSpecies, speciesCode), opts...)
}

func (c *Client) NearestObservationsOfSpecies(ctx context.Context, speciesCode string, opts ...RequestOption) ([]Observation, error) {
	if speciesCode == "" {
		return nil, fmt.Errorf("speciesCode cannot be empty")
	}
	return c.getObservations(ctx, fmt.Sprintf(APIEndpoints.NearestObservationsOfSpecies, speciesCode), opts...)
}

func (c *Client) RecentNearbyNotableObservations(ctx context.Context, opts ...RequestOption) ([]Observation, error) {
	return c.getObservations(ctx, APIEndpoints.RecentNearbyNotableObservations, opts...)
}

func (c *Client) RecentChecklistsFeed(ctx context.Context, regionCode string, opts ...RequestOption) ([]RecentChecklistFeed, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}
	endpoint := fmt.Sprintf(APIEndpoints.RecentChecklistsFeed, regionCode)
	params := processOptions(opts...)

	var checklists []RecentChecklistFeed
	err := c.get(ctx, endpoint, params.URLParams, &checklists)
	if err != nil {
		return nil, fmt.Errorf("failed to get recent checklists feed: %w", err)
	}
	return checklists, nil
}

func (c *Client) HistoricObservationsOnDate(ctx context.Context, regionCode string, date time.Time, opts ...RequestOption) ([]Observation, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}
	endpoint := fmt.Sprintf(APIEndpoints.HistoricObservationsOnDate, regionCode, date.Year(), date.Month(), date.Day())
	return c.getObservations(ctx, endpoint, opts...)
}

func (c *Client) getObservations(ctx context.Context, endpoint string, opts ...RequestOption) ([]Observation, error) {
	params := processOptions(opts...)
	var observations []Observation
	err := c.get(ctx, endpoint, params.URLParams, &observations)
	if err != nil {
		return nil, fmt.Errorf("failed to get observations: %w", err)
	}
	return observations, nil
}
