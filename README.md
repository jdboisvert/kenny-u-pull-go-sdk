# Kenny U-Pull Go SDK v0.1.0

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


### Git Workflow

This repo is configured for trunk-based development. When adding a new fix or feature, create a new branch off of `main`.

Merges into main *must always be rebased and squashed*. This can be done manually or with GitHub's "Squash and Merge" feature.

### Testing

All test files are named *_test.go. Github workflow automatically run the tests when code is pushed and will return a report with results when finished.

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

### Badges

* [Go Report Card](https://goreportcard.com/) - It will scan your code with `gofmt`, `go vet`, `gocyclo`, `golint`, `ineffassign`, `license` and `misspell`. Replace `github.com/golang-standards/project-layout` with your project reference.

    [![Go Report Card](https://goreportcard.com/badge/github.com/golang-standards/project-layout?style=flat-square)](https://goreportcard.com/report/github.com/golang-standards/project-layout)

* ~~[GoDoc](http://godoc.org) - It will provide online version of your GoDoc generated documentation. Change the link to point to your project.~~

    [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/golang-standards/project-layout)

* [Pkg.go.dev](https://pkg.go.dev) - Pkg.go.dev is a new destination for Go discovery & docs. You can create a badge using the [badge generation tool](https://pkg.go.dev/badge).

    [![PkgGoDev](https://pkg.go.dev/badge/github.com/golang-standards/project-layout)](https://pkg.go.dev/github.com/golang-standards/project-layout)

* Release - It will show the latest release number for your project. Change the github link to point to your project.

    [![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/golang-standards/project-layout/releases/latest)
