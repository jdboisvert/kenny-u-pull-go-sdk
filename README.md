# Kenny U-Pull Go SDK v0.1.2

A Collection of useful Go functions and utilities to interact with Kenny U-Pull's API.

## Usage

### Installation

    go get github.com/jdboisvert/kenny-u-pull-go-sdk

### Functions and Usage

#### GetInventory
Used to get a list of inventory listings based on the search criteria provided from the Kenny U-Pull Website/API.
This will include all the information about the vehicle and the branch it is located at.
##### Example
```go
package main

import (
	"fmt"

	kennyupull "github.com/jdboisvert/kenny-u-pull-go-sdk"
)

func main() {
	inventorySearch := kennyupull.InventorySearch{
		Year:   "2010",
		Make:   "toyota",
		Model:  "corolla",
		Branch: "all-branches", // Optional by default will be set to "all-branches"
	}
	inventoryListings, err := kennyupull.GetInventory(inventorySearch)
	if err != nil {
		panic(err)
	}
	fmt.Println(inventoryListings) // Should display all the InventoryListenings matching your search criteria
}

```

## Development

### Getting Started

    # install golang
    brew install golang

    # install the golangci linter
    # more details: https://golangci-lint.run/
    brew install golangci-lint

    # install pre-commit
    pip install pre-commit
    pre-commit install

    # Download all dependencies
    go mod download

### Pre-commit

A number of pre-commit hooks are set up to ensure all commits meet basic code quality standards.

If one of the hooks changes a file, you will need to `git add` that file and re-run `git commit` before being able to continue.


### Testing

All test files are named *_test.go. Github workflow automatically run the tests when code is pushed and will return a report with results when finished.

You can also run the tests locally:

    go test ./...

### PRs and Releases

GitHub Actions is configured to perform unit tests for all new PRs.

It will also check if the version has been bumped. To do that, use `bumpver`:

    # "patch" bumps are for minor non-breaking changes, hotfixes,
    # documentation updates, new tests, etc.
    bumpver update --patch
    # or
    bumper update -p

    # "minor" bumps are for significant backwards-compatible changes
    bumpver update --minor
    # or
    bumper update -m

    # "major" bumps are for breaking changes
    bumpver update --major


[![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Go Reference](https://pkg.go.dev/badge/github.com/jdboisvert/kenny-u-pull-go-sdk.svg)](https://pkg.go.dev/github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/jdboisvert/kenny-u-pull-go-sdk/releases/latest)
