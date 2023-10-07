package ebird

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type HostSpotTestSuite struct {
	suite.Suite
}

func TestHotspotsSuites(t *testing.T) {
	suite.Run(t, new(HostSpotTestSuite))
}

func (ts *HostSpotTestSuite) TestHotspotsInRegion() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []HotspotsInRegion
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct without query params",
			Input: `
[
  {
    "locId": "L4466534",
    "locName": "12 Mile Coulee Reservoir",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-TW",
    "lat": 50.2252452,
    "lng": -111.6423368,
    "latestObsDt": "2023-09-20 09:00",
    "numSpeciesAllTime": 80
},
{
    "locId": "L18410127",
    "locName": "160 Street Marsh",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-SI",
    "lat": 50.594727,
    "lng": -113.767492,
    "latestObsDt": "2023-05-13 14:49",
    "numSpeciesAllTime": 46
}
]`,
			ExpRes: []HotspotsInRegion{
				{
					LocId:             "L4466534",
					LocName:           "12 Mile Coulee Reservoir",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-TW",
					Lat:               50.2252452,
					Lng:               -111.6423368,
					LatestObsDt:       "2023-09-20 09:00",
					NumSpeciesAllTime: 80,
				},
				{
					LocId:             "L18410127",
					LocName:           "160 Street Marsh",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-SI",
					Lat:               50.594727,
					Lng:               -113.767492,
					LatestObsDt:       "2023-05-13 14:49",
					NumSpeciesAllTime: 46,
				},
			},
			ExpErr: nil,
			Params: nil,
		}, {
			Desc: "ensure that the response is correct with csv format",
			Input: `
[
  {
    "locId": "L4466534",
    "locName": "12 Mile Coulee Reservoir",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-TW",
    "lat": 50.2252452,
    "lng": -111.6423368,
    "latestObsDt": "2023-09-20 09:00",
    "numSpeciesAllTime": 80
},
{
    "locId": "L18410127",
    "locName": "160 Street Marsh",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-SI",
    "lat": 50.594727,
    "lng": -113.767492,
    "latestObsDt": "2023-05-13 14:49",
    "numSpeciesAllTime": 46
}
]`,
			ExpRes: []HotspotsInRegion{
				{
					LocId:             "L4466534",
					LocName:           "12 Mile Coulee Reservoir",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-TW",
					Lat:               50.2252452,
					Lng:               -111.6423368,
					LatestObsDt:       "2023-09-20 09:00",
					NumSpeciesAllTime: 80,
				},
				{
					LocId:             "L18410127",
					LocName:           "160 Street Marsh",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-SI",
					Lat:               50.594727,
					Lng:               -113.767492,
					LatestObsDt:       "2023-05-13 14:49",
					NumSpeciesAllTime: 46,
				},
			},
			ExpErr: nil,
			Params: Fmt("csv"),
		}, {
			Desc: "ensure that the response is correct with json format",
			Input: `
[
  {
    "locId": "L4466534",
    "locName": "12 Mile Coulee Reservoir",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-TW",
    "lat": 50.2252452,
    "lng": -111.6423368,
    "latestObsDt": "2023-09-20 09:00",
    "numSpeciesAllTime": 80
},
{
    "locId": "L18410127",
    "locName": "160 Street Marsh",
    "countryCode": "CA",
    "subnational1Code": "CA-AB",
    "subnational2Code": "CA-AB-SI",
    "lat": 50.594727,
    "lng": -113.767492,
    "latestObsDt": "2023-05-13 14:49",
    "numSpeciesAllTime": 46
}
]`,
			ExpRes: []HotspotsInRegion{
				{
					LocId:             "L4466534",
					LocName:           "12 Mile Coulee Reservoir",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-TW",
					Lat:               50.2252452,
					Lng:               -111.6423368,
					LatestObsDt:       "2023-09-20 09:00",
					NumSpeciesAllTime: 80,
				},
				{
					LocId:             "L18410127",
					LocName:           "160 Street Marsh",
					CountryCode:       "CA",
					Subnational1Code:  "CA-AB",
					Subnational2Code:  "CA-AB-SI",
					Lat:               50.594727,
					Lng:               -113.767492,
					LatestObsDt:       "2023-05-13 14:49",
					NumSpeciesAllTime: 46,
				},
			},
			ExpErr: nil,
			Params: Fmt("json"),
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "CA-AB"
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.HotspotsInRegion(ctx, regionCode)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.HotspotsInRegion(ctx, regionCode, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *HostSpotTestSuite) TestNearbyHotspots() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []NearbyHotspots
		ExpErr error
		Params []RequestOption
	}{
		{
			Desc: "ensure that the response is correct with lat and lng parameters",
			Input: `
[
	{
    "locId": "L1670452",
    "locName": "Vatnajökulsþjóðgarður NP--Skaftafell",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 64.0172413,
    "lng": -16.9721603,
    "latestObsDt": "2023-09-24 10:25",
    "numSpeciesAllTime": 62
  },
	{
    "locId": "L14359747",
    "locName": "Öræfi--Fagurhólsmýri",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 63.8781272,
    "lng": -16.6440526,
    "latestObsDt": "2023-09-23 15:59",
    "numSpeciesAllTime": 14
	}
]`,
			ExpRes: []NearbyHotspots{
				{
					LocId:             "L1670452",
					LocName:           "Vatnajökulsþjóðgarður NP--Skaftafell",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               64.0172413,
					Lng:               -16.9721603,
					LatestObsDt:       "2023-09-24 10:25",
					NumSpeciesAllTime: 62,
				},
				{
					LocId:             "L14359747",
					LocName:           "Öræfi--Fagurhólsmýri",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               63.8781272,
					Lng:               -16.6440526,
					LatestObsDt:       "2023-09-23 15:59",
					NumSpeciesAllTime: 14,
				},
			},
			ExpErr: nil,
			Params: []RequestOption{Lat(64), Lng(-17)},
		},
		{
			Desc: "ensure that the response is correct with lat parameter",
			Input: `
[
	{
    "locId": "L1670452",
    "locName": "Vatnajökulsþjóðgarður NP--Skaftafell",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 64.0172413,
    "lng": -16.9721603,
    "latestObsDt": "2023-09-24 10:25",
    "numSpeciesAllTime": 62
  },
	{
    "locId": "L14359747",
    "locName": "Öræfi--Fagurhólsmýri",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 63.8781272,
    "lng": -16.6440526,
    "latestObsDt": "2023-09-23 15:59",
    "numSpeciesAllTime": 14
	}
]`,
			ExpRes: []NearbyHotspots{
				{
					LocId:             "L1670452",
					LocName:           "Vatnajökulsþjóðgarður NP--Skaftafell",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               64.0172413,
					Lng:               -16.9721603,
					LatestObsDt:       "2023-09-24 10:25",
					NumSpeciesAllTime: 62,
				},
				{
					LocId:             "L14359747",
					LocName:           "Öræfi--Fagurhólsmýri",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               63.8781272,
					Lng:               -16.6440526,
					LatestObsDt:       "2023-09-23 15:59",
					NumSpeciesAllTime: 14,
				},
			},
			ExpErr: fmt.Errorf("must provide Lat and Lng parameters"),
			Params: []RequestOption{Lat(64)},
		},
		{
			Desc: "ensure that the response is correct with lng parameter",
			Input: `
[
	{
    "locId": "L1670452",
    "locName": "Vatnajökulsþjóðgarður NP--Skaftafell",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 64.0172413,
    "lng": -16.9721603,
    "latestObsDt": "2023-09-24 10:25",
    "numSpeciesAllTime": 62
  },
	{
    "locId": "L14359747",
    "locName": "Öræfi--Fagurhólsmýri",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 63.8781272,
    "lng": -16.6440526,
    "latestObsDt": "2023-09-23 15:59",
    "numSpeciesAllTime": 14
	}
]`,
			ExpRes: []NearbyHotspots{
				{
					LocId:             "L1670452",
					LocName:           "Vatnajökulsþjóðgarður NP--Skaftafell",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               64.0172413,
					Lng:               -16.9721603,
					LatestObsDt:       "2023-09-24 10:25",
					NumSpeciesAllTime: 62,
				},
				{
					LocId:             "L14359747",
					LocName:           "Öræfi--Fagurhólsmýri",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               63.8781272,
					Lng:               -16.6440526,
					LatestObsDt:       "2023-09-23 15:59",
					NumSpeciesAllTime: 14,
				},
			},
			ExpErr: fmt.Errorf("must provide Lat and Lng parameters"),
			Params: []RequestOption{Lng(-17)},
		},
		{
			Desc: "ensure that the response is correct without any parameter",
			Input: `
[
	{
    "locId": "L1670452",
    "locName": "Vatnajökulsþjóðgarður NP--Skaftafell",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 64.0172413,
    "lng": -16.9721603,
    "latestObsDt": "2023-09-24 10:25",
    "numSpeciesAllTime": 62
  },
	{
    "locId": "L14359747",
    "locName": "Öræfi--Fagurhólsmýri",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 63.8781272,
    "lng": -16.6440526,
    "latestObsDt": "2023-09-23 15:59",
    "numSpeciesAllTime": 14
	}
]`,
			ExpRes: []NearbyHotspots{
				{
					LocId:             "L1670452",
					LocName:           "Vatnajökulsþjóðgarður NP--Skaftafell",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               64.0172413,
					Lng:               -16.9721603,
					LatestObsDt:       "2023-09-24 10:25",
					NumSpeciesAllTime: 62,
				},
				{
					LocId:             "L14359747",
					LocName:           "Öræfi--Fagurhólsmýri",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               63.8781272,
					Lng:               -16.6440526,
					LatestObsDt:       "2023-09-23 15:59",
					NumSpeciesAllTime: 14,
				},
			},
			ExpErr: fmt.Errorf("must provide Lat and Lng parameters"),
			Params: nil,
		},
		{
			Desc: "ensure that the response is correct with csv parameter",
			Input: `
[
	{
    "locId": "L1670452",
    "locName": "Vatnajökulsþjóðgarður NP--Skaftafell",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 64.0172413,
    "lng": -16.9721603,
    "latestObsDt": "2023-09-24 10:25",
    "numSpeciesAllTime": 62
  },
	{
    "locId": "L14359747",
    "locName": "Öræfi--Fagurhólsmýri",
    "countryCode": "IS",
    "subnational1Code": "IS-7",
    "lat": 63.8781272,
    "lng": -16.6440526,
    "latestObsDt": "2023-09-23 15:59",
    "numSpeciesAllTime": 14
	}
]`,
			ExpRes: []NearbyHotspots{
				{
					LocId:             "L1670452",
					LocName:           "Vatnajökulsþjóðgarður NP--Skaftafell",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               64.0172413,
					Lng:               -16.9721603,
					LatestObsDt:       "2023-09-24 10:25",
					NumSpeciesAllTime: 62,
				},
				{
					LocId:             "L14359747",
					LocName:           "Öræfi--Fagurhólsmýri",
					CountryCode:       "IS",
					Subnational1Code:  "IS-7",
					Lat:               63.8781272,
					Lng:               -16.6440526,
					LatestObsDt:       "2023-09-23 15:59",
					NumSpeciesAllTime: 14,
				},
			},
			ExpErr: fmt.Errorf("must provide Lat and Lng parameters"),
			Params: []RequestOption{Lat(64), Lng(-17), Fmt("csv")},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		ctx := context.Background()
		if test.Params == nil {
			result, err := client.NearbyHotspots(ctx)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.NearbyHotspots(ctx, test.Params...)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *HostSpotTestSuite) TestHotspotInfo() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes *HotspotInfo
		ExpErr error
	}{
		{
			Desc: "ensure that the response is correct",
			Input: `
{
    "locId": "L1131038",
    "name": "Reykjanes peninsula (please consider using more specific locations)",
    "latitude": 63.9108121,
    "longitude": -22.5357056,
    "countryCode": "IS",
    "countryName": "Iceland",
    "subnational1Name": "Suðurnes",
    "subnational1Code": "IS-2",
    "isHotspot": true,
    "locName": "Reykjanes peninsula (please consider using more specific locations)",
    "lat": 63.9108121,
    "lng": -22.5357056,
    "hierarchicalName": "Reykjanes peninsula (please consider using more specific locations), Suðurnes, IS",
    "locID": "L1131038"
}
			`,
			ExpRes: &HotspotInfo{
				LocId:            "L1131038",
				Name:             "Reykjanes peninsula (please consider using more specific locations)",
				Latitude:         63.9108121,
				Longitude:        -22.5357056,
				CountryCode:      "IS",
				CountryName:      "Iceland",
				Subnational1Name: "Suðurnes",
				Subnational1Code: "IS-2",
				IsHotspot:        true,
				LocName:          "Reykjanes peninsula (please consider using more specific locations)",
				Lat:              63.9108121,
				Lng:              -22.5357056,
				HierarchicalName: "Reykjanes peninsula (please consider using more specific locations), Suðurnes, IS",
				LocID:            "L1131038",
			},
			ExpErr: nil,
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		ctx := context.Background()
		locId := "L1131038"
		result, err := client.HotspotInfo(ctx, locId)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}
