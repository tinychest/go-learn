package ip

import (
	"net"
)

/*
IsLoopback：是否是本地回环地址
127.0.0.1、::1

IsPrivate：是否是内网地址（Go 1.7 新增）
10.0.0.0 ~ 10.255.255.255
172.16.0.0 ~ 172.31.255.255
192.168.0.0 ~ 192.168.255.255
*/

func IsLoopback(ipStr string) bool {
	return net.ParseIP(ipStr).IsLoopback()
}

func IsInternalIP(ipStr string) bool {
	ip := net.ParseIP(ipStr)
	return ip.IsPrivate() || ip.IsLoopback()
}
