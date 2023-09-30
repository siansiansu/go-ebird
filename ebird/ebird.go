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
	APITaxonomicForms       = "ref/taxon/forms/%s"
	APIEbirdTaxonomy        = "ref/taxonomy/ebird"
	APIEndointRegionInfo    = "ref/region/info/%s"
	APIEndointSubRegionInfo = "ref/region/list/%s/%s"
	APIEndpointBase         = "https://api.ebird.org/v2"
)

type Client struct {
	APIKey     string
	BaseURL    *url.URL
	httpClient *http.Client
}

func NewClient(key string) (*Client, error) {
	if key == "" {
		return nil, errors.New("wololo")
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

func (c *Client) get(ctx context.Context, endpoint string, query url.Values) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.url(c.BaseURL, endpoint), nil)
	if err != nil {
		return nil, err
	}
	if query != nil {
		req.URL.RawQuery = query.Encode()
	}
	fmt.Println("8============D", req.URL, req.URL.RawQuery)
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
