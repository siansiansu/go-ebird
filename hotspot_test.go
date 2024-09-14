package ebird

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHotspotsInRegion(t *testing.T) {
	tests := []struct {
		name       string
		regionCode string
		input      string
		wantResult []HotspotInRegion
		wantErr    bool
		params     []RequestOption
	}{
		{
			name:       "Valid request without params",
			regionCode: "CA-AB",
			input: `[
				{"locId": "L4466534", "locName": "12 Mile Coulee Reservoir", "countryCode": "CA", "subnational1Code": "CA-AB", "subnational2Code": "CA-AB-TW", "lat": 50.2252452, "lng": -111.6423368, "latestObsDt": "2023-09-20 09:00", "numSpeciesAllTime": 80},
				{"locId": "L18410127", "locName": "160 Street Marsh", "countryCode": "CA", "subnational1Code": "CA-AB", "subnational2Code": "CA-AB-SI", "lat": 50.594727, "lng": -113.767492, "latestObsDt": "2023-05-13 14:49", "numSpeciesAllTime": 46}
			]`,
			wantResult: []HotspotInRegion{
				{LocId: "L4466534", LocName: "12 Mile Coulee Reservoir", CountryCode: "CA", Subnational1Code: "CA-AB", Subnational2Code: "CA-AB-TW", Lat: 50.2252452, Lng: -111.6423368, LatestObsDt: "2023-09-20 09:00", NumSpeciesAllTime: 80},
				{LocId: "L18410127", LocName: "160 Street Marsh", CountryCode: "CA", Subnational1Code: "CA-AB", Subnational2Code: "CA-AB-SI", Lat: 50.594727, Lng: -113.767492, LatestObsDt: "2023-05-13 14:49", NumSpeciesAllTime: 46},
			},
			wantErr: false,
		},
		{
			name:       "Valid request with format param",
			regionCode: "CA-AB",
			input: `[
				{"locId": "L4466534", "locName": "12 Mile Coulee Reservoir", "countryCode": "CA", "subnational1Code": "CA-AB", "subnational2Code": "CA-AB-TW", "lat": 50.2252452, "lng": -111.6423368, "latestObsDt": "2023-09-20 09:00", "numSpeciesAllTime": 80}
			]`,
			wantResult: []HotspotInRegion{
				{LocId: "L4466534", LocName: "12 Mile Coulee Reservoir", CountryCode: "CA", Subnational1Code: "CA-AB", Subnational2Code: "CA-AB-TW", Lat: 50.2252452, Lng: -111.6423368, LatestObsDt: "2023-09-20 09:00", NumSpeciesAllTime: 80},
			},
			wantErr: false,
			params:  []RequestOption{Fmt("json")},
		},
		{
			name:       "Empty region code",
			regionCode: "",
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Contains(t, r.URL.Path, tt.regionCode)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.input))
			}))
			defer server.Close()

			client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
			require.NoError(t, err)

			ctx := context.Background()
			got, err := client.HotspotsInRegion(ctx, tt.regionCode, tt.params...)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestNearbyHotspots(t *testing.T) {
	tests := []struct {
		name       string
		input      string
		wantResult []NearbyHotspot
		wantErr    bool
		params     []RequestOption
	}{
		{
			name: "Valid request with lat and lng",
			input: `[
				{"locId": "L1670452", "locName": "Vatnajökulsþjóðgarður NP--Skaftafell", "countryCode": "IS", "subnational1Code": "IS-7", "lat": 64.0172413, "lng": -16.9721603, "latestObsDt": "2023-09-24 10:25", "numSpeciesAllTime": 62},
				{"locId": "L14359747", "locName": "Öræfi--Fagurhólsmýri", "countryCode": "IS", "subnational1Code": "IS-7", "lat": 63.8781272, "lng": -16.6440526, "latestObsDt": "2023-09-23 15:59", "numSpeciesAllTime": 14}
			]`,
			wantResult: []NearbyHotspot{
				{LocId: "L1670452", LocName: "Vatnajökulsþjóðgarður NP--Skaftafell", CountryCode: "IS", Subnational1Code: "IS-7", Lat: 64.0172413, Lng: -16.9721603, LatestObsDt: "2023-09-24 10:25", NumSpeciesAllTime: 62},
				{LocId: "L14359747", LocName: "Öræfi--Fagurhólsmýri", CountryCode: "IS", Subnational1Code: "IS-7", Lat: 63.8781272, Lng: -16.6440526, LatestObsDt: "2023-09-23 15:59", NumSpeciesAllTime: 14},
			},
			wantErr: false,
			params:  []RequestOption{Lat(64), Lng(-17)},
		},
		{
			name:    "Missing lat parameter",
			wantErr: true,
			params:  []RequestOption{Lng(-17)},
		},
		{
			name:    "Missing lng parameter",
			wantErr: true,
			params:  []RequestOption{Lat(64)},
		},
		{
			name:    "No parameters",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.input))
			}))
			defer server.Close()

			client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
			require.NoError(t, err)

			ctx := context.Background()
			got, err := client.NearbyHotspots(ctx, tt.params...)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}

func TestHotspotInfo(t *testing.T) {
	tests := []struct {
		name       string
		locId      string
		input      string
		wantResult *HotspotInfo
		wantErr    bool
	}{
		{
			name:  "Valid request",
			locId: "L1131038",
			input: `{
				"locId": "L1131038",
				"name": "Reykjanes peninsula",
				"latitude": 63.9108121,
				"longitude": -22.5357056,
				"countryCode": "IS",
				"countryName": "Iceland",
				"subnational1Name": "Suðurnes",
				"subnational1Code": "IS-2",
				"isHotspot": true,
				"locName": "Reykjanes peninsula",
				"lat": 63.9108121,
				"lng": -22.5357056,
				"hierarchicalName": "Reykjanes peninsula, Suðurnes, IS",
				"locID": "L1131038"
			}`,
			wantResult: &HotspotInfo{
				LocId: "L1131038", Name: "Reykjanes peninsula", Latitude: 63.9108121, Longitude: -22.5357056,
				CountryCode: "IS", CountryName: "Iceland", Subnational1Name: "Suðurnes", Subnational1Code: "IS-2",
				IsHotspot: true, LocName: "Reykjanes peninsula", Lat: 63.9108121, Lng: -22.5357056,
				HierarchicalName: "Reykjanes peninsula, Suðurnes, IS", LocID: "L1131038",
			},
			wantErr: false,
		},
		{
			name:    "Empty locId",
			locId:   "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				assert.Equal(t, http.MethodGet, r.Method)
				assert.Contains(t, r.URL.Path, tt.locId)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(tt.input))
			}))
			defer server.Close()

			client, err := NewClient("test-api-key", WithBaseURL(server.URL+"/"))
			require.NoError(t, err)

			ctx := context.Background()
			got, err := client.HotspotInfo(ctx, tt.locId)

			if tt.wantErr {
				assert.Error(t, err)
				return
			}

			assert.NoError(t, err)
			assert.Equal(t, tt.wantResult, got)
		})
	}
}
