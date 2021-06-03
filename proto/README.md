# Golan ECM Protocol Buffer Definitions
This section describes the request and response interface structures exposed
by this service.

## Getting Started

Install Pre-requisites
```shell
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

~/.bash_profile
```
...
# This could be different on your local machine
export GOPATH=/usr/local/Cellar/go/1.16.3

export PATH="$PATH:$GOPATH/bin"
...

```

If you are using zsh

~/.zshrc
```
...
if [ -f ~/.bash_profile ]; then
    . ~/.bash_profile;
fi
```

## Building

```shell
$ protoc -I=proto --go_out=internal proto/reqres.proto && \
protoc -I=proto --go-grpc_out=internal proto/service.proto
```
