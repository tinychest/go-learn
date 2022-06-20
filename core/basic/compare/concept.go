package compare

// [官方文档] https://golang.org/ref/spec#Comparison_operators
//
// bool、int、float、complex、string、pointer、chan、interface、array 都是可比较的
//
// [new]
// Go 1.18 引出的新概念见：core/basic/generic/icomparable_test.go
//
// [number]
// 浮点数相关的比较见：core/basic/number
//
// [slice]
// 不支持比较
//     go test 是通过反射包的 reflect.DeepEqual 方法来对实际结果和预期进行比较的
//
// [struct]
// 在 Go 中，基础类型可以比较；结构体，如果都是基础类型，那么结构体实例可以比较
//    可以通过在 结构体 中添加一个 _ [0]func() 来达到，无法比较的效果
