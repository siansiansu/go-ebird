package ebird

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	APIEndpointBase = "https://api.ebird.org/v2/"
	defaultTimeout  = 10 * time.Second
)

// APIEndpoints holds all the API endpoints
var APIEndpoints = struct {
	HistoricObservationsOnDate          string
	NearestObservationsOfSpecies        string
	RecentChecklistsFeed                string
	RecentNearbyNotableObservations     string
	RecentNearbyObservations            string
	RecentNearbyObservationsOfSpecies   string
	RecentNotableObservationsInRegion   string
	RecentObservationsInRegion          string
	RecentObservationsOfSpeciesInRegion string
	ChecklistFeedOnDate                 string
	RegionalStatisticsOnDate            string
	SpeciesListForRegion                string
	Top100                              string
	ViewChecklist                       string
	AdjacentRegions                     string
	HotspotInfo                         string
	HotspotsInRegion                    string
	NearbyHotspots                      string
	EbirdTaxonomy                       string
	TaxaLocaleCodes                     string
	TaxonomicForms                      string
	TaxonomicGroups                     string
	TaxonomyVersions                    string
	RegionInfo                          string
	SubRegionInfo                       string
}{
	HistoricObservationsOnDate:          "data/obs/%s/historic/%d/%d/%d",
	NearestObservationsOfSpecies:        "data/nearest/geo/recent/%s",
	RecentChecklistsFeed:                "product/lists/%s",
	RecentNearbyNotableObservations:     "data/obs/geo/recent/notable",
	RecentNearbyObservations:            "data/obs/geo/recent",
	RecentNearbyObservationsOfSpecies:   "data/obs/geo/recent/%s",
	RecentNotableObservationsInRegion:   "data/obs/%s/recent/notable",
	RecentObservationsInRegion:          "data/obs/%s/recent",
	RecentObservationsOfSpeciesInRegion: "data/obs/%s/recent/%s",
	ChecklistFeedOnDate:                 "product/lists/%s/%d/%d/%d",
	RegionalStatisticsOnDate:            "product/stats/%s/%d/%d/%d",
	SpeciesListForRegion:                "product/spplist/%s",
	Top100:                              "product/top100/%s/%d/%d/%d",
	ViewChecklist:                       "product/checklist/view/%s",
	AdjacentRegions:                     "ref/adjacent/%s",
	HotspotInfo:                         "ref/hotspot/info/%s",
	HotspotsInRegion:                    "ref/hotspot/%s",
	NearbyHotspots:                      "ref/hotspot/geo",
	EbirdTaxonomy:                       "ref/taxonomy/ebird",
	TaxaLocaleCodes:                     "ref/taxa-locales/ebird",
	TaxonomicForms:                      "ref/taxon/forms/%s",
	TaxonomicGroups:                     "ref/sppgroup/%s",
	TaxonomyVersions:                    "ref/taxonomy/versions",
	RegionInfo:                          "ref/region/info/%s",
	SubRegionInfo:                       "ref/region/list/%s/%s",
}

type Client struct {
	apikey         string
	baseURL        *url.URL
	httpClient     *http.Client
	acceptLanguage string
}

func WithAcceptLanguage(lang string) ClientOption {
	return func(client *Client) {
		client.acceptLanguage = lang
	}
}

func WithBaseURL(urlStr string) ClientOption {
	return func(client *Client) {
		parsedURL, err := url.Parse(urlStr)
		if err != nil {
			panic(err)
		}
		client.baseURL = parsedURL
	}
}

func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(client *Client) {
		client.httpClient = httpClient
	}
}

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e Error) Error() string {
	return fmt.Sprintf("eBird API error (status %d): %s", e.Status, e.Message)
}

func NewClient(key string, opts ...ClientOption) (*Client, error) {
	if key == "" {
		return nil, errors.New("eBird API key is missing. Please set the 'EBIRD_API_KEY' environment variable")
	}

	c := &Client{
		apikey: key,
		httpClient: &http.Client{
			Timeout: defaultTimeout,
		},
	}

	baseURL, err := url.ParseRequestURI(APIEndpointBase)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}
	c.baseURL = baseURL

	for _, opt := range opts {
		opt(c)
	}

	return c, nil
}

func (c *Client) get(ctx context.Context, endpoint string, params url.Values, result interface{}) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("invalid endpoint URL: %w", err)
	}

	u = c.baseURL.ResolveReference(u)
	u.RawQuery = params.Encode()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, u.String(), nil)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	if c.acceptLanguage != "" {
		req.Header.Set("Accept-Language", c.acceptLanguage)
	}

	req.Header.Set("X-eBirdApiToken", c.apikey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	if err := json.NewDecoder(resp.Body).Decode(result); err != nil {
		return fmt.Errorf("failed to decode response: %w", err)
	}

	return nil
}

func (c *Client) decodeError(resp *http.Response) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read error response: %w", err)
	}

	if len(body) == 0 {
		return fmt.Errorf("HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	var apiError struct {
		Error Error `json:"error"`
	}

	if err := json.Unmarshal(body, &apiError); err != nil {
		return fmt.Errorf("failed to decode error response: %w", err)
	}

	if apiError.Error.Message == "" {
		apiError.Error.Message = fmt.Sprintf("unexpected HTTP %d: %s", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	apiError.Error.Status = resp.StatusCode
	return apiError.Error
}

func convertToJsonFormat(url string) string {
	if !strings.Contains(url, "?") && !strings.Contains(url, "fmt=json") {
		url += "?fmt=json"
	}

	if strings.Contains(url, "?") && !strings.Contains(url, "fmt=json") {
		url += "&fmt=json"
	}

	url = strings.Replace(url, "fmt=csv", "fmt=json", 1)
	return url
}
