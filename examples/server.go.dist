package main

import (
	"net/rpc"
	"net"
	"net/http"
	"log"
	"github.com/baldur-io/uid"
)

func main() {
	uidService := new(uid.Service)
	rpc.Register(uidService)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":4001")
	if e != nil {
		log.Fatal("listen error: ", e)
	}
	http.Serve(l, nil)
}
