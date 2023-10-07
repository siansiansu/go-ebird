package ebird

import (
	"context"
	"fmt"
)

type TaxonomicGroups struct {
	GroupName        string      `json:"groupName,omitempty"`
	GroupOrder       int32       `json:"groupOrder,omitempty"`
	TaxonOrderBounds [][]float32 `json:"taxonOrderBounds,omitempty"`
}

type TaxonomyVersions struct {
	AuthorityVer float32 `json:"authorityVer,omitempty"`
	Latest       bool    `json:"latest,omitempty"`
}

type TaxaLocaleCodes struct {
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	LastUpdate string `json:"lastUpdate,omitempty"`
}

type EbirdTaxonomy struct {
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

func (c *Client) TaxonomicGroups(ctx context.Context, speciesGrouping string, opts ...RequestOption) ([]TaxonomicGroups, error) {
	if speciesGrouping == "" {
		return nil, fmt.Errorf("speciesGrouping cannot be empty")
	}
	ebirdURL := fmt.Sprintf(APIEndpointTaxonomicGroups, speciesGrouping)

	var t []TaxonomicGroups

	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) TaxonomyVersions(ctx context.Context, opts ...RequestOption) ([]TaxonomyVersions, error) {
	ebirdURL := APIEndpointTaxonomyVersions

	var t []TaxonomyVersions

	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) TaxaLocaleCodes(ctx context.Context, opts ...RequestOption) ([]TaxaLocaleCodes, error) {
	ebirdURL := APIEndpointTaxaLocaleCodes

	var t []TaxaLocaleCodes

	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

func (c *Client) TaxonomicForms(ctx context.Context, speciesCode string, opts ...RequestOption) ([]string, error) {
	if speciesCode == "" {
		return nil, fmt.Errorf("speciesCode cannot be empty")
	}

	ebirdURL := fmt.Sprintf(APIEndpointTaxonomicForms, speciesCode)

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

func (c *Client) EbirdTaxonomy(ctx context.Context, opts ...RequestOption) ([]EbirdTaxonomy, error) {
	ebirdURL := APIEndpointEbirdTaxonomy

	var t []EbirdTaxonomy

	if params := processOptions(opts...).urlParams.Encode(); params != "" {
		ebirdURL += "?" + params
	}

	ebirdURL = convertToJsonFormat(ebirdURL)

	err := c.get(ctx, ebirdURL, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
