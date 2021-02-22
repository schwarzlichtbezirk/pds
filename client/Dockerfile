############
# 1 stage. #
############

FROM golang:1.15 AS build

# Let's get gRPC and application source code with `golang.org` blocking bypass.
# Then get application source code with others dependencies.
RUN git clone https://github.com/grpc/grpc-go.git ${GOPATH}/src/google.golang.org/grpc &&\
    git clone https://github.com/protocolbuffers/protobuf-go.git ${GOPATH}/src/google.golang.org/protobuf &&\
    git clone https://github.com/googleapis/go-genproto.git ${GOPATH}/src/google.golang.org/genproto &&\
    git clone https://github.com/golang/protobuf.git ${GOPATH}/src/github.com/golang/protobuf &&\
    git clone https://github.com/golang/text.git ${GOPATH}/src/golang.org/x/text &&\
    git clone https://github.com/golang/net.git ${GOPATH}/src/golang.org/x/net &&\
    git clone https://github.com/golang/sys.git ${GOPATH}/src/golang.org/x/sys &&\
    go get -d -v github.com/gorilla/mux &&\
    git clone https://github.com/schwarzlichtbezirk/pds-grpc.git ${GOPATH}/src/github.com/schwarzlichtbezirk/pds-grpc

# Install application.
COPY . ${GOPATH}/src/github.com/schwarzlichtbezirk/pds-grpc/client
RUN go install -v github.com/schwarzlichtbezirk/pds-grpc/client

# REST listen port.
EXPOSE 8008

# Ping web service to check health.
HEALTHCHECK --interval=15s --timeout=3s\
    CMD curl -f -d "healthcheck" http://localhost:8008/api/tool/ping || exit 1

# Simulate 'Ctrl+C'.
STOPSIGNAL SIGINT

# Run application with full path representation.
ENTRYPOINT ["/go/bin/client"]

############
# 2 stage. #
############

FROM ubuntu:latest AS prod

# Install toolset.
RUN apt-get update && apt-get install -y curl

# Copy compiled executables to new image destination.
COPY --from=build /go/bin/ /go/bin/
# Copy configuration files.
COPY --from=build /go/src/github.com/schwarzlichtbezirk/pds-grpc/config/ /go/bin/config/

# REST listen port.
EXPOSE 8008

# Ping web service to check health.
HEALTHCHECK --interval=15s --timeout=3s\
    CMD curl -f -d "healthcheck" http://localhost:8008/api/tool/ping || exit 1

# Simulate 'Ctrl+C', it's brings to graceful shutdown on all platforms.
# Because golang:1.13 and previous does not have SIGTERM implementation for windows.
STOPSIGNAL SIGINT

# Run application with full path representation.
ENTRYPOINT ["/go/bin/client"]
