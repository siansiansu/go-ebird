package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Obs struct {
	SpeciesCode      string   `json:"speciesCode,omitempty"`
	HideFlags        []string `json:"hideFlags,omitempty"`
	ObsDt            string   `json:"obsDt,omitempty"`
	Subnational1Code string   `json:"subnational1Code,omitempty"`
	HowManyAtleast   uint32   `json:"howManyAtleast,omitempty"`
	HowManyAtmost    uint32   `json:"howManyAtmost,omitempty"`
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
type ViewChecklistResponse struct {
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
	NumObservers                uint32  `json:"numObservers,omitempty"`
	EffortDistanceKm            float32 `json:"effortDistanceKm,omitempty"`
	EffortDistanceEnteredUnit   string  `json:"effortDistanceEnteredUnit,omitempty"`
	Subnational1Code            string  `json:"subnational1Code,omitempty"`
	SubmissionMethodCode        string  `json:"submissionMethodCode,omitempty"`
	SubmissionMethodVersion     string  `json:"submissionMethodVersion,omitempty"`
	UserDisplayName             string  `json:"userDisplayName,omitempty"`
	NumSpecies                  uint32  `json:"numSpecies,omitempty"`
	SubmissionMethodVersionDisp string  `json:"submissionMethodVersionDisp,omitempty"`
	SubAux                      []SubAux
	SubAuxAi                    []string `json:"subAuxAi,omitempty"`
	Obs                         []Obs
}

type RegionalStatisticsOnDateResponse struct {
	NumChecklists   uint32 `json:"numChecklists,omitempty"`
	NumContributors uint32 `json:"numContributors,omitempty"`
	NumSpecies      uint32 `json:"numSpecies,omitempty"`
}
type ChecklistFeedOnDateResponse struct {
	LocId           string              `json:"locId,omitempty"`
	SubId           string              `json:"subId,omitempty"`
	UserDisplayName string              `json:"userDisplayName,omitempty"`
	NumSpecies      uint32              `json:"numSpecies,omitempty"`
	ObsDt           string              `json:"obsDt,omitempty"`
	ObsTime         string              `json:"obsTime,omitempty"`
	IsoObsDate      string              `json:"isoObsDate,omitempty"`
	SubID           string              `json:"subID,omitempty"`
	Loc             HotspotInfoResponse `json:"loc,omitempty"`
}

type Top100Response struct {
	ProfileHandle         string `json:"profileHandle,omitempty"`
	UserDisplayName       string `json:"userDisplayName,omitempty"`
	NumSpecies            uint32 `json:"numSpecies,omitempty"`
	NumCompleteChecklists uint32 `json:"numCompleteChecklists,omitempty"`
	RowNum                uint32 `json:"rowNum,omitempty"`
	UserId                string `json:"userId,omitempty"`
}
type Top100Options struct {
	RankedBy   string `url:"rankedBy"`
	MaxResults uint32 `url:"maxResults"`
}

type ChecklistFeedOnDateOptions struct {
	SortKey    string `url:"sortKey"`
	MaxResults uint32 `url:"maxResults"`
}

func (c *Client) ViewChecklist(ctx context.Context, subId string) (*ViewChecklistResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointViewChecklist, subId)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToViewChecklist(req)
}

func (c *Client) SpeciesListForRegion(ctx context.Context, regionCode string) ([]string, error) {
	endpoint := fmt.Sprintf(APIEndpointSpeciesListForRegion, regionCode)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToSpeciesListForRegion(req)
}

func (c *Client) RegionalStatisticsOnDate(ctx context.Context, regionCode string, y, m, d uint32) (*RegionalStatisticsOnDateResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointRegionalStatisticsOnDate, regionCode, y, m, d)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToRegionalStatisticsOnDate(req)
}

func (c *Client) ChecklistFeedOnDate(ctx context.Context, regionCode string, y, m, d uint32, query interface{}) ([]ChecklistFeedOnDateResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointChecklistFeedOnDate, regionCode, y, m, d)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToChecklistFeedOnDate(req)
}

func (c *Client) Top100(ctx context.Context, regionCode string, y, m, d uint32, query interface{}) ([]Top100Response, error) {
	endpoint := fmt.Sprintf(APIEndpointTop100, regionCode, y, m, d)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToTop100(req)
}

func decodeToTop100(res *http.Response) ([]Top100Response, error) {
	decoder := json.NewDecoder(res.Body)
	var result []Top100Response
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToChecklistFeedOnDate(res *http.Response) ([]ChecklistFeedOnDateResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []ChecklistFeedOnDateResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToRegionalStatisticsOnDate(res *http.Response) (*RegionalStatisticsOnDateResponse, error) {
	decoder := json.NewDecoder(res.Body)
	result := &RegionalStatisticsOnDateResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToSpeciesListForRegion(res *http.Response) ([]string, error) {
	decoder := json.NewDecoder(res.Body)
	var result []string
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

func decodeToViewChecklist(res *http.Response) (*ViewChecklistResponse, error) {
	decoder := json.NewDecoder(res.Body)
	result := &ViewChecklistResponse{}
	if err := decoder.Decode(result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}
