# Butcher

![LOGO](logo.jpeg)

## Description

Helpfull hooks provided by Pudge, sets of dubbo utils.

Butcher is a simple wrapper around dubbo-telnet with go, aimed to made debuging
dubbo more esayly.

## Usage

```shell
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

### Contribution

1. Fork the repository
2. Create Feat_xxx branch
3. Commit your code
4. Create Pull Request
