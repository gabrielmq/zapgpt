.PHONY: build clean deploy

build:
	export GO111MODULE=on
	env GOARCH=amd64 GOOS=linux go build -ldflags="-s -w" -o bin/zapgpt cmd/lambda/main.go

clean:
	rm -rf ./bin ./vendor

deploy: clean build
	sls deploy -s dev --verbose
