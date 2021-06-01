# Golan

Golan is a content engine, written in Go, for the project-name-redacted `Enterprise Content Management` solution. Golan 
manages the storage & retrieval of arbitrary contents across a wide variety of storage mediums 
(databases, AWS S3, Raw Storage - At least that's the plan).

Golan pluggable's architecture allows for new storage mediums to be added with no to minimal changes in the core functionality.

## Authentication & Authorization
Golan supports OAuth2 authentication and participate in this protocol as a resource server.

## Architecture & Design

## Getting Started

## Testing gRPC APIs

See [https://github.com/ktr0731/evans](https://github.com/ktr0731/evans)

If you're using MacOS
```shell
$ brew tap ktr0731/evans
$ brew install evans
```

Build & Start Golan ECM `golan/cmd/main.go`. By default, ecm will run on port 9000
```shell
$ go build -o ecm-exec cmd/main.go
$ ./ecm-exec --config samples/config.yaml
```

Start Evans Cli
```shell
$ cd proto
$ evans --host 127.0.0.1 --port 9000 repl --proto service.proto
```
![Evans](docs/evans.png?raw=true "Evans")

Calling Service Using Evans

![Calling gRPC Service in Evans](docs/calling-service-in-evans.png?raw=true "Evans")