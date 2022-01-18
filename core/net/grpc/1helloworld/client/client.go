package main

import (
	"fmt"
	"go-learn/core/net/grpc/proto/hello"
	"net/rpc"
)

/* rpc 客户端 */

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(fmt.Errorf("dialing err: %w", err))
	}

	var reply = &hello.String{}
	var args = &hello.String{
		Value: "hello ming",
	}

	err = client.Call("HelloService.Hello", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println("success rpc result:", reply)
}
