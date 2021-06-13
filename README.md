# Golan

Golan is a content engine, written in Go, for the project-name-redacted `Enterprise Content Management` solution. Golan 
manages the storage & retrieval of arbitrary contents across a wide variety of storage mediums 
(databases, AWS S3, Raw Storage - At least that's the plan).

Golan pluggable's architecture allows for new storage mediums to be added with no to minimal changes in the core functionality.

## Authentication & Authorization
Golan supports OAuth2 authentication and participate in this protocol as a resource server.

## Architecture & Design

## Getting Started


### Developing for DocumentDB / MongoDB

Starting MongoDB
```shell
$ docker compose up -d mongo mongo-express
```

You can access the Web-based mongo client at: [http://localhost:8081/](http://localhost:8081/)

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

Using Evans (REPL)
```shell
$ cd proto
$ evans --host 127.0.0.1 --port 9000 repl --proto service.proto
```
![Evans](docs/evans.png?raw=true "Evans")

Calling gRPC Service Using Evans

![Calling gRPC Service in Evans](docs/calling-service-in-evans.png?raw=true "Evans")

Using Evans (CLI)

If you intend to use `seed.sh` to test the Golan ECM's gRPC endpoints, you need to install
[jq](https://stedolan.github.io/jq/). [jq](https://stedolan.github.io/jq/) is used to parse the output from stdout and use the parsed output as an input on 
subsequent calls.

On MacOS, you can install [jq](https://stedolan.github.io/jq/) using homebrew
```shell
$ brew install jq
Updating Homebrew...
...

$ which jq
/usr/local/bin/jq
```
