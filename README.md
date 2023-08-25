# Kenny U-Pull Go SDK v0.2.1

A Collection of useful Go functions and utilities to interact with Kenny U-Pull's API.

## Usage

### Installation

    go get github.com/jdboisvert/kenny-u-pull-go-sdk

### How to use

Please refer to the [examples](./docs/examples.md) for more information on how to use the SDK.

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

To run the tests with coverage:

    go test -coverprofile=coverage.out ./...
    go tool cover -html=coverage.out

### PRs and Releases

GitHub Actions is configured to perform unit tests for all new PRs.

[![Go Report Card](https://goreportcard.com/badge/github.com/jdboisvert/kenny-u-pull-go-sdk)](https://goreportcard.com/report/github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](https://github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Go Reference](https://pkg.go.dev/badge/github.com/jdboisvert/kenny-u-pull-go-sdk.svg)](https://pkg.go.dev/github.com/jdboisvert/kenny-u-pull-go-sdk)

[![Release](https://img.shields.io/github/release/golang-standards/project-layout.svg?style=flat-square)](https://github.com/jdboisvert/kenny-u-pull-go-sdk/releases/latest)
