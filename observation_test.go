package ebird

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ObservationTestSuite struct {
	suite.Suite
}

func TestObservationSuites(t *testing.T) {
	suite.Run(t, new(ObservationTestSuite))
}

func (ts *ObservationTestSuite) TestRecentObservationsInRegion() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct without query params",
			Input: `
[
  {
    "speciesCode": "rufgro",
    "comName": "Ruffed Grouse",
    "sciName": "Bonasa umbellus",
    "locId": "L366849",
    "locName": "Kananaskis--Gorge Creek Trail (North Access)",
    "obsDt": "2023-10-05 21:42",
    "howMany": 1,
    "lat": 50.7541732,
    "lng": -114.5556483,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": false,
    "subId": "S151517643"
  },
  {
    "speciesCode": "grhowl",
    "comName": "Great Horned Owl",
    "sciName": "Bubo virginianus",
    "locId": "L27391698",
    "locName": "283112 Township Road 254, Rocky View County CA-AB 51.15436, -113.85223",
    "obsDt": "2023-10-05 19:29",
    "howMany": 1,
    "lat": 51.154359,
    "lng": -113.85223,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151513406"
  }
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "rufgro",
					ComName:         "Ruffed Grouse",
					SciName:         "Bonasa umbellus",
					LocId:           "L366849",
					LocName:         "Kananaskis--Gorge Creek Trail (North Access)",
					ObsDt:           "2023-10-05 21:42",
					HowMany:         1,
					Lat:             50.7541732,
					Lng:             -114.5556483,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: false,
					SubId:           "S151517643",
				},
				{
					SpeciesCode:     "grhowl",
					ComName:         "Great Horned Owl",
					SciName:         "Bubo virginianus",
					LocId:           "L27391698",
					LocName:         "283112 Township Road 254, Rocky View County CA-AB 51.15436, -113.85223",
					ObsDt:           "2023-10-05 19:29",
					HowMany:         1,
					Lat:             51.154359,
					Lng:             -113.85223,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151513406",
				},
			},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "CA-AB"
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.RecentObservationsInRegion(ctx, regionCode)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.RecentObservationsInRegion(ctx, regionCode, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *ObservationTestSuite) TestRecentNotableObservationsInRegion() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct without query params",
			Input: `
[
  {
    "speciesCode": "rebwoo",
    "comName": "Red-bellied Woodpecker",
    "sciName": "Melanerpes carolinus",
    "locId": "L3089074",
    "locName": "Nacmine (community)",
    "obsDt": "2023-10-05 17:46",
    "howMany": 1,
    "lat": 51.4708064,
    "lng": -112.7871633,
    "obsValid": true,
    "obsReviewed": true,
    "locationPrivate": false,
    "subId": "S151515981"
},
{
    "speciesCode": "rebwoo",
    "comName": "Red-bellied Woodpecker",
    "sciName": "Melanerpes carolinus",
    "locId": "L3089074",
    "locName": "Nacmine (community)",
    "obsDt": "2023-10-05 17:46",
    "howMany": 1,
    "lat": 51.4708064,
    "lng": -112.7871633,
    "obsValid": true,
    "obsReviewed": true,
    "locationPrivate": false,
    "subId": "S151511421"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "rebwoo",
					ComName:         "Red-bellied Woodpecker",
					SciName:         "Melanerpes carolinus",
					LocId:           "L3089074",
					LocName:         "Nacmine (community)",
					ObsDt:           "2023-10-05 17:46",
					HowMany:         1,
					Lat:             51.4708064,
					Lng:             -112.7871633,
					ObsValid:        true,
					ObsReviewed:     true,
					LocationPrivate: false,
					SubId:           "S151515981",
				},
				{
					SpeciesCode:     "rebwoo",
					ComName:         "Red-bellied Woodpecker",
					SciName:         "Melanerpes carolinus",
					LocId:           "L3089074",
					LocName:         "Nacmine (community)",
					ObsDt:           "2023-10-05 17:46",
					HowMany:         1,
					Lat:             51.4708064,
					Lng:             -112.7871633,
					ObsValid:        true,
					ObsReviewed:     true,
					LocationPrivate: false,
					SubId:           "S151511421",
				},
			},
			ExpErr: nil,
			Params: nil,
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "CA-AB"
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.RecentNotableObservationsInRegion(ctx, regionCode)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.RecentNotableObservationsInRegion(ctx, regionCode, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *ObservationTestSuite) TestRecentObservationsOfSpeciesInRegion() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct without query params",
			Input: `
[
  {
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L17449098",
    "locName": "Home",
    "obsDt": "2023-10-05 18:34",
    "howMany": 2,
    "lat": 38.572064,
    "lng": -90.420457,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508867",
    "exoticCategory": "N"
},
{
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L21993870",
    "locName": "Mattese Meadows",
    "obsDt": "2023-10-05 18:20",
    "howMany": 2,
    "lat": 38.475536,
    "lng": -90.35024,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508532",
    "exoticCategory": "N"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L17449098",
					LocName:         "Home",
					ObsDt:           "2023-10-05 18:34",
					HowMany:         2,
					Lat:             38.572064,
					Lng:             -90.420457,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508867",
					ExoticCategory:  "N",
				},
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L21993870",
					LocName:         "Mattese Meadows",
					ObsDt:           "2023-10-05 18:20",
					HowMany:         2,
					Lat:             38.475536,
					Lng:             -90.35024,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508532",
					ExoticCategory:  "N",
				},
			},
			ExpErr: nil,
			Params: nil,
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "US"
		specidesCode := "eutspa"
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.RecentObservationsOfSpeciesInRegion(ctx, regionCode, specidesCode)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.RecentObservationsOfSpeciesInRegion(ctx, regionCode, specidesCode, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *ObservationTestSuite) TestRecentNearbyObservations() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params []RequestOption
	}{
		{
			Desc: "ensure that the response is correct with query params",
			Input: `
[
  {
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L17449098",
    "locName": "Home",
    "obsDt": "2023-10-05 18:34",
    "howMany": 2,
    "lat": 38.572064,
    "lng": -90.420457,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508867",
    "exoticCategory": "N"
},
{
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L21993870",
    "locName": "Mattese Meadows",
    "obsDt": "2023-10-05 18:20",
    "howMany": 2,
    "lat": 38.475536,
    "lng": -90.35024,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508532",
    "exoticCategory": "N"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L17449098",
					LocName:         "Home",
					ObsDt:           "2023-10-05 18:34",
					HowMany:         2,
					Lat:             38.572064,
					Lng:             -90.420457,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508867",
					ExoticCategory:  "N",
				},
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L21993870",
					LocName:         "Mattese Meadows",
					ObsDt:           "2023-10-05 18:20",
					HowMany:         2,
					Lat:             38.475536,
					Lng:             -90.35024,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508532",
					ExoticCategory:  "N",
				},
			},
			ExpErr: nil,
			Params: []RequestOption{Lat(64), Lng(-17)},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		ctx := context.Background()
		result, err := client.RecentNearbyObservations(ctx, test.Params...)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}

func (ts *ObservationTestSuite) TestRecentNearbyObservationsOfSpecies() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params []RequestOption
	}{
		{
			Desc: "ensure that the response is correct with query params",
			Input: `
[
  {
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L17449098",
    "locName": "Home",
    "obsDt": "2023-10-05 18:34",
    "howMany": 2,
    "lat": 38.572064,
    "lng": -90.420457,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508867",
    "exoticCategory": "N"
},
{
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L21993870",
    "locName": "Mattese Meadows",
    "obsDt": "2023-10-05 18:20",
    "howMany": 2,
    "lat": 38.475536,
    "lng": -90.35024,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151508532",
    "exoticCategory": "N"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L17449098",
					LocName:         "Home",
					ObsDt:           "2023-10-05 18:34",
					HowMany:         2,
					Lat:             38.572064,
					Lng:             -90.420457,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508867",
					ExoticCategory:  "N",
				},
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L21993870",
					LocName:         "Mattese Meadows",
					ObsDt:           "2023-10-05 18:20",
					HowMany:         2,
					Lat:             38.475536,
					Lng:             -90.35024,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151508532",
					ExoticCategory:  "N",
				},
			},
			ExpErr: nil,
			Params: []RequestOption{Lat(64), Lng(-17)},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		ctx := context.Background()
		result, err := client.RecentNearbyObservations(ctx, test.Params...)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}

func (ts *ObservationTestSuite) TestNearestObservationsOfSpecies() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params []RequestOption
	}{
		{
			Desc: "ensure that the response is correct with query params",
			Input: `
[
  {
    "speciesCode": "eutspa",
    "comName": "Eurasian Tree Sparrow",
    "sciName": "Passer montanus",
    "locId": "L21181579",
    "locName": "名古屋市--勅使池  (Nagoya--Chokushi-ike Park)",
    "obsDt": "2023-09-24 08:00",
    "howMany": 2,
    "lat": 35.083182,
    "lng": 137.013097,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": false,
    "subId": "S150855601"
},
{
  "speciesCode": "eutspa",
  "comName": "Eurasian Tree Sparrow",
  "sciName": "Passer montanus",
  "locId": "L17164894",
  "locName": "勅使池",
  "obsDt": "2023-09-23 07:52",
  "howMany": 6,
  "lat": 35.083271,
  "lng": 137.013995,
  "obsValid": true,
  "obsReviewed": false,
  "locationPrivate": true,
  "subId": "S150527663"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L21181579",
					LocName:         "名古屋市--勅使池  (Nagoya--Chokushi-ike Park)",
					ObsDt:           "2023-09-24 08:00",
					HowMany:         2,
					Lat:             35.083182,
					Lng:             137.013097,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: false,
					SubId:           "S150855601",
				},
				{
					SpeciesCode:     "eutspa",
					ComName:         "Eurasian Tree Sparrow",
					SciName:         "Passer montanus",
					LocId:           "L17164894",
					LocName:         "勅使池",
					ObsDt:           "2023-09-23 07:52",
					HowMany:         6,
					Lat:             35.083271,
					Lng:             137.013995,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S150527663",
				},
			},
			ExpErr: nil,
			Params: []RequestOption{Lat(64), Lng(-17)},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		speciesCode := "eutspa"

		ctx := context.Background()
		result, err := client.NearestObservationsOfSpecies(ctx, speciesCode, test.Params...)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}

func (ts *ObservationTestSuite) TestRecentNearbyNotableObservations() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params []RequestOption
	}{
		{
			Desc: "ensure that the response is correct with query params",
			Input: `
[
	{
    "speciesCode": "masboo",
    "comName": "Masked Booby",
    "sciName": "Sula dactylatra",
    "locId": "L27325351",
    "locName": "岐阜市--上佐波西９丁目の水田 (Gifu--Kamisabanishi 9-Chome ricefields)",
    "obsDt": "2023-10-01 15:50",
    "howMany": 1,
    "lat": 35.3662473,
    "lng": 136.7154931,
    "obsValid": false,
    "obsReviewed": false,
    "locationPrivate": false,
    "subId": "S151144322"
},
{
  "speciesCode": "corplo",
  "comName": "Common Ringed Plover",
  "sciName": "Charadrius hiaticula",
  "locId": "L18620654",
  "locName": "楠町南五味塚1545-1, 四日市市 JP-三重県 34.90576, 136.64237",
  "obsDt": "2023-09-28 14:29",
  "howMany": 1,
  "lat": 34.905757,
  "lng": 136.642367,
  "obsValid": true,
  "obsReviewed": true,
  "locationPrivate": true,
  "subId": "S150910536"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "masboo",
					ComName:         "Masked Booby",
					SciName:         "Sula dactylatra",
					LocId:           "L27325351",
					LocName:         "岐阜市--上佐波西９丁目の水田 (Gifu--Kamisabanishi 9-Chome ricefields)",
					ObsDt:           "2023-10-01 15:50",
					HowMany:         1,
					Lat:             35.3662473,
					Lng:             136.7154931,
					ObsValid:        false,
					ObsReviewed:     false,
					LocationPrivate: false,
					SubId:           "S151144322",
				},
				{
					SpeciesCode:     "corplo",
					ComName:         "Common Ringed Plover",
					SciName:         "Charadrius hiaticula",
					LocId:           "L18620654",
					LocName:         "楠町南五味塚1545-1, 四日市市 JP-三重県 34.90576, 136.64237",
					ObsDt:           "2023-09-28 14:29",
					HowMany:         1,
					Lat:             34.905757,
					Lng:             136.642367,
					ObsValid:        true,
					ObsReviewed:     true,
					LocationPrivate: true,
					SubId:           "S150910536",
				},
			},
			ExpErr: nil,
			Params: []RequestOption{Lat(64), Lng(-17)},
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		ctx := context.Background()
		result, err := client.RecentNearbyNotableObservations(ctx, test.Params...)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}

func (ts *ObservationTestSuite) TestRecentChecklistsFeed() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []RecentChecklistsFeed
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct",
			Input: `
[
  {
  "locId": "L13557069",
  "subId": "S151523092",
  "userDisplayName": "Ed Bailey",
  "numSpecies": 1,
  "obsDt": "6 Oct 2023",
  "obsTime": "03:35",
  "isoObsDate": "2023-10-06 03:35",
  "subID": "S151523092",
  "loc": {
    "locId": "L13557069",
    "name": "my yard",
    "latitude": 41.3242778,
    "longitude": -73.1988271,
    "countryCode": "US",
    "countryName": "United States",
    "subnational1Name": "Connecticut",
    "subnational1Code": "US-CT",
    "subnational2Code": "US-CT-001",
    "subnational2Name": "Fairfield",
    "isHotspot": false,
    "locID": "L13557069",
    "locName": "my yard",
    "lat": 41.3242778,
    "lng": -73.1988271,
    "hierarchicalName": "my yard, Fairfield, Connecticut, US"
  }
},
{
  "locId": "L27393260",
  "subId": "S151521762",
  "userDisplayName": "Eric Fay",
  "numSpecies": 1,
  "obsDt": "6 Oct 2023",
  "obsTime": "03:28",
  "isoObsDate": "2023-10-06 03:28",
  "subID": "S151521762",
  "loc": {
    "locId": "L27393260",
    "name": "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391",
    "latitude": 43.563074,
    "longitude": -70.213914,
    "countryCode": "US",
    "countryName": "United States",
    "subnational1Name": "Maine",
    "subnational1Code": "US-ME",
    "subnational2Code": "US-ME-005",
    "subnational2Name": "Cumberland",
    "isHotspot": false,
    "locID": "L27393260",
    "locName": "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391",
    "lat": 43.563074,
    "lng": -70.213914,
    "hierarchicalName": "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391, Cumberland, Maine, US"
  }
}
]`,
			ExpRes: []RecentChecklistsFeed{
				{
					LocId:           "L13557069",
					SubId:           "S151523092",
					UserDisplayName: "Ed Bailey",
					NumSpecies:      1,
					ObsDt:           "6 Oct 2023",
					ObsTime:         "03:35",
					IsoObsDate:      "2023-10-06 03:35",
					SubID:           "S151523092",
					Loc: Loc{
						LocId:            "L13557069",
						Name:             "my yard",
						Latitude:         41.3242778,
						Longitude:        -73.1988271,
						CountryCode:      "US",
						CountryName:      "United States",
						Subnational1Name: "Connecticut",
						Subnational1Code: "US-CT",
						Subnational2Code: "US-CT-001",
						Subnational2Name: "Fairfield",
						IsHotspot:        false,
						LocID:            "L13557069",
						LocName:          "my yard",
						Lat:              41.3242778,
						Lng:              -73.1988271,
						HierarchicalName: "my yard, Fairfield, Connecticut, US",
					},
				},
				{
					LocId:           "L27393260",
					SubId:           "S151521762",
					UserDisplayName: "Eric Fay",
					NumSpecies:      1,
					ObsDt:           "6 Oct 2023",
					ObsTime:         "03:28",
					IsoObsDate:      "2023-10-06 03:28",
					SubID:           "S151521762",
					Loc: Loc{
						LocId:            "L27393260",
						Name:             "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391",
						Latitude:         43.563074,
						Longitude:        -70.213914,
						CountryCode:      "US",
						CountryName:      "United States",
						Subnational1Name: "Maine",
						Subnational1Code: "US-ME",
						Subnational2Code: "US-ME-005",
						Subnational2Name: "Cumberland",
						IsHotspot:        false,
						LocID:            "L27393260",
						LocName:          "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391",
						Lat:              43.563074,
						Lng:              -70.213914,
						HierarchicalName: "91 Two Lights Rd, Cape Elizabeth US-ME 43.56307, -70.21391, Cumberland, Maine, US",
					},
				},
			},
			ExpErr: nil,
			Params: MaxResults(10),
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "US"
		ctx := context.Background()
		result, err := client.RecentChecklistsFeed(ctx, regionCode, test.Params)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}

func (ts *ObservationTestSuite) TestHistoricObservationsOnDate() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Observations
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct",
			Input: `
[
  {
    "speciesCode": "grhowl",
    "comName": "Great Horned Owl",
    "sciName": "Bubo virginianus",
    "locId": "L4122129",
    "locName": "Hays Home",
    "obsDt": "2023-09-30 23:40",
    "howMany": 2,
    "lat": 35.1337729,
    "lng": -120.5566585,
    "obsValid": true,
    "obsReviewed": false,
    "locationPrivate": true,
    "subId": "S151140860"
},
{
  "speciesCode": "killde",
  "comName": "Killdeer",
  "sciName": "Charadrius vociferus",
  "locId": "L5236749",
  "locName": "Fairmont Hill",
  "obsDt": "2023-09-30 23:38",
  "howMany": 1,
  "lat": 33.8672523,
  "lng": -117.7810312,
  "obsValid": true,
  "obsReviewed": false,
  "locationPrivate": true,
  "subId": "S151169205"
}
]`,
			ExpRes: []Observations{
				{
					SpeciesCode:     "grhowl",
					ComName:         "Great Horned Owl",
					SciName:         "Bubo virginianus",
					LocId:           "L4122129",
					LocName:         "Hays Home",
					ObsDt:           "2023-09-30 23:40",
					HowMany:         2,
					Lat:             35.1337729,
					Lng:             -120.5566585,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151140860",
				},
				{
					SpeciesCode:     "killde",
					ComName:         "Killdeer",
					SciName:         "Charadrius vociferus",
					LocId:           "L5236749",
					LocName:         "Fairmont Hill",
					ObsDt:           "2023-09-30 23:38",
					HowMany:         1,
					Lat:             33.8672523,
					Lng:             -117.7810312,
					ObsValid:        true,
					ObsReviewed:     false,
					LocationPrivate: true,
					SubId:           "S151169205",
				},
			},
			ExpErr: nil,
			Params: MaxResults(10),
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "US-CA"
		y := 2023
		m := 10
		d := 6
		ctx := context.Background()
		result, err := client.HistoricObservationsOnDate(ctx, regionCode, y, m, d, test.Params)
		if err != nil {
			ts.EqualError(test.ExpErr, err.Error(), test.Desc)
		} else {
			ts.Equal(test.ExpRes, result, test.Desc)
		}
	}
}
