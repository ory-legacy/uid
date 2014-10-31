baldur.io - uid generator
===

This go module aims to provide unique identifiers (64 bit integers) in a concurrent high performance environment,
taking into account the id's type, the timestamp, the node which created this id and a unique offset.

This work is based upon: https://github.com/Metalcon/muid

## Status

[![Build Status](https://travis-ci.org/baldur-io/uid.svg)](https://travis-ci.org/baldur-io/uid)

## Usage

```
import "github.com/baldur-io/uid"

var idType int64 = 123
var timestamp int64 = 124124
var Node int64 = 321
var offset int64 = 313
var err error
var id uid.Uid

id, err = New(idType, Node, timestamp, offset)

if err != nil {
    // Something happened
}

fmt.Println(id.Type()) // 123
```

## RPC

This module has an RPC wrapper for running an independent uid generator service: `src/service.go`.
See `examples/server.go` and `examples/client.go` for example implementations.

## TBDs

* Add more test cases
* Add test cases for the service wrapper
* Test RPC server/client
* Improve string serialisation
* Use uint32 / byte / uint16 instead of int64 at: type, node, offset