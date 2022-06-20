package _package

// 包名为 internal 的包具有特殊含义，会阻断来自包外的所有访问
// internal 当前所在的包下的 子文件 和 子包下的文件（internal 包同级、internal 包的子文件），可以正常访问 internal 包下的所有内容
// https://golang.google.cn/doc/go1.4#internalpackages
