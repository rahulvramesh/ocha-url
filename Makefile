GOCMD=go
BINARY_NAME=ocha
VERSION?=1.0
.PHONY: all test build vendor

## Build:
build-release: vendor swagger
	 env GIN_MODE=release GO111MODULE=on $(GOCMD) build -mod vendor -o build/bin/$(BINARY_NAME) cmd/ocha/main.go

clean:
	rm -fr ./build
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml

vendor:
	$(GOCMD) mod vendor

unit:
	go test --tags=unit ./...

integration:
	 go test --tags=integration ./...

test: vendor
	 go test -v ./...

build-docker: swagger
	docker-compose -f deployments/docker-compose.yml build

docker-run: swagger
	docker-compose -f deployments/docker-compose.yml up

swagger:
	swag init -g cmd/ocha/main.go -o ./api