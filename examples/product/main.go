package main

import (
	"context"
	"fmt"

	"github.com/siansiansu/go-ebird"
)

const (
	EBIRD_API_KEY = ""
)

var (
	regionCode = "TW"
	y          = 2023
	m          = 10
	d          = 6
)

func main() {
	var ctx = context.Background()
	client, err := ebird.NewClient(EBIRD_API_KEY)
	if err != nil {
		panic(err)
	}

	r, err := client.Top100(ctx, regionCode, y, m, d, ebird.RankedBy("spp"), ebird.MaxResults(10))
	if err != nil {
		panic(err)
	}

	for _, e := range r {
		fmt.Println(e.NumCompleteChecklists, e.NumSpecies, e.UserDisplayName)
	}
}
