# Barry and Dave's Artifact-Tracking Index Of Doom
We'll figure out a pithy description soon.

## Setup

- Clone this repository.
- Install [Direnv](https://direnv.net/)
  - Ensure you [hook direnv into your shell](https://direnv.net/docs/hook.html)
  - *DO NOT* put API keys, or any secrets in `.envrc` -- use `.envrc.local` which should be excluded in `.gitignore`
- This uses [gorm datatypes](https://github.com/go-gorm/datatypes)
  - Requires running `go get gorm.io/datatypes` and then running `go build --tags json1`
    - You may also need to run `go mod vendor` once to appease the Go packaging gods.
- Locally requires two environment variables
  - `export PORT=3000`
  - `export SENTRY_DSN='foo'`

### Build

- `make build`
- `bin/wtfizit`

## Testing (sort of)

You can use a VSCode extension called [rest-client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client). The `.http` files are located in `REST_API_TEST` directory. This makes it a little easier to do a `GET` or a `POST` and it's possible to share.

- `go run main.go`

## Core features to build (high-level)

## TODO

## DONE
