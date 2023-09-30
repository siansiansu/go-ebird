package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type TaxonomicFormsResponse []string

type TaxonomicFormsOptions struct {
	SpeciesCode string `url:"speciesCode"`
}

type EbirdTaxonomyOptions struct {
	Cat     string `url:"cat"`
	Fmt     string `url:"fmt"`
	Locale  string `url:"locale"`
	Species string `url:"species"`
	Version string `url:"version"`
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

func (c *Client) TaxonomicForms(ctx context.Context, options interface{}) (*TaxonomicFormsResponse, error) {
	endpoint := APITaxonomicForms
	opts := addOptions(options)
	req, err := c.get(ctx, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return decodeToTaxonomicForms(req)
}

func (c *Client) EbirdTaxonomy(ctx context.Context, options interface{}) ([]EbirdTaxonomyResponse, error) {
	endpoint := APIEbirdTaxonomy
	opts := addOptions(options)
	opts.Add("Fmt", "json")
	req, err := c.get(ctx, endpoint, opts)
	if err != nil {
		return nil, err
	}
	return decodeToEbirdTaxonomyResponse(req)
}

func decodeToTaxonomicForms(res *http.Response) (*TaxonomicFormsResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result TaxonomicFormsResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return &result, nil
}

// Ebird Taxomony
func decodeToEbirdTaxonomyResponse(res *http.Response) ([]EbirdTaxonomyResponse, error) {
	decoder := json.NewDecoder(res.Body)
	var result []EbirdTaxonomyResponse
	if err := decoder.Decode(&result); err != nil {
		return nil, fmt.Errorf("failed to decode response: %v", err)
	}
	return result, nil
}
