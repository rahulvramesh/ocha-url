GOCMD=go
BINARY_NAME=ocha
VERSION?=1.0
.PHONY: all test build vendor

## Build:
build-release: vendor
	 env GIN_MODE=release GO111MODULE=on $(GOCMD) build -mod vendor -o build/bin/$(BINARY_NAME) cmd/ocha/main.go

clean:
	rm -fr ./build
	rm -f ./junit-report.xml checkstyle-report.xml ./coverage.xml ./profile.cov yamllint-checkstyle.xml

vendor: ## Copy of all packages needed to support builds and tests in the vendor directory
	$(GOCMD) mod vendor
