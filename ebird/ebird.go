package ebird

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"reflect"
	"strings"
)

const (
	APIEndpointBase = "https://api.ebird.org/v2"

	// ref/region
	APIEndointRegionInfo    = "ref/region/info/%s"
	APIEndointSubRegionInfo = "ref/region/list/%s/%s"

	// ref/taxonomy
	APIEndpointEbirdTaxonomy    = "ref/taxonomy/ebird"
	APIEndpointTaxaLocaleCodes  = "ref/taxa-locales/ebird"
	APIEndpointTaxonomicForms   = "ref/taxon/forms/%s"
	APIEndpointTaxonomicGroups  = "ref/sppgroup/%s"
	APIEndpointTaxonomyVersions = "ref/taxonomy/versions"

	// ref/hotspot
	APIEndpointHotspotInfo      = "ref/hotspot/info/%s"
	APIEndpointHotspotsInRegion = "ref/hotspot/%s"
	APIEndpointNearbyHotspots   = "ref/hotspot/geo"

	// ref/geo
	APIEndpointAdjacentRegions = "ref/adjacent/%s"

	// product
	APIEndpointChecklistFeedOnDate      = "product/lists/%s/%d/%d/%d"
	APIEndpointRegionalStatisticsOnDate = "product/stats/%s/%d/%d/%d"
	APIEndpointSpeciesListForRegion     = "product/spplist/%s"
	APIEndpointTop100                   = "product/top100/%s/%d/%d/%d"
	APIEndpointViewChecklist            = "product/checklist/view/%s"

	// data/obs
	APIEndpointHistoricObservationsOnDate          = "data/obs/%s/historic/%d/%d/%d"
	APIEndpointNearestObservationsOfSpecies        = "data/nearest/geo/recent/%s"
	APIEndpointRecentChecklistsFeed                = "product/lists/%s"
	APIEndpointRecentNearbyNotableObservations     = "data/obs/geo/recent/notable"
	APIEndpointRecentNearbyObservations            = "data/obs/geo/recent"
	APIEndpointRecentNearbyObservationsOfSpecies   = "data/obs/geo/recent/%s"
	APIEndpointRecentNotableObservationsInRegion   = "data/obs/%s/recent/notable"
	APIEndpointRecentObservationsInRegion          = "data/obs/%s/recent"
	APIEndpointRecentObservationsOfSpeciesInRegion = "data/obs/%s/recent/%s"
)

type Client struct {
	APIKey     string
	BaseURL    *url.URL
	httpClient *http.Client
}

type RequestOptions struct {
	Query   url.Values
	Headers map[string]string
}

func NewClient(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("please input an auth key")
	}
	c := &Client{
		APIKey:     key,
		httpClient: http.DefaultClient,
	}
	if c.BaseURL == nil {
		u, err := url.ParseRequestURI(APIEndpointBase)
		if err != nil {
			return nil, err
		}
		c.BaseURL = u
	}
	return c, nil
}

func (client *Client) url(base *url.URL, endpoint string) string {
	u := *base
	u.Path = path.Join(u.Path, endpoint)
	return u.String()
}

func (c *Client) do(ctx context.Context, req *http.Request) (*http.Response, error) {
	req.Header.Add("x-ebirdapitoken", c.APIKey)
	// req.Header.Set("Content-Type", "application/json")
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.httpClient.Do(req)
}

func (c *Client) get(ctx context.Context, endpoint string, options RequestOptions) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url(c.BaseURL, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if options.Query != nil {
		req.URL.RawQuery = options.Query.Encode()
	}
	if options.Headers != nil {
		for k, v := range options.Headers {
			req.Header.Add(k, v)
		}
	}
	fmt.Println("8===========D", req.URL.RawQuery)
	return c.do(ctx, req)
}

func addOptions(options interface{}) url.Values {
	values := url.Values{}
	if options == nil {
		return values
	}

	optValue := reflect.ValueOf(options)
	if optValue.Kind() != reflect.Struct {
		return values
	}

	optType := optValue.Type()
	for i := 0; i < optType.NumField(); i++ {
		field := optType.Field(i)
		value := optValue.Field(i)
		if value.IsValid() && !reflect.DeepEqual(value.Interface(), reflect.Zero(field.Type).Interface()) {
			tag := field.Tag.Get("url")
			if tag == "" {
				tag = strings.ToLower(field.Name)
			}
			values.Add(tag, fmt.Sprintf("%v", value.Interface()))
		}
	}
	return values
}
