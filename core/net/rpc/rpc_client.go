package rpc

import (
	"fmt"
	"go-learn/core/net/rpc/dto"
	"net/rpc"
)

func RPCClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		panic(fmt.Errorf("dialing err: %w", err))
	}

	args := dto.Args{}
	reply := dto.Reply{}

	err = client.Call("hello.Hello", args, &reply)
	if err != nil {
		panic(err)
	}
	fmt.Println(reply)

	// fmt.Println("success rpc result:", reply)
	// 不关闭，服务端不会释放连接，导致第二次发起连接阻塞
	err = client.Close()
	if err != nil {
		panic(err)
	}
}

var conn *rpc.Client

func getClient() *rpc.Client {
	var err error

	if conn == nil {
		conn, err = rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			panic(fmt.Errorf("dialing err: %w", err))
		}
	}

	return conn
}

func RPCClientReuse() {
	client := getClient()

	args := dto.Args{}
	reply := dto.Reply{}

	err := client.Call("hello.Hello", args, &reply)
	if err != nil {
		panic(err)
	}
}
