package _step

// 模拟服务发现，这里是静态写死的

func ServerIP() string {
	return "127.0.0.1"
}

func ServerPort() int {
	return 1234
}

func ServerPortStr() string {
	return "1234"
}

func ServerAddr() string {
	return "127.0.0.1:1234"
}

func ServiceName() string {
	return "HelloService"
}

func ServiceFuncName() string {
	// 注意首字母大小写
	return "Hello"
}