package http_rpc

import (
	"go-learn/special/performance/benchmark/http_rpc/dto"
	"log"
	"net/rpc"
)

func RPCClient() {
	client, err := rpc.Dial("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing err:", err)
	}

	param := dto.Param{}
	reply := dto.Reply{}

	err = client.Call("hello.Hello", param, &reply)
	if err != nil {
		log.Fatal(err)
	}
	// log.Println("success rpc result:", reply)
	// 不关闭，服务端不会释放连接，导致第二次发起连接阻塞
	err = client.Close()
	if err != nil {
		log.Fatal(err)
	}
}

var conn *rpc.Client

func getClient() *rpc.Client {
	var err error

	if conn == nil {
		conn, err = rpc.Dial("tcp", "localhost:1234")
		if err != nil {
			log.Fatal("dialing err:", err)
		}
	}

	return conn
}

func RPCClientReuse() {
	client := getClient()

	param := dto.Param{}
	reply := dto.Reply{}

	err := client.Call("hello.Hello", param, &reply)
	if err != nil {
		log.Fatal(err)
	}
}
