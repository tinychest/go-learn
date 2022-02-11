package example2

/*
极简的数据实体，接口方法定义

命令 protoc --go_out=. hello.proto
	生成的东西不包含接口定义、客户端、服务端 代码实现
命令 protoc --go_out=plugins=grpc:. hello.proto
	生成的东西则包含

生成的接口定义、客户端、服务端时间看下来的感受是：
接口没什么好说，就是你自己定义的
客户端具体实现，麻烦一些
服务端具体实现，就是指定接口的实现类，简单一些
*/
