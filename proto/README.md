# Golan ECM Protocol Buffer Definitions
This section describes the request and response interface structures exposed
by this service.

## Getting Started

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
$ protoc -I=proto --go_out=./pb proto/service.proto proto/reqres.proto
```
