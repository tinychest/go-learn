package ip

import (
	"fmt"
	"net"
	"testing"
)

// https://www.jianshu.com/p/3c8a4cce9cd1
// https://www.cnblogs.com/psy-sdudio/p/11530592.html#:~:text=localhost,%E5%AE%83%E5%B9%B6%E4%B8%8D%E6%98%AFIP%EF%BC%8C%E8%80%8C%E6%98%AF%E4%B8%80%E7%A7%8D%E7%89%B9%E6%AE%8A%E7%9A%84%E5%9F%9F%E5%90%8D%EF%BC%88%E6%B2%A1%E6%9C%89%E5%90%8E%E7%BC%80%EF%BC%89%EF%BC%8C%E9%BB%98%E8%AE%A4%E7%9A%84%E6%83%85%E5%86%B5%E4%B8%8B%E5%AE%83%E8%A7%A3%E6%9E%90%E5%88%B0%E7%9A%84%E6%98%AF%E6%9C%AC%E5%9C%B0IP%EF%BC%88127.0.0.1%EF%BC%89%EF%BC%8C%E4%B8%BB%E8%A6%81%E9%80%9A%E8%BF%87%E6%9C%AC%E6%9C%BA%E7%9A%84host%E6%96%87%E4%BB%B6%E8%BF%9B%E8%A1%8C%E7%AE%A1%E7%90%86%EF%BC%8C%E5%A6%82%E6%9E%9C%E4%BD%A0%E6%84%BF%E6%84%8F%EF%BC%8C%E4%B9%9F%E5%8F%AF%E4%BB%A5%E6%8A%8Alocalhost%E5%9F%9F%E5%90%8D%E8%A7%A3%E6%9E%90%E5%88%B0%E6%9F%90%E4%B8%AA%E5%85%AC%E7%BD%91IP%E4%B8%8A%E5%8E%BB%E3%80%82%20127.0.0.1%E6%98%AF%E4%B8%80%E7%A7%8D%E6%9C%AC%E6%9C%BA%E4%BF%9D%E7%95%99%E7%9A%84%E7%A7%81%E6%9C%89IP
func TestParseIP(t *testing.T) {
	ipLocalV4 := "127.0.0.1"
	// ipLocalV6 := "::1"
	ipLocalV6 := "0:0:0:0:0:0:0:1"

	ipv4 := net.ParseIP(ipLocalV4)
	fmt.Println(ipv4)
	fmt.Println(ipv4.To16())

	ipv6 := net.ParseIP(ipLocalV6)
	fmt.Println(ipv6)
	fmt.Println(ipv6.To4())
}