# Barry and Dave's Artifact-Tracking Index Of Doom
We'll figure out a pithy description soon.

## Setup

- Clone this repository.
- This uses [gorm datatypes](https://github.com/go-gorm/datatypes)
  - Requires running `go get gorm.io/datatypes` and then running `go build --tags json1`
    - You may also need to run `go mod vendor` once to appease the Go packaging gods.
- Locally requires two environment variables
  - `export PORT=5000`
  - `export SENTRY_DSN='foo'`

### Build

- `make build`
- `bin/wtfizit`

### Run

- `go run main.go`

## Core features to build (high-level)

## TODO

## DONE
