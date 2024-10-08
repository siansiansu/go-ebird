package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/siansiansu/go-ebird"
)

const (
	EBIRD_API_KEY = "abc123"
	LATITUDE      = 35.0
	LONGITUDE     = 137.0
)

func main() {
	apiKey := EBIRD_API_KEY
	if apiKey == "" {
		apiKey = os.Getenv("EBIRD_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required. Set EBIRD_API_KEY constant or environment variable.")
		}
	}

	ctx := context.Background()
	client, err := ebird.NewClient(apiKey)
	if err != nil {
		log.Fatalf("Failed to create eBird client: %v", err)
	}

	hotspots, err := client.NearbyHotspots(ctx, ebird.Lat(LATITUDE), ebird.Lng(LONGITUDE))
	if err != nil {
		log.Fatalf("Failed to get nearby hotspots: %v", err)
	}

	fmt.Printf("Nearby hotspots for coordinates (%f, %f):\n", LATITUDE, LONGITUDE)
	for _, hotspot := range hotspots {
		fmt.Printf("- %s (%s): %d species\n", hotspot.LocName, hotspot.LocId, hotspot.NumSpeciesAllTime)
	}
}
