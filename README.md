# Butcher

![LOGO](logo.jpeg)

## Description

A set of Dubbo utils provided by Pudge meant to be helpful.

Butcher is a simple wrapper around dubbo-telnet with go, aimed to made debuging
dubbo more esayly.

## Usage

```shell
##
# get help and show avaliable commands
$ buthcer -h

A set of Dubbo utils provided by Pudge meant to be helpful.

Usage:
  Butcher [command]

Available Commands:
  help        Help about any command
  invoke      Invoke given command on dubbo instance.
  ls          Show all provider and consumers

Flags:
  -h, --help          help for Butcher
  -H, --host string   Dubbo connection host (default "0.0.0.0")
  -P, --port int      Dubbo connection port (default 20880)

Use "Butcher [command] --help" for more information about a command.

##
# list dubbo container services and methods
$ buthcer -H 10.0.11.1 -P 20880 ls

PROVIDER:
com.example.Provier

CONSUMER:
com.example.Consumer

##
# Invoke the service method on dubbo container
$ buthcer -H 10.0.11.1 -P 20880 invoke 'com.example.Provier.hello("world")'

Use default com.example.Provier.
result: {"response": "world"}
elapsed: 1 ms.

##
# Batch invoke methods from a given file.
# invoke-list.txt:
#   com.example.Provier.hello("world a")
#   com.example.Provier.hello("world b")
$ buthcer -H 10.0.11.1 -P 20880 invoke --file invoke-list.txt

Use default com.example.Provier.
result: {"response": "world a"}
elapsed: 1 ms.

Use default com.example.Provier.
result: {"response": "world b"}
elapsed: 1 ms.
```

### Installation

1. From source code.
   1. Install gox by `go get github.com/mitchellh/gox`
   2. Build by `make all`
2. Download from pre-build releases.

### Known Issues

Butcher building around Dubbo telnet protocol, synced sending commands and async reading response, there's no guarantee of invocation order and may cause issues when invoke frequency is too high. Also, you need to specify "--sleep" flag for "invoke" command if RPC will cost a lot of time, the default sleep time between each invocation is 500 milliseconds.

### Contribution

1. Fork the repository
2. Create Feat_xxx branch
3. Commit your code
4. Create Pull Request
