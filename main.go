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
	r, err := client.ViewChecklist(ctx, "S78057631")
	if err != nil {
		panic(err)
	}
	fmt.Println(r.SubmissionMethodCode)
}
