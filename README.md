# go-ebird: eBird API Client for Go

[![GoDoc](https://godoc.org/github.com/siansiansu/go-ebird?status.svg)](http://godoc.org/github.com/siansiansu/go-ebird)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Go Report Card](https://goreportcard.com/badge/github.com/siansiansu/go-ebird)](https://goreportcard.com/report/github.com/siansiansu/go-ebird)

## Overview

go-ebird is a comprehensive Go client library for the eBird API. It provides an easy-to-use interface for developers to integrate bird observation data from eBird into their Go applications.

## Features

- **Comprehensive Data Access**: Retrieve various types of bird observation data, including recent observations, notable sightings, and checklists.
- **Flexible Filtering**: Filter data based on location, date, species, and more.
- **Robust Error Handling**: Clear error messages and proper handling of API rate limits.
- **Customizable Client**: Configure timeout, base URL, and other HTTP client options.
- **Full API Coverage**: Support for all major eBird API endpoints.

## Installation

Install go-ebird using `go get`:

```shell
go get -u github.com/siansiansu/go-ebird
```

## Quick Start

Here's a simple example to get you started:

```go
package main

import (
    "context"
    "fmt"
    "log"

    "github.com/siansiansu/go-ebird"
)

func main() {
    client, err := ebird.NewClient("YOUR_EBIRD_API_KEY")
    if err != nil {
        log.Fatalf("Failed to create client: %v", err)
    }

    observations, err := client.RecentObservationsInRegion(context.Background(), "US-NY", ebird.MaxResults(5))
    if err != nil {
        log.Fatalf("Failed to get observations: %v", err)
    }

    for _, obs := range observations {
        fmt.Printf("%s spotted at %s\n", obs.ComName, obs.LocName)
    }
}
```

## Detailed Usage

Check out the [examples directory](./examples/) for more detailed usage examples, including:

- Retrieving notable observations
- Getting nearby hotspots
- Fetching recent checklists
- Accessing taxonomy information

## API Endpoints

go-ebird supports all major eBird API endpoints, including:

- Observations
- Hotspots
- Taxonomy
- Checklists
- Region information

Refer to the [GoDoc](https://pkg.go.dev/github.com/siansiansu/go-ebird) for a complete list of supported endpoints and their usage.

## Configuration

You can configure the client with various options:

```go
client, err := ebird.NewClient(
    "YOUR_EBIRD_API_KEY",
    ebird.WithBaseURL("https://api.ebird.org/v2/"),
    ebird.WithHTTPClient(&http.Client{Timeout: 30 * time.Second}),
    ebird.WithAcceptLanguage("en"),
)
```

## Rate Limiting

The eBird API has rate limits. This client does not automatically handle rate limiting, so be sure to implement appropriate backoff and retry logic in your application.

## Contributing

Contributions are welcome! Here's how you can contribute:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

Please ensure your code adheres to the existing style and includes appropriate tests.

## Testing

Run the tests using:

```shell
go test -v ./...
```

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

## Acknowledgments

- Thanks to the eBird team for providing the API
- Inspired by other excellent Go API clients

## Support

If you encounter any issues or have questions, please [open an issue](https://github.com/siansiansu/go-ebird/issues/new) on GitHub.