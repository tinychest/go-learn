package todo

// 在 Go 中，基础类型可以比较；结构体，如果都是基础类型，那么结构体实例可以比较
// - 可以通过在 结构体 中添加一个 _ [0]func() 来达到，无法比较的效果

// 可以通过在 结构体 中添加一个 _ struct{} 来避免结构体的纯值实例化方式

// Go 不支持切片类型进行比较（参照自动生成的 func test，可以看到深比较 reflect.DeepEqual
//    https://golang.org/ref/spec#Comparison_operators