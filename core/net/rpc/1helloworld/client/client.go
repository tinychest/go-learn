package main

import (
	"fmt"
	"go-learn/special/proto/hello"
	"log"
	"net/rpc"
)

/* rpc 客户端 */

func main() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing err:", err)
	}

	var reply = &hello.String{}
	var param = &hello.String{
		Value: "hello ming",
	}

	err = client.Call("HelloService.Hello", param, &reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success rpc result:", reply)
}
