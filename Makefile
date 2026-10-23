IMAGE ?= go-htmx
ENGINE ?= podman
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: run test build image image-run

run:
	go run -ldflags="$(LDFLAGS)" .

test:
	go vet ./...
	go test -v ./...

build:
	CGO_ENABLED=0 go build -ldflags="$(LDFLAGS)" -o bin/app .

image:
	$(ENGINE) build -f Containerfile --build-arg VERSION=$(VERSION) -t $(IMAGE) .

image-run:
	$(ENGINE) run --rm -p 8080:8080 $(IMAGE)
