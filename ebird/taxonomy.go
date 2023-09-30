package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaxonomicGroupsResponse struct {
	GroupName        string      `json:"groupName,omitempty"`
	GroupOrder       uint32      `json:"groupOrder,omitempty"`
	TaxonOrderBounds [][]float32 `json:"taxonOrderBounds,omitempty"`
}

type TaxonomyVersionsResponse struct {
	AuthorityVer float32 `json:"authorityVer,omitempty"`
	Latest       bool    `json:"latest,omitempty"`
}
type TaxaLocaleCodesResponse struct {
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	LastUpdate string `json:"lastUpdate,omitempty"`
}

type EbirdTaxonomyResponse struct {
	SciName       string   `json:"sciName,omitempty"`
	ComName       string   `json:"comName,omitempty"`
	SpeciesCode   string   `json:"speciesCode,omitempty"`
	Category      string   `json:"category,omitempty"`
	TaxonOrder    float32  `json:"taxonOrder,omitempty"`
	BandingCodes  []string `json:"bandingCodes,omitempty"`
	ComNameCodes  []string `json:"comNameCodes,omitempty"`
	SciNameCodes  []string `json:"sciNameCodes,omitempty"`
	Order         string   `json:"order,omitempty"`
	FamilyCode    string   `json:"familyCode,omitempty"`
	FamilyComName string   `json:"familyComName,omitempty"`
	FamilySciName string   `json:"familySciName,omitempty"`
}
type EbirdTaxonomyOptions struct {
	Cat     string `url:"cat"`
	Fmt     string `url:"fmt"`
	Locale  string `url:"locale"`
	Species string `url:"species"`
	Version string `url:"version"`
}

// Taxonomic Groups
func (c *Client) TaxonomicGroups(ctx context.Context, speciesGrouping string, query interface{}) ([]TaxonomicGroupsResponse, error) {
	endpoint := fmt.Sprintf(APIEndpointTaxonomicGroups, speciesGrouping)
	options := RequestOptions{
		Query:   addOptions(query),
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToTaxonomicGroups(req)
}

func decodeToTaxonomicGroups(res *http.Response) ([]TaxonomicGroupsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []TaxonomicGroupsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// Taxonomy Versions
func (c *Client) TaxonomyVersions(ctx context.Context) ([]TaxonomyVersionsResponse, error) {
	endpoint := APIEndpointTaxonomyVersions
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToTaxonomyVersions(req)
}

func decodeToTaxonomyVersions(res *http.Response) ([]TaxonomyVersionsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []TaxonomyVersionsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// Taxa Locale Codes
func (c *Client) TaxaLocaleCodes(ctx context.Context, headers map[string]string) ([]TaxaLocaleCodesResponse, error) {
	endpoint := APIEndpointTaxaLocaleCodes
	options := RequestOptions{
		Query:   nil,
		Headers: headers,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToTaxaLocaleCodes(req)
}

func decodeToTaxaLocaleCodes(res *http.Response) ([]TaxaLocaleCodesResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []TaxaLocaleCodesResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// Taxonomic Forms
func (c *Client) TaxonomicForms(ctx context.Context, speciesCode string) ([]string, error) {
	endpoint := fmt.Sprintf(APIEndpointTaxonomicForms, speciesCode)
	options := RequestOptions{
		Query:   nil,
		Headers: nil,
	}
	req, err := c.get(ctx, endpoint, options)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	return decodeToTaxonomicForms(req)
}

func decodeToTaxonomicForms(res *http.Response) ([]string, error) {
	decoder := json.NewDecoder(res.Body)
	var result []string
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}

// Ebird Taxomony
func (c *Client) EbirdTaxonomy(ctx context.Context, query interface{}) ([]EbirdTaxonomyResponse, error) {
	endpoint := APIEndpointEbirdTaxonomy
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
	return decodeToEbirdTaxonomyResponse(req)
}

func decodeToEbirdTaxonomyResponse(res *http.Response) ([]EbirdTaxonomyResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []EbirdTaxonomyResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}
