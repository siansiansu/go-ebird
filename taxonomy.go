package ebird

import (
	"context"
	"fmt"
)

type TaxonomicGroup struct {
	GroupName        string      `json:"groupName,omitempty"`
	GroupOrder       int         `json:"groupOrder,omitempty"`
	TaxonOrderBounds [][]float64 `json:"taxonOrderBounds,omitempty"`
}

type TaxonomyVersion struct {
	AuthorityVer float64 `json:"authorityVer,omitempty"`
	Latest       bool    `json:"latest,omitempty"`
}

type TaxaLocaleCode struct {
	Code       string `json:"code,omitempty"`
	Name       string `json:"name,omitempty"`
	LastUpdate string `json:"lastUpdate,omitempty"`
}

type EbirdTaxon struct {
	SciName       string   `json:"sciName,omitempty"`
	ComName       string   `json:"comName,omitempty"`
	SpeciesCode   string   `json:"speciesCode,omitempty"`
	Category      string   `json:"category,omitempty"`
	TaxonOrder    float64  `json:"taxonOrder,omitempty"`
	BandingCodes  []string `json:"bandingCodes,omitempty"`
	ComNameCodes  []string `json:"comNameCodes,omitempty"`
	SciNameCodes  []string `json:"sciNameCodes,omitempty"`
	Order         string   `json:"order,omitempty"`
	FamilyCode    string   `json:"familyCode,omitempty"`
	FamilyComName string   `json:"familyComName,omitempty"`
	FamilySciName string   `json:"familySciName,omitempty"`
}

func (c *Client) TaxonomicGroups(ctx context.Context, speciesGrouping string, opts ...RequestOption) ([]TaxonomicGroup, error) {
	if speciesGrouping == "" {
		return nil, fmt.Errorf("speciesGrouping cannot be empty")
	}
	endpoint := fmt.Sprintf(APIEndpoints.TaxonomicGroups, speciesGrouping)
	params := processOptions(opts...)

	var groups []TaxonomicGroup
	err := c.get(ctx, endpoint, params.URLParams, &groups)
	if err != nil {
		return nil, fmt.Errorf("failed to get taxonomic groups: %w", err)
	}

	return groups, nil
}

func (c *Client) TaxonomyVersions(ctx context.Context, opts ...RequestOption) ([]TaxonomyVersion, error) {
	params := processOptions(opts...)

	var versions []TaxonomyVersion
	err := c.get(ctx, APIEndpoints.TaxonomyVersions, params.URLParams, &versions)
	if err != nil {
		return nil, fmt.Errorf("failed to get taxonomy versions: %w", err)
	}

	return versions, nil
}

func (c *Client) TaxaLocaleCodes(ctx context.Context, opts ...RequestOption) ([]TaxaLocaleCode, error) {
	params := processOptions(opts...)

	var codes []TaxaLocaleCode
	err := c.get(ctx, APIEndpoints.TaxaLocaleCodes, params.URLParams, &codes)
	if err != nil {
		return nil, fmt.Errorf("failed to get taxa locale codes: %w", err)
	}

	return codes, nil
}

func (c *Client) TaxonomicForms(ctx context.Context, speciesCode string, opts ...RequestOption) ([]string, error) {
	if speciesCode == "" {
		return nil, fmt.Errorf("speciesCode cannot be empty")
	}

	endpoint := fmt.Sprintf(APIEndpoints.TaxonomicForms, speciesCode)
	params := processOptions(opts...)

	var forms []string
	err := c.get(ctx, endpoint, params.URLParams, &forms)
	if err != nil {
		return nil, fmt.Errorf("failed to get taxonomic forms: %w", err)
	}

	return forms, nil
}

func (c *Client) EbirdTaxonomy(ctx context.Context, opts ...RequestOption) ([]EbirdTaxon, error) {
	params := processOptions(opts...)
	params.URLParams.Set("fmt", "json")

	var taxonomy []EbirdTaxon
	err := c.get(ctx, APIEndpoints.EbirdTaxonomy, params.URLParams, &taxonomy)
	if err != nil {
		return nil, fmt.Errorf("failed to get eBird taxonomy: %w", err)
	}

	return taxonomy, nil
}
