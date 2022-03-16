GO_BUILD_ENV := CGO_ENABLED=0 GOOS=linux GOARCH=amd64

build:
	go build -v -o wilddata .

local:
	make build
	heroku local