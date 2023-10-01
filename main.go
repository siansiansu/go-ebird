package main

import (
	"context"
	"fmt"

	"github.com/siansiansu/go-ebird/ebird"
)

func main() {
	var ctx = context.Background()
	client, err := ebird.NewClient("694uloho5a0b")
	if err != nil {
		panic(err)
	}

	r, err := client.HistoricObservationsOnDate(ctx, "TW", 2023, 9, 30, ebird.RecentObservationsOptions{
		MaxResults: 1,
	})
	if err != nil {
		panic(err)
	}
	for _, e := range r {
		fmt.Println(e.ComName)
	}

	data, err := client.ViewChecklist(ctx, "S78057631")
	if err != nil {
		panic(err)
	}
	fmt.Println(data.ChecklistId, data.NumSpecies, data.LocId)
}
