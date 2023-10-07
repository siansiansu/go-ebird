package ebird

import (
	"bytes"
	"context"
	"net/http"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ProductTestSuite struct {
	suite.Suite
}

func TestProductSuites(t *testing.T) {
	suite.Run(t, new(ProductTestSuite))
}

func (ts *ProductTestSuite) TestTop100() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []Top100
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct",
			Input: `
[
  {
		"profileHandle": "MTE3NDA2NA",
		"userDisplayName": "Rajan Rao",
		"numSpecies": 132,
		"numCompleteChecklists": 0,
		"rowNum": 1,
		"userId": "USER1174064"
  },
  {
		"profileHandle": "MTIzMjAy",
		"userDisplayName": "Aaron Haiman",
		"numSpecies": 132,
		"numCompleteChecklists": 0,
		"rowNum": 1,
		"userId": "USER123202"
  }
]`,
			ExpRes: []Top100{
				{
					ProfileHandle:         "MTE3NDA2NA",
					UserDisplayName:       "Rajan Rao",
					NumSpecies:            132,
					NumCompleteChecklists: 0,
					RowNum:                1,
					UserId:                "USER1174064",
				},
				{
					ProfileHandle:         "MTIzMjAy",
					UserDisplayName:       "Aaron Haiman",
					NumSpecies:            132,
					NumCompleteChecklists: 0,
					RowNum:                1,
					UserId:                "USER123202",
				},
			},
			ExpErr: nil,
			Params: RankedBy("spp"),
		},
	}
	for _, test := range tests {
		client, server := testClient(http.StatusOK, bytes.NewBufferString(test.Input))
		defer server.Close()
		client.httpClient = server.Client()

		regionCode := "CA-AB"
		y := 2023
		m := 10
		d := 6
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.Top100(ctx, regionCode, y, m, d)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.Top100(ctx, regionCode, y, m, d, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}

func (ts *ProductTestSuite) TestChecklistFeedOnDate() {
	tests := []struct {
		Desc   string
		Input  string
		ExpRes []ChecklistFeedOnDate
		ExpErr error
		Params RequestOption
	}{
		{
			Desc: "ensure that the response is correct",
			Input: `
[
  {
		"locId": "L4122129",
		"subId": "S151140860",
		"userDisplayName": "Mark Hays",
		"numSpecies": 1,
		"obsDt": "30 Sep 2023",
		"obsTime": "23:40",
		"isoObsDate": "2023-09-30 23:40",
		"subID": "S151140860",
		"loc": {
				"locId": "L4122129",
				"name": "Hays Home",
				"latitude": 35.1337729,
				"longitude": -120.5566585,
				"countryCode": "US",
				"countryName": "United States",
				"subnational1Name": "California",
				"subnational1Code": "US-CA",
				"subnational2Code": "US-CA-079",
				"subnational2Name": "San Luis Obispo",
				"isHotspot": false,
				"locName": "Hays Home",
				"lat": 35.1337729,
				"lng": -120.5566585,
				"hierarchicalName": "Hays Home, San Luis Obispo, California, US",
				"locID": "L4122129"
		}
  },
  {
		"locId": "L24844801",
		"subId": "S151140446",
		"userDisplayName": "Sarah Brown",
		"numSpecies": 1,
		"obsDt": "30 Sep 2023",
		"obsTime": "23:38",
		"isoObsDate": "2023-09-30 23:38",
		"subID": "S151140446",
		"loc": {
				"locId": "L24844801",
				"name": "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165",
				"latitude": 37.382722,
				"longitude": -122.081653,
				"countryCode": "US",
				"countryName": "United States",
				"subnational1Name": "California",
				"subnational1Code": "US-CA",
				"subnational2Code": "US-CA-085",
				"subnational2Name": "Santa Clara",
				"isHotspot": false,
				"locName": "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165",
				"lat": 37.382722,
				"lng": -122.081653,
				"hierarchicalName": "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165, Santa Clara, California, US",
				"locID": "L24844801"
		}
  }
]`,
			ExpRes: []ChecklistFeedOnDate{
				{
					LocId:           "L4122129",
					SubId:           "S151140860",
					UserDisplayName: "Mark Hays",
					NumSpecies:      1,
					ObsDt:           "30 Sep 2023",
					ObsTime:         "23:40",
					IsoObsDate:      "2023-09-30 23:40",
					SubID:           "S151140860",
					Loc: HotspotInfo{
						LocId:            "L4122129",
						Name:             "Hays Home",
						Latitude:         35.1337729,
						Longitude:        -120.5566585,
						CountryCode:      "US",
						CountryName:      "United States",
						Subnational1Name: "California",
						Subnational1Code: "US-CA",
						Subnational2Code: "US-CA-079",
						Subnational2Name: "San Luis Obispo",
						IsHotspot:        false,
						LocName:          "Hays Home",
						Lat:              35.1337729,
						Lng:              -120.5566585,
						HierarchicalName: "Hays Home, San Luis Obispo, California, US",
						LocID:            "L4122129",
					},
				},
				{
					LocId:           "L24844801",
					SubId:           "S151140446",
					UserDisplayName: "Sarah Brown",
					NumSpecies:      1,
					ObsDt:           "30 Sep 2023",
					ObsTime:         "23:38",
					IsoObsDate:      "2023-09-30 23:38",
					SubID:           "S151140446",
					Loc: HotspotInfo{
						LocId:            "L24844801",
						Name:             "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165",
						Latitude:         37.382722,
						Longitude:        -122.081653,
						CountryCode:      "US",
						CountryName:      "United States",
						Subnational1Name: "California",
						Subnational1Code: "US-CA",
						Subnational2Code: "US-CA-085",
						Subnational2Name: "Santa Clara",
						IsHotspot:        false,
						LocName:          "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165",
						Lat:              37.382722,
						Lng:              -122.081653,
						HierarchicalName: "1005 Boranda Ave, Mountain View US-CA 37.38272, -122.08165, Santa Clara, California, US",
						LocID:            "L24844801",
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

		regionCode := "CA-AB"
		y := 2023
		m := 10
		d := 6
		ctx := context.Background()
		if test.Params == nil {
			result, err := client.ChecklistFeedOnDate(ctx, regionCode, y, m, d)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		} else {
			result, err := client.ChecklistFeedOnDate(ctx, regionCode, y, m, d, test.Params)
			if err != nil {
				ts.EqualError(test.ExpErr, err.Error(), test.Desc)
			} else {
				ts.Equal(test.ExpRes, result, test.Desc)
			}
		}
	}
}
