GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	go build -v -o wilddata .

docker:
	docker build -t ghcr.io/esacteksab/wilddata:local .

local:
	make build
	heroku local