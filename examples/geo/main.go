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
)

func main() {
	var ctx = context.Background()
	client, err := ebird.NewClient(EBIRD_API_KEY)
	if err != nil {
		panic(err)
	}

	r, err := client.AdjacentRegions(ctx, regionCode)
	if err != nil {
		panic(err)
	}

	for _, e := range r {
		fmt.Println(e.Code, e.Name)
	}
}
