package ebird

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const (
	// Base URL for eBird API 2.0 requests.
	// baseURL should always be specified with a trailing slash.
	APIEndpointBase = "https://api.ebird.org/v2/"

	// The data/obs end-points are used to fetch observations submitted to eBird in checklists.
	// https://documenter.getpostman.com/view/664302/S1ENwy59#4e020bc2-fc67-4fb6-a926-570cedefcc34
	APIEndpointHistoricObservationsOnDate          = "data/obs/%s/historic/%d/%d/%d"
	APIEndpointNearestObservationsOfSpecies        = "data/nearest/geo/recent/%s"
	APIEndpointRecentChecklistsFeed                = "product/lists/%s"
	APIEndpointRecentNearbyNotableObservations     = "data/obs/geo/recent/notable"
	APIEndpointRecentNearbyObservations            = "data/obs/geo/recent"
	APIEndpointRecentNearbyObservationsOfSpecies   = "data/obs/geo/recent/%s"
	APIEndpointRecentNotableObservationsInRegion   = "data/obs/%s/recent/notable"
	APIEndpointRecentObservationsInRegion          = "data/obs/%s/recent"
	APIEndpointRecentObservationsOfSpeciesInRegion = "data/obs/%s/recent/%s"

	// The product end-points make it easy to get the information shown in various pages on the eBird web site
	// https://documenter.getpostman.com/view/664302/S1ENwy59#af04604f-e406-4cea-991c-a9baef24cd78
	APIEndpointChecklistFeedOnDate      = "product/lists/%s/%d/%d/%d"
	APIEndpointRegionalStatisticsOnDate = "product/stats/%s/%d/%d/%d"
	APIEndpointSpeciesListForRegion     = "product/spplist/%s"
	APIEndpointTop100                   = "product/top100/%s/%d/%d/%d"
	APIEndpointViewChecklist            = "product/checklist/view/%s"

	// ref/geo
	// With the ref/geo end-point you can find a country's or region's neighbours.
	// https://documenter.getpostman.com/view/664302/S1ENwy59#c9947c5c-2dce-4c6d-9911-7d702235506c
	APIEndpointAdjacentRegions = "ref/adjacent/%s"

	// ref/hotspot
	// With the ref/hotspot end-points you can find the hotspots for a given country or region or nearby hotspots
	// https://documenter.getpostman.com/view/664302/S1ENwy59#c9947c5c-2dce-4c6d-9911-7d702235506c
	APIEndpointHotspotInfo      = "ref/hotspot/info/%s"
	APIEndpointHotspotsInRegion = "ref/hotspot/%s"
	APIEndpointNearbyHotspots   = "ref/hotspot/geo"

	// ref/taxonomy
	// https://documenter.getpostman.com/view/664302/S1ENwy59#36c95b76-e18e-4788-9c9e-e539045f9166
	APIEndpointEbirdTaxonomy    = "ref/taxonomy/ebird"
	APIEndpointTaxaLocaleCodes  = "ref/taxa-locales/ebird"
	APIEndpointTaxonomicForms   = "ref/taxon/forms/%s"
	APIEndpointTaxonomicGroups  = "ref/sppgroup/%s"
	APIEndpointTaxonomyVersions = "ref/taxonomy/versions"

	// ref/region
	// The ref/region end-points return information on regions.
	// https://documenter.getpostman.com/view/664302/S1ENwy59#e18ea3b5-e80c-479f-87db-220ce8d9f3b6
	APIEndointRegionInfo    = "ref/region/info/%s"
	APIEndointSubRegionInfo = "ref/region/list/%s/%s"
)

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

type Error struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func (e Error) Error() string {
	return e.Message
}

func NewClient(key string, opts ...ClientOption) (*Client, error) {
	if key == "" {
		return nil, errors.New("eBird API key is missing. Please set the 'EBIRD_API_KEY' environment variable")
	}

	c := &Client{
		apikey:     key,
		httpClient: http.DefaultClient,
	}

	if c.baseURL == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.baseURL = u
	}

	for _, opt := range opts {
		opt(c)
	}
	return c, nil
}

func (c *Client) get(ctx context.Context, query string, result interface{}) error {
	ebirdURL := c.baseURL.String() + query
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, ebirdURL, nil)
	if err != nil {
		return err
	}

	if c.acceptLanguage != "" {
		req.Header.Set("Accept-Language", c.acceptLanguage)
	}

	req.Header.Add("x-ebirdapitoken", c.apikey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNoContent {
		return nil
	}

	if resp.StatusCode != http.StatusOK {
		return c.decodeError(resp)
	}

	err = json.NewDecoder(resp.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}

func (c *Client) decodeError(resp *http.Response) error {
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if len(responseBody) == 0 {
		return fmt.Errorf("HTTP %d: %s (body empty)", resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	buf := bytes.NewBuffer(responseBody)

	var e struct {
		E Error `json:"error"`
	}
	err = json.NewDecoder(buf).Decode(&e)
	if err != nil {
		return fmt.Errorf("couldn't decode error: (%d) [%s]", len(responseBody), responseBody)
	}

	if e.E.Message == "" {
		e.E.Message = fmt.Sprintf("unexpected HTTP %d: %s (empty error)",
			resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return e.E
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
