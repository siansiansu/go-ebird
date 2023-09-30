package ebird

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

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

// Taxonomic Forms
func (c *Client) TaxonomicForms(ctx context.Context, speciesCode string) ([]string, error) {
	endpoint := fmt.Sprintf(APITaxonomicForms, speciesCode)
	req, err := c.get(ctx, endpoint, nil)
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
func (c *Client) EbirdTaxonomy(ctx context.Context, options interface{}) ([]EbirdTaxonomyResponse, error) {
	endpoint := APIEbirdTaxonomy
	opts := addOptions(options)
	opts.Add("Fmt", "json")
	req, err := c.get(ctx, endpoint, opts)
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
