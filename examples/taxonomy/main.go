package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/siansiansu/go-ebird"
)

const (
	EBIRD_API_KEY = "abc123"
	LOCALE        = "zh"
	MAX_RESULTS   = 10
	TIMEOUT       = 60 * time.Second
)

func main() {
	apiKey := EBIRD_API_KEY
	if apiKey == "" {
		apiKey = os.Getenv("EBIRD_API_KEY")
		if apiKey == "" {
			log.Fatal("API key is required. Set EBIRD_API_KEY constant or environment variable.")
		}
	}

	httpClient := &http.Client{
		Timeout: TIMEOUT,
	}

	ctx := context.Background()
	client, err := ebird.NewClient(apiKey, ebird.WithHTTPClient(httpClient))
	if err != nil {
		log.Fatalf("Failed to create eBird client: %v", err)
	}

	ctx, cancel := context.WithTimeout(ctx, TIMEOUT)
	defer cancel()

	taxonomy, err := client.EbirdTaxonomy(ctx, ebird.Fmt("json"), ebird.Locale(LOCALE))
	if err != nil {
		log.Fatalf("Failed to get eBird taxonomy: %v", err)
	}

	fmt.Printf("eBird Taxonomy (Locale: %s):\n", LOCALE)
	for i, taxon := range taxonomy {
		if i >= MAX_RESULTS {
			fmt.Printf("... and %d more entries\n", len(taxonomy)-MAX_RESULTS)
			break
		}
		fmt.Printf("%d. %s (Common Name: %s, Species Code: %s)\n", i+1, taxon.SciName, taxon.ComName, taxon.SpeciesCode)
	}
}
