# syntax=docker/dockerfile:1

# Build application from golang image
FROM golang:1.22.5 AS build-stage

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ARG TARGETOS
ARG TARGETARCH

RUN CGO_ENABLED=0 GOOS=${TARGETOS} GOARCH=${TARGETARCH} go build -o /go/bin/cve-processor ./cmd/main.go

# Deploy the application binary into a lean image
FROM scratch AS build-release-stage

COPY --from=build-stage /go/bin/cve-processor /go/bin/cve-processor
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
WORKDIR /go
ENTRYPOINT ["/go/bin/cve-processor"]
