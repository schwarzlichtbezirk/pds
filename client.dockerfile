##
## Build stage
##

# Use image with golang last version as builder.
FROM golang:1.23-bullseye AS build

# See https://stackoverflow.com/questions/64462922/docker-multi-stage-build-go-image-x509-certificate-signed-by-unknown-authorit
RUN apt-get update && apt-get install -y ca-certificates openssl
ARG cert_location=/usr/local/share/ca-certificates
# Get certificate from "github.com".
RUN openssl s_client -showcerts -connect github.com:443 </dev/null 2>/dev/null|openssl x509 -outform PEM > ${cert_location}/github.crt
# Get certificate from "proxy.golang.org".
RUN openssl s_client -showcerts -connect proxy.golang.org:443 </dev/null 2>/dev/null|openssl x509 -outform PEM >  ${cert_location}/proxy.golang.crt
# Update certificates.
RUN update-ca-certificates

# Make project root folder as current dir.
WORKDIR $GOPATH/src/github.com/schwarzlichtbezirk/pds/
# Copy only go.mod and go.sum to prevent downloads all dependencies again on any code changes.
COPY go.mod go.sum ./
# Download all dependencies pointed at go.mod file.
RUN go mod download
# Copy all files and subfolders in current state as is.
COPY . .

# Build service and put executable.
RUN go build -o /go/bin/client -ldflags="-X 'main.buildvers=`cat semver`' -X 'main.builddate=$(date +%F)'" ./client

##
## Deploy stage
##

# Thin deploy image.
FROM gcr.io/distroless/base-debian11

# Copy compiled executables to new image destination.
COPY --from=build /go/bin/client /go/bin/client
# Copy configuration files.
COPY --from=build /go/src/github.com/schwarzlichtbezirk/pds/config/* /go/bin/config/

# Open REST listen port.
EXPOSE 8008

# Run application with full path representation.
# Without shell to get signal for graceful shutdown.
ENTRYPOINT ["/go/bin/client"]
