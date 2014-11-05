package main

import (
	"log"
	"net/rpc"
	"fmt"
	"github.com/baldur-io/uid"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:4001")

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &uid.CreatorArguments{7}

	var result uid.Uid

	err = client.Call("Service.New", args, &result)
	if err != nil {
		log.Fatal("Call error:", err)
	}

	fmt.Printf("Identifier %d (%s) at adress %X created \n", result, result.String(), &result)
}
