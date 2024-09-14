package main

import (
	"context"
	"fmt"
	"time"

	"github.com/siansiansu/go-ebird"
)

const (
	EBIRD_API_KEY = "abc123"
)

var (
	regionCode = "TW"
	date       = time.Date(2023, 10, 6, 0, 0, 0, 0, time.UTC)
)

func main() {
	ctx := context.Background()
	client, err := ebird.NewClient(EBIRD_API_KEY)
	if err != nil {
		panic(err)
	}

	r, err := client.Top100(ctx, regionCode, date, ebird.RankedBy("spp"), ebird.MaxResults(10))
	if err != nil {
		panic(err)
	}

	for _, e := range r {
		fmt.Printf("Checklists: %d, Species: %d, User: %s\n", e.NumCompleteChecklists, e.NumSpecies, e.UserDisplayName)
	}
}
