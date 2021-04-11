package _proxy

// 当前文件是模型定义，也就是被代理类（这里，被代理类连接口方法都没有实现）

// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
// @proxy IUser
type User struct {
}
