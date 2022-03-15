# Name TBD

We'll figure out a pithy description soon.

## Setup

- Clone this repository.
- Clone the front end [repository](https://github.com/esacteksab/wilddata-fe)
- Install [Direnv](https://direnv.net/)
  - Ensure you [hook direnv into your shell](https://direnv.net/docs/hook.html)
  - *DO NOT* put API keys, or any secrets in `.envrc` -- use `.envrc.local` which should be excluded in `.gitignore`
- This uses [gorm datatypes](https://github.com/go-gorm/datatypes)
  - Requires running `go get gorm.io/datatypes` and then running `go build --tags json1`
    - You may also need to run `go mod vendor` once to appease the Go packaging gods.

### Build

- `make build`

### Docker Stuff

A Docker container is pushed to Github Packages via `.github/workflows/docker-publish.yml`. By default the `IMAGE_TAG` is the git SHA of the branch. But you can tag a commit with `git tag -a vX.X.X $SHA` then `git push --tags`. This pushing of tags will trigger GitHub Actions to run the workflow and build a Docker container with an `IMAGE_TAG` of the semver from the tag based on that git commit SHA from above.

## Testing (sort of)

You can use a VSCode extension called [rest-client](https://marketplace.visualstudio.com/items?itemName=humao.rest-client). The `.http` files are located in `REST_API_TEST` directory. This makes it a little easier to do a `GET` or a `POST` and it's possible to share.

- `go run wilddata.go`
