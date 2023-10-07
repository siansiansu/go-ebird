# eBird client for Go

## Overview

go-ebird is a Go client library designed for the eBird API. This library enables developers to seamlessly integrate bird observation data from eBird into their Go applications.

## Features

- **Data Retrieval**: Access bird observation data from eBird effortlessly.
- **Flexible Filtering**: Filter data based on parameters such as location, date, and species.
- **Easy Integration**: Integrate eBird data seamlessly into your Go applications.

## Installation

Install `go-ebird` using `go get`:

```shell
go get -u github.com/siansiansu/go-ebird
```

## Usage

```go
package main

import (
  "context"
  "fmt"

  "github.com/siansiansu/go-ebird"
)

func main() {
  var ctx = context.Background()
  client, err := ebird.NewClient("YOUR_EBIRD_API_KEY")
  if err != nil {
    panic(err)
  }
  r, err := client.RecentNotableObservationsInRegion(ctx, "TW")
  if err != nil {
    panic(err)
  }
  for _, e := range r {
    fmt.Println(e.ComName, e.LocName, e.HowMany)
  }
}
```

Remember to replace `"YOUR_EBIRD_API_KEY"` with your actual eBird API key. You can obtain an API key by creating an account on the [eBird website](https://ebird.org/api/keygen).

Refer to the [GoDoc](https://pkg.go.dev/github.com/siansiansu/go-ebird) page for comprehensive documentation and more examples.

## Contributing

Contributions are welcome! Report bugs or request features by opening an issue. If you want to contribute code, fork the repository and submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.