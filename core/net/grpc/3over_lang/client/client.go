package main

import (
	"fmt"
	over "go-learn/core/net/grpc/3over_lang"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(fmt.Errorf("net.Dial: %w", err))
	}

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

	var reply string
	err = client.Call(over.HelloServiceName+".Hello", "hello", &reply)
	if err != nil {
		panic(err)
	}

	fmt.Println(reply)
}
