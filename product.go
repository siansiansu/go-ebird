package ebird

import (
	"context"
	"fmt"
)

type Obs struct {
	SpeciesCode      string   `json:"speciesCode,omitempty"`
	HideFlags        []string `json:"hideFlags,omitempty"`
	ObsDt            string   `json:"obsDt,omitempty"`
	Subnational1Code string   `json:"subnational1Code,omitempty"`
	HowManyAtleast   int32    `json:"howManyAtleast,omitempty"`
	HowManyAtmost    int32    `json:"howManyAtmost,omitempty"`
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
	ProjId                      string  `json:"projId,omitempty"`
	SubId                       string  `json:"subId,omitempty"`
	ProtocolId                  string  `json:"protocolId,omitempty"`
	LocId                       string  `json:"locId,omitempty"`
	GroupId                     string  `json:"groupId,omitempty"`
	DurationHrs                 float32 `json:"durationHrs,omitempty"`
	AllObsReported              bool    `json:"allObsReported,omitempty"`
	CreationDt                  string  `json:"creationDt,omitempty"`
	LastEditedDt                string  `json:"lastEditedDt,omitempty"`
	ObsDt                       string  `json:"obsDt,omitempty"`
	ObsTimeValid                bool    `json:"obsTimeValid,omitempty"`
	ChecklistId                 string  `json:"checklistId,omitempty"`
	NumObservers                int32   `json:"numObservers,omitempty"`
	EffortDistanceKm            float32 `json:"effortDistanceKm,omitempty"`
	EffortDistanceEnteredUnit   string  `json:"effortDistanceEnteredUnit,omitempty"`
	Subnational1Code            string  `json:"subnational1Code,omitempty"`
	SubmissionMethodCode        string  `json:"submissionMethodCode,omitempty"`
	SubmissionMethodVersion     string  `json:"submissionMethodVersion,omitempty"`
	UserDisplayName             string  `json:"userDisplayName,omitempty"`
	NumSpecies                  int32   `json:"numSpecies,omitempty"`
	SubmissionMethodVersionDisp string  `json:"submissionMethodVersionDisp,omitempty"`
	SubAux                      []SubAux
	SubAuxAi                    []string `json:"subAuxAi,omitempty"`
	Obs                         []Obs
}

type RegionalStatisticsOnDate struct {
	NumChecklists   int32 `json:"numChecklists,omitempty"`
	NumContributors int32 `json:"numContributors,omitempty"`
	NumSpecies      int32 `json:"numSpecies,omitempty"`
}
type ChecklistFeedOnDate struct {
	LocId           string      `json:"locId,omitempty"`
	SubId           string      `json:"subId,omitempty"`
	UserDisplayName string      `json:"userDisplayName,omitempty"`
	NumSpecies      int32       `json:"numSpecies,omitempty"`
	ObsDt           string      `json:"obsDt,omitempty"`
	ObsTime         string      `json:"obsTime,omitempty"`
	IsoObsDate      string      `json:"isoObsDate,omitempty"`
	SubID           string      `json:"subID,omitempty"`
	Loc             HotspotInfo `json:"loc,omitempty"`
}

type Top100 struct {
	ProfileHandle         string `json:"profileHandle,omitempty"`
	UserDisplayName       string `json:"userDisplayName,omitempty"`
	NumSpecies            int32  `json:"numSpecies,omitempty"`
	NumCompleteChecklists int32  `json:"numCompleteChecklists,omitempty"`
	RowNum                int32  `json:"rowNum,omitempty"`
	UserId                string `json:"userId,omitempty"`
}

func (c *Client) Top100(ctx context.Context, regionCode string, y, m, d int, opts ...RequestOption) ([]Top100, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	if y <= 0 || m <= 0 || d <= 0 {
		return nil, fmt.Errorf("invalid date components: year (y), month (m), and day (d) must be greater than 0")
	}

	ebirdURL := fmt.Sprintf(APIEndpointTop100, regionCode, y, m, d)

	var t []Top100
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) ChecklistFeedOnDate(ctx context.Context, regionCode string, y, m, d int, opts ...RequestOption) ([]ChecklistFeedOnDate, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	if y <= 0 || m <= 0 || d <= 0 {
		return nil, fmt.Errorf("invalid date components: year (y), month (m), and day (d) must be greater than 0")
	}

	ebirdURL := fmt.Sprintf(APIEndpointChecklistFeedOnDate, regionCode, y, m, d)

	var t []ChecklistFeedOnDate
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) RegionalStatisticsOnDate(ctx context.Context, regionCode string, y, m, d int, opts ...RequestOption) (*RegionalStatisticsOnDate, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	if y <= 0 || m <= 0 || d <= 0 {
		return nil, fmt.Errorf("invalid date components: year (y), month (m), and day (d) must be greater than 0")
	}

	ebirdURL := fmt.Sprintf(APIEndpointRegionalStatisticsOnDate, regionCode, y, m, d)

	var t RegionalStatisticsOnDate
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}
	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (c *Client) SpeciesListForRegion(ctx context.Context, regionCode string, opts ...RequestOption) ([]string, error) {
	if regionCode == "" {
		return nil, fmt.Errorf("regionCode cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndpointSpeciesListForRegion, regionCode)

	var t []string

	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) ViewChecklist(ctx context.Context, subId string, opts ...RequestOption) (*ViewChecklist, error) {
	if subId == "" {
		return nil, fmt.Errorf("subId cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndpointViewChecklist, subId)

	var t ViewChecklist
	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}
	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
