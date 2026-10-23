# Stage 1: Build static Go binary with embedded HTMX templates.
# The builder runs on the build machine's own architecture and Go cross-compiles
# for each target (amd64, arm64), so multi-arch builds need no emulation.
FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder
WORKDIR /src
# go.sum* glob: a stdlib-only module has no go.sum, and a plain COPY would fail.
COPY go.mod go.sum* ./
RUN go mod download
COPY . .
# VERSION is shown in the page footer; CI passes the short commit SHA.
ARG VERSION=dev
ARG TARGETOS
ARG TARGETARCH
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w -X main.version=${VERSION}" -o /bin/app .

# Stage 2: Minimal rootless Distroless runtime
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=builder /bin/app /app
USER nonroot:nonroot
EXPOSE 8080
ENTRYPOINT ["/app"]
