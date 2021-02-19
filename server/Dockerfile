FROM golang:1.15

# Let's get gRPC source code with `golang.org` blocking bypass.
RUN git clone https://github.com/grpc/grpc-go.git ${GOPATH}/src/google.golang.org/grpc &&\
    git clone https://github.com/protocolbuffers/protobuf-go.git ${GOPATH}/src/google.golang.org/protobuf &&\
    git clone https://github.com/googleapis/go-genproto.git ${GOPATH}/src/google.golang.org/genproto &&\
    git clone https://github.com/golang/protobuf.git ${GOPATH}/src/github.com/golang/protobuf &&\
    git clone https://github.com/golang/text.git ${GOPATH}/src/golang.org/x/text &&\
    git clone https://github.com/golang/net.git ${GOPATH}/src/golang.org/x/net &&\
    git clone https://github.com/golang/sys.git ${GOPATH}/src/golang.org/x/sys

# Fetch application source code and install it.
RUN git clone https://github.com/schwarzlichtbezirk/pds-grpc.git ${GOPATH}/src/github.com/schwarzlichtbezirk/pds-grpc &&\
    go install -v github.com/schwarzlichtbezirk/pds-grpc/server

# gRPC listens ports.
EXPOSE 50051 50052

# Simulate 'Ctrl+C'.
STOPSIGNAL SIGINT

# Run application with full path representation.
ENTRYPOINT ["/go/bin/server"]