package ebird

import (
	"context"
	"fmt"
	"time"
)

type Obs struct {
	SpeciesCode      string   `json:"speciesCode,omitempty"`
	HideFlags        []string `json:"hideFlags,omitempty"`
	ObsDt            string   `json:"obsDt,omitempty"`
	Subnational1Code string   `json:"subnational1Code,omitempty"`
	HowManyAtleast   int      `json:"howManyAtleast,omitempty"`
	HowManyAtmost    int      `json:"howManyAtmost,omitempty"`
	SubId            string   `json:"subId,omitempty"`
	ProjId           string   `json:"projId,omitempty"`
	ObsId            string   `json:"obsId,omitempty"`
	HowManyStr       string   `json:"howManyStr,omitempty"`
	Present          bool     `json:"present,omitempty"`
}

type SubAux struct {
	SubId           string `json:"subId,omitempty"`
	FieldName       string `json:"fieldName,omitempty"`
	EntryMethodCode string `json:"entryMethodCode,omitempty"`
	AuxCode         string `json:"auxCode,omitempty"`
}

type ViewChecklist struct {
	ProjId                      string        `json:"projId,omitempty"`
	SubId                       string        `json:"subId,omitempty"`
	ProtocolId                  string        `json:"protocolId,omitempty"`
	LocId                       string        `json:"locId,omitempty"`
	GroupId                     string        `json:"groupId,omitempty"`
	DurationHrs                 float32       `json:"durationHrs,omitempty"`
	AllObsReported              bool          `json:"allObsReported,omitempty"`
	CreationDt                  string        `json:"creationDt,omitempty"`
	LastEditedDt                string        `json:"lastEditedDt,omitempty"`
	ObsDt                       string        `json:"obsDt,omitempty"`
	ObsTimeValid                bool          `json:"obsTimeValid,omitempty"`
	ChecklistId                 string        `json:"checklistId,omitempty"`
	NumObservers                int           `json:"numObservers,omitempty"`
	EffortDistanceKm            float32       `json:"effortDistanceKm,omitempty"`
	EffortDistanceEnteredUnit   string        `json:"effortDistanceEnteredUnit,omitempty"`
	Subnational1Code            string        `json:"subnational1Code,omitempty"`
	SubmissionMethodCode        string        `json:"submissionMethodCode,omitempty"`
	SubmissionMethodVersion     string        `json:"submissionMethodVersion,omitempty"`
	UserDisplayName             string        `json:"userDisplayName,omitempty"`
	NumSpecies                  int           `json:"numSpecies,omitempty"`
	SubmissionMethodVersionDisp string        `json:"submissionMethodVersionDisp,omitempty"`
	SubAux                      []SubAux      `json:"subAux,omitempty"`
	SubAuxAi                    []string      `json:"subAuxAi,omitempty"`
	Obs                         []Observation `json:"obs,omitempty"`
}

type RegionalStatisticsOnDate struct {
	NumChecklists   int `json:"numChecklists,omitempty"`
	NumContributors int `json:"numContributors,omitempty"`
	NumSpecies      int `json:"numSpecies,omitempty"`
}

type ChecklistFeedOnDate struct {
	LocId           string      `json:"locId,omitempty"`
	SubId           string      `json:"subId,omitempty"`
	UserDisplayName string      `json:"userDisplayName,omitempty"`
	NumSpecies      int         `json:"numSpecies,omitempty"`
	ObsDt           string      `json:"obsDt,omitempty"`
	ObsTime         string      `json:"obsTime,omitempty"`
	IsoObsDate      string      `json:"isoObsDate,omitempty"`
	SubID           string      `json:"subID,omitempty"`
	Loc             HotspotInfo `json:"loc,omitempty"`
}

type Top100 struct {
	ProfileHandle         string `json:"profileHandle,omitempty"`
	UserDisplayName       string `json:"userDisplayName,omitempty"`
	NumSpecies            int    `json:"numSpecies,omitempty"`
	NumCompleteChecklists int    `json:"numCompleteChecklists,omitempty"`
	RowNum                int    `json:"rowNum,omitempty"`
	UserId                string `json:"userId,omitempty"`
}

func (c *Client) Top100(ctx context.Context, regionCode string, date time.Time, opts ...RequestOption) ([]Top100, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.Top100, regionCode, date.Year(), date.Month(), date.Day())
	params := processOptions(opts...)

	var top100 []Top100
	err := c.get(ctx, endpoint, params.URLParams, &top100)
	if err != nil {
		return nil, fmt.Errorf("failed to get Top 100: %w", err)
	}

	return top100, nil
}

func (c *Client) ChecklistFeedOnDate(ctx context.Context, regionCode string, date time.Time, opts ...RequestOption) ([]ChecklistFeedOnDate, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.ChecklistFeedOnDate, regionCode, date.Year(), date.Month(), date.Day())
	params := processOptions(opts...)

	var feed []ChecklistFeedOnDate
	err := c.get(ctx, endpoint, params.URLParams, &feed)
	if err != nil {
		return nil, fmt.Errorf("failed to get checklist feed: %w", err)
	}

	return feed, nil
}

func (c *Client) RegionalStatisticsOnDate(ctx context.Context, regionCode string, date time.Time, opts ...RequestOption) (*RegionalStatisticsOnDate, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.RegionalStatisticsOnDate, regionCode, date.Year(), date.Month(), date.Day())
	params := processOptions(opts...)

	var stats RegionalStatisticsOnDate
	err := c.get(ctx, endpoint, params.URLParams, &stats)
	if err != nil {
		return nil, fmt.Errorf("failed to get regional statistics: %w", err)
	}

	return &stats, nil
}

func (c *Client) SpeciesListForRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]string, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.SpeciesListForRegion, regionCode)
	params := processOptions(opts...)

	var speciesList []string
	err := c.get(ctx, endpoint, params.URLParams, &speciesList)
	if err != nil {
		return nil, fmt.Errorf("failed to get species list for region: %w", err)
	}

	return speciesList, nil
}

func (c *Client) ViewChecklist(ctx context.Context, subId string, opts ...RequestOption) (*ViewChecklist, error) {
	if subId == "" {
		return nil, fmt.Errorf("subId cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.ViewChecklist, subId)
	params := processOptions(opts...)

	var checklist ViewChecklist
	err := c.get(ctx, endpoint, params.URLParams, &checklist)
	if err != nil {
		return nil, fmt.Errorf("failed to view checklist: %w", err)
	}

	return &checklist, nil
}
