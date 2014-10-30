package main

import (
	"log"
	"net/rpc"
	"fmt"
	"uid"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:4001")
	count := 2

	if err != nil {
		log.Fatal("dialing:", err)
	}

	args := &uid.CreatorArguments{7}
	res := make(chan *rpc.Call, count)

	//var result uid.Uid
	//client.Call("Service.Create", args, &result)
	//fmt.Printf("Identifier: %d \n", result)

	for i := 0; i < count; i++ {
		var result uid.Uid
		client.Go("Service.Create", args, &result, res)
	}

	for call := range res {
		var result *uid.Uid
		if call.Error != nil {
			log.Fatal("Server error: ", err)
		}
		result = call.Reply.(*uid.Uid)
		fmt.Printf("Identifier %s %d created \n", result, result)
	}
}
