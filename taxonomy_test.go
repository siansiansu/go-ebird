package ebird

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTaxonomicGroups(t *testing.T) {
	input := `[
		{"groupName": "Waterfowl", "groupOrder": 1, "taxonOrderBounds": [[1.0, 100.0]]},
		{"groupName": "Raptors", "groupOrder": 2, "taxonOrderBounds": [[101.0, 200.0]]}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/sppgroup/ebird", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.TaxonomicGroups(ctx, "ebird")
	require.NoError(t, err)

	var want []TaxonomicGroup
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestTaxonomyVersions(t *testing.T) {
	input := `[
		{"authorityVer": 2022.0, "latest": true},
		{"authorityVer": 2021.0, "latest": false}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/taxonomy/versions", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.TaxonomyVersions(ctx)
	require.NoError(t, err)

	var want []TaxonomyVersion
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestTaxaLocaleCodes(t *testing.T) {
	input := `[
		{"code": "en", "name": "English", "lastUpdate": "2022-01-01"},
		{"code": "es", "name": "Spanish", "lastUpdate": "2022-01-01"}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/taxa-locales/ebird", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.TaxaLocaleCodes(ctx)
	require.NoError(t, err)

	var want []TaxaLocaleCode
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestTaxonomicForms(t *testing.T) {
	input := `["nominate", "subspecies1", "subspecies2"]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/taxon/forms/amecro", r.URL.Path)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.TaxonomicForms(ctx, "amecro")
	require.NoError(t, err)

	var want []string
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestEbirdTaxonomy(t *testing.T) {
	input := `[
		{
			"sciName": "Corvus brachyrhynchos",
			"comName": "American Crow",
			"speciesCode": "amecro",
			"category": "species",
			"taxonOrder": 22990.0,
			"bandingCodes": ["AMCR"],
			"comNameCodes": [],
			"sciNameCodes": [],
			"order": "Passeriformes",
			"familyCode": "corvi1",
			"familyComName": "Crows, Jays, and Magpies",
			"familySciName": "Corvidae"
		}
	]`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/ref/taxonomy/ebird", r.URL.Path)
		assert.Equal(t, "json", r.URL.Query().Get("fmt"))
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(input))
	}))
	defer server.Close()

	client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
	require.NoError(t, err)

	ctx := context.Background()
	got, err := client.EbirdTaxonomy(ctx)
	require.NoError(t, err)

	var want []EbirdTaxon
	err = json.Unmarshal([]byte(input), &want)
	require.NoError(t, err)

	assert.Equal(t, want, got)
}

func TestTaxonomicGroupsWithEmptyGrouping(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()
	_, err = client.TaxonomicGroups(ctx, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "speciesGrouping cannot be empty")
}

func TestTaxonomicFormsWithEmptySpeciesCode(t *testing.T) {
	client, err := NewClient("test-api-key")
	require.NoError(t, err)

	ctx := context.Background()
	_, err = client.TaxonomicForms(ctx, "")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "speciesCode cannot be empty")
}
