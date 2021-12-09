# Port Domain Servece

Client and server test sample with gRPC streaming and REST functionality. Sample features:

- REST service.
- gRPC service.
- load balancing for gRPC backend.
- uniform code style with documentation.
- errors checking at all source code points.
- domain driven design.
- single application configuration.
- safely servers start and graceful shutdown.
- runnable at pure host, at standalone containers, at composite bundle.

## Source code structure

### client

- `main.go` is file with main function only, contains workflow functions calls.
- `workflow.go` contains functions for initialization, starts of services, wait for break, and finalization function with graceful shutdown of services.
- `config.go`, all settings of application are collected into single structure with single initialization. This singleton can be streamed into JSON or YAML file.
- `router.go` have a routing for HTTP-server, and some auxiliary functions for HTTP handlers.
- `handlers.go` contains the list of HTTP handlers and error codes for them.
- `io.go` reads settings from configuration file. Reads `port.json` file with predefined data format, and sends items step-by-step to gRPC server. File does not limited by size.
- `envfmt.go` have helper function to expand environment variables in the file path.

### server

- `main.go` is file with main function only, contains workflow functions calls.
- `workflow.go` contains functions for initialization, starts of services, wait for break, and finalization function with graceful shutdown of services.
- `config.go`, all settings of application are collected into single structure with single initialization. This singleton can be streamed into JSON or YAML file.
- `grpcserv.go` have gRPC interface implementation for server.
- `io.go` reads settings from configuration file.

### pds

Here is `pds.proto` with gRPC interface declaration, and files produced by protobuf compiler.

## How to run on localhost

1. First of all install [Golang](https://golang.org/) of last version. Requires that [GOPATH is set](https://golang.org/doc/code.html#GOPATH).

2. Fetch golang `grpc` library.

```batch
go get -u google.golang.org/grpc
```

Note: if there is no access to `golang.org` host, use VPN (via Netherlands/USA) or git repositories cloning.

3. Fetch this source code and compile application.

```batch
go get github.com/schwarzlichtbezirk/pds
```

Folder `github.com\schwarzlichtbezirk\pds\tool` contains batch helpers to compile services for Windows for x86 and amd64 platforms. Also it has shell-scripts to compile for Linux amd64 platforms.

4. Edit config-files `github.com/schwarzlichtbezirk/pds/config/client.yaml` and `github.com/schwarzlichtbezirk/pds/config/server.yaml`.

5. Run services.

```batch
start "PDS server" %GOPATH%/bin/pds.server.x64.exe
start "PDS client" %GOPATH%/bin/pds.client.x64.exe
```

or run `github.com\schwarzlichtbezirk\pds\tool\start.x64.cmd` batch-file to start composition of client and server.

## Connections

Ports thats used at network are defined in configuration of server and client (see files `config.go`).

Server listen on `50051` and `50052` ports by default, and it can be a list for load balancing.

Client creates connection to gRPC server on the same ports. There is used `round_robin` load balancer policy. Host can be defined by environment variable `SERVERURL`, and if it not defined or empty, `localhost` is used. Also client opens `8008` port by default to listen for incoming connections to serve REST API, and it can be a list for load balancing.

On localhost server and client can be run as is without any modifications in configuration.

## How to run in docker

1. Change current directory to project root.

```batch
cd /d %GOPATH%/src/github.com/schwarzlichtbezirk/dfs
```

2. Build docker images for `server` and for `client` services.

```batch
docker build --pull --rm -f "Dockerfile.server" -t dfs-server:latest "."
docker build --pull --rm -f "Dockerfile.client" -t dfs-client:latest "."
```

3. Then run docker compose file.

```batch
docker-compose -f "docker-compose.yaml" up -d --build
```

### Run standalone containers

Create a network, its created only once.

```batch
docker network create -d bridge --subnet 172.20.0.0/16 pds-net
```

Then it should be run containers on `pds-net` network.

```batch
docker run --rm -d -p 50051:50051 -p 50052:50052 --network=pds-net --ip=172.20.1.7 --name server pds-server
docker run --rm -d -p 8008:8008 --network=pds-net --ip=172.20.1.8 -e SERVERURL="172.20.1.7" --name client pds-client
```

### Run by docker compose file

Docker compose file uses already builded images and creates internal network for containers.

## What its need else to modify code

If you want to modify `.go`-code and `.proto` file, you should [download](https://github.com/protocolbuffers/protobuf/blob/master/README.md#protocol-compiler-installation) and install protocol buffer compiler. Then fetch and compile protocol buffer compiler plugins:

```batch
go get -u google.golang.org/protobuf
go get -u google.golang.org/genproto
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
go install google.golang.org/protobuf/cmd/protoc-gen-go
```

To generate protocol buffer code, run `tool/pb.cmd` batch file.

## REST API

Arguments of all API calls placed as JSON-objects at request body. Replies comes also only as JSON-objects in all cases.

Errors come on replies with status >= 300 as objects like `{"what":"some error message","when":1613251727492,"code":3}` where `when` is Unix time in milliseconds of error occurrence, `code` is unique error source point code.

### Store port object `/api/port/set`

Store port object to database, or replace with existing key (that placed in `unlocs` field of object).

```batch
curl -d "{\"name\":\"Dubai\",\"city\":\"Dubai\",\"country\":\"United Arab Emirates\",\"coordinates\":[55.27,25.25],\"province\":\"Dubayy [Dubai]\",\"timezone\":\"Asia/Dubai\",\"unlocs\":[\"AEDXB\"],\"code\":\"52005\"}" -X POST localhost:8008/api/port/set

{"value":"AEDXB"}
```

### Get port object by key `/api/port/get`

Returns port object with given associated key.

```batch
curl -d "{\"value\":\"AEDXB\"}" -X POST localhost:8008/api/port/get

{"name":"Dubai","city":"Dubai","country":"United Arab Emirates","coordinates":[55.27,25.25],"province":"Dubayy [Dubai]","timezone":"Asia/Dubai","unlocs":["AEDXB"],"code":"52005"}
```

### Get port object by name `/api/port/name`

Returns port object with given name. It's looking for port with strict name match.

```batch
curl -d "{\"value\":\"Dubai\"}" -X POST localhost:8008/api/port/name

{"name":"Dubai","city":"Dubai","country":"United Arab Emirates","coordinates":[55.27,25.25],"province":"Dubayy [Dubai]","timezone":"Asia/Dubai","unlocs":["AEDXB"],"code":"52005"}
```

### Find nearest port `/api/port/near`

Finds nearest Port to given coordinates. Recieves `Point` with searching latitude and longitude and returns port with nearest coodinates to given point. Be considered that at port coordinates first value is longitude, second value is latitude.

```batch
curl -d "{\"latitude\":25.873280,\"longitude\":55.011377}" -X POST localhost:8008/api/port/near

{"name":"Umm al Qaiwain","city":"Umm al Qaiwain","country":"United Arab Emirates","coordinates":[55.55,25.57],"province":"Umm Al Quwain","timezone":"Asia/Dubai","unlocs":["AEQIW"]}
```

### Find ports in circle `/api/port/circle`

Finds all ports in given circle. Circle determined by latitude/longitude point of center, and radius in meters.

```batch
curl -d "{\"center\":{\"latitude\":25.458155,\"longitude\":55.148621},\"radius\":40000}" -X POST localhost:8008/api/port/circle

{"list":[{"name":"Sharjah","city":"Sharjah","country":"United Arab Emirates","coordinates":[55.38,25.35],"province":"Ash Shariqah [Sharjah]","timezone":"Asia/Dubai","unlocs":["AESHJ"],"code":"52070"},{"name":"Dubai","city":"Dubai","country":"United Arab Emirates","coordinates":[55.27,25.25],"province":"Dubayy [Dubai]","timezone":"Asia/Dubai","unlocs":["AEDXB"],"code":"52005"},{"name":"Ajman","city":"Ajman","country":"United Arab Emirates","coordinates":[55.513645,25.405216],"province":"Ajman","timezone":"Asia/Dubai","unlocs":["AEAJM"],"code":"52000"},{"name":"Port Rashid","city":"Port Rashid","country":"United Arab Emirates","coordinates":[55.27565,25.284756],"province":"Dubai","timezone":"Asia/Dubai","unlocs":["AEPRA"],"code":"52005"}]}
```

### Find ports with text `/api/port/text`

Finds all ports each of which contains given text in one of the fields: name, city, province, country. Field `sensitive` of argument makes search case sensitive; `whole` matches entire string. Returns list of founded ports if it has.

```batch
curl -d "{\"value\":\"dubai\",\"whole\":true}" -X POST localhost:8008/api/port/text

{"list":[{"name":"Dubai","city":"Dubai","country":"United Arab Emirates","coordinates":[55.27,25.25],"province":"Dubayy [Dubai]","timezone":"Asia/Dubai","unlocs":["AEDXB"],"code":"52005"},{"name":"Port Rashid","city":"Port Rashid","country":"United Arab Emirates","coordinates":[55.27565,25.284756],"province":"Dubai","timezone":"Asia/Dubai","unlocs":["AEPRA"],"code":"52005"},{"name":"Jebel Ali","city":"Jebel Ali","country":"United Arab Emirates","coordinates":[55.02729,24.985714],"province":"Dubai","timezone":"Asia/Dubai","unlocs":["AEJEA"],"code":"52051"}]}
```

```batch
curl -d "{\"value\":\"miam\",\"whole\":false}" -X POST localhost:8008/api/port/text

{"list":[{"name":"Miami","city":"Miami","country":"United States","coordinates":[-80.19179,25.76168],"province":"Florida","timezone":"America/New_York","unlocs":["USMIA"],"code":"5201"}]}
```

---
(c) schwarzlichtbezirk, 2021.
