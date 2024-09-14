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

	regionInfo, err := client.RegionInfo(ctx, REGION_CODE, ebird.RegionNameFormat("detailed"))
	if err != nil {
		log.Fatalf("Failed to get region info: %v", err)
	}

	fmt.Printf("Region Information for %s:\n", REGION_CODE)
	fmt.Printf("Result: %s\n", regionInfo.Result)
	fmt.Printf("Code: %s\n", regionInfo.Code)
	fmt.Printf("Type: %s\n", regionInfo.Type)
	fmt.Printf("Latitude: %f\n", regionInfo.Latitude)
	fmt.Printf("Longitude: %f\n", regionInfo.Longitude)
	fmt.Printf("Bounds: MinX=%f, MaxX=%f, MinY=%f, MaxY=%f\n",
		regionInfo.Bounds.MinX, regionInfo.Bounds.MaxX, regionInfo.Bounds.MinY, regionInfo.Bounds.MaxY)
}
