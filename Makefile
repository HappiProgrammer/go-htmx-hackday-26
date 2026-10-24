IMAGE ?= go-htmx
ENGINE ?= podman
VERSION ?= $(shell git describe --tags --always --dirty 2>/dev/null || echo dev)
LDFLAGS := -s -w -X main.version=$(VERSION)

.PHONY: run test build image image-run slides slides-pdf

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

# Workshop deck (docs/slides.md) with Marp. Needs Node.js; the PDF also needs Chrome.
MARP := npx -y @marp-team/marp-cli@4.0.3

slides:
	$(MARP) docs/slides.md --preview

slides-pdf:
	$(MARP) docs/slides.md --html --pdf --allow-local-files -o docs/slides.pdf
