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

	regions, err := client.AdjacentRegions(ctx, REGION_CODE)
	if err != nil {
		log.Fatalf("Failed to get adjacent regions: %v", err)
	}

	fmt.Printf("Adjacent regions for %s:\n", REGION_CODE)
	for _, region := range regions {
		fmt.Printf("- %s (%s)\n", region.Name, region.Code)
	}
}
