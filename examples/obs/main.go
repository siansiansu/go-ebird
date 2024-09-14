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
	REGION_CODE   = "TW"
	MAX_RESULTS   = 2
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

	observations, err := client.RecentObservationsInRegion(ctx, REGION_CODE, ebird.MaxResults(MAX_RESULTS), ebird.Hotspot(true))
	if err != nil {
		log.Fatalf("Failed to get recent observations: %v", err)
	}

	fmt.Printf("Recent observations in %s (max %d, hotspots only):\n", REGION_CODE, MAX_RESULTS)
	for _, obs := range observations {
		fmt.Printf("- %s: %d seen at %s\n", obs.ComName, obs.HowMany, obs.LocName)
	}
}
