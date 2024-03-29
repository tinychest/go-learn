package _embed

// go 1.16 的新特性

// 假如应用程序会连带一个些静态文件资源，你总是要不可避免的保证这些文件要匹配二进制程序的相对目录
// 但是，假如这些静态文件不会更改，且更希望隐藏起来，那就可以考虑使用 go embed

// 这个思路是由某些三方类库先提出来（go-bindata），到了 1.16 go 才开始自身支持了