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
	// regionInfoOptions := ebird.RegionInfoOptions{
	// 	RegionNameFormat: "nameonly",
	// 	Delim:            ",",
	// }

	// test 1
	// r, err := client.RegionInfo(ctx, "CA", regionInfoOptions)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(r.Result)

	// test2
	// subRegionListOption := ebird.SubRegionListOptions{
	// 	Fmt: "csv",
	// }
	// r, err := client.SubRegionList(ctx, "country", "world", subRegionListOption)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, e := range r {
	// 	fmt.Println(e.Code)
	// }

	// test3
	// ebirdTaxonomyOptions := ebird.EbirdTaxonomyOptions{
	// 	Locale:  "zh",
	// 	Species: "bkfbun1,coatit2",
	// }
	// r, err := client.EbirdTaxonomy(ctx, ebirdTaxonomyOptions)
	// if err != nil {
	// 	panic(err)
	// }
	// for _, e := range r {
	// 	fmt.Println(e.SciName)
	// }
	// test4

	ebirdTaxonomyOptions := ebird.TaxonomicFormsOptions{
		SpeciesCode: "virrai",
	}
	r, err := client.TaxonomicForms(ctx, ebirdTaxonomyOptions)
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
