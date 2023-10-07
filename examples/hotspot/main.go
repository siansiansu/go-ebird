package main

import (
	"context"
	"fmt"

	"github.com/siansiansu/go-ebird"
)

const (
	EBIRD_API_KEY = ""
)

func main() {
	var ctx = context.Background()
	client, err := ebird.NewClient(EBIRD_API_KEY)
	if err != nil {
		panic(err)
	}

	r, err := client.NearbyHotspots(ctx, ebird.Lat(35), ebird.Lng(137))
	if err != nil {
		panic(err)
	}

	for _, e := range r {
		fmt.Println(e.LocId, e.LocName, e.NumSpeciesAllTime)
	}
}

