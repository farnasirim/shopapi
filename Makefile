GIT ?= git
GO ?= go
COMMIT := $(shell $(GIT) rev-parse HEAD)
VERSION ?= $(shell $(GIT) describe --abbrev=0 --tags 2>/dev/null)
BUILD_TIME := $(shell LANG=en_US date +"%F_%T_%z")
TARGET := github.com/farnasirim/shopapi/cmd/shopapi
ROOT := github.com/farnasirim/shopapi/cmd/shopapi
LD_FLAGS := -X $(ROOT).Version=$(VERSION) -X $(ROOT).BuildTime=$(BUILD_TIME) -X $(ROOT).Commit=$(COMMIT)
DOCKER ?= docker
IMAGE_NAME ?= colonelmo/shopapi
IMAGE_VERSION ?= $(VERSION)

help:
	@echo "use 'make <target>' where <target> is one of the following"
	@echo "  shopapi to build the main binary"
	@echo "  docker to build the docker image"
	@echo "  test-dependencies to get the test dependencies"
	@echo "  test to run tests"
	@echo "  clean to remove generated files"

shopapi: *.go */*.go */*/*.go generate
	$(GO) build -o="shopapi" -ldflags="$(LD_FLAGS)" $(TARGET)

generate: schema mocks

schema: api/graphql/schema.graphql
	$(GO) generate api/graphql/schema.go

mocks: shopapi.go
	$(GO) generate shopapi.go


docker: shopapi 
	$(DOCKER) build -t $(IMAGE_NAME):$(IMAGE_VERSION) .

push: docker
	$(DOCKER) push $(IMAGE_NAME):$(VERSION)

dependencies:
	$(GO) list -f='{{ join .Deps "\n" }}' $(TARGET) | grep -v $(ROOT) | tr '\n' ' ' | xargs $(GO) get -u -v
	$(GO) list -f='{{ join .Deps "\n" }}' $(TARGET) | grep -v $(ROOT) | tr '\n' ' ' | xargs $(GO) install -v

test-dependencies: dependencies
	$(GO) list -f='{{ join .TestImports "\n" }}' ./... | grep -v $(ROOT) | tr '\n' ' ' | xargs $(GO) get -u -v

test: $(shell find | grep _test.go)
	$(GO) test ./...

cover-data: $(shell find | grep _test.go)
	rm -rf .cover
	mkdir .cover
	for pkg in $$($(GO) list ./...) ; do f=".cover/$$(echo $$pkg | tr / -).cover"; $(GO) test -covermode=count -coverprofile="$$f" "$$pkg" ; done
	echo "mode: count" >.cover/cover.out
	grep -h -v "^mode:" .cover/*.cover >>".cover/cover.out"

coverage-report: cover-data
	$(GO) tool cover -func=".cover/cover.out"

clean:
	rm -f shopapi

