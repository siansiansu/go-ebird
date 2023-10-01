# eBird Go SDK

Country: Taiwan
Name: Min-Sian, Su
Video Demo:
Description: This package provides a golang client for eBird's REST API. Please [sign up](https://ebird.org/api/keygen) for a key on the [eBird](https://ebird.org/home) website.

Notice: This project is still in progress. I plan to incorporate the following features in the future:

- Implementing tests.
- Enhancing flexibility.
- Improving readability.

Example:

```go
package main

import (
  "context"
  "fmt"

  "github.com/siansiansu/go-ebird/ebird"
)

func main() {
  var ctx = context.Background()
  client, err := ebird.NewClient("<your_key>")
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
}
```
