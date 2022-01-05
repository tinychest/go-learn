package performance

/*
Dead code elimination（下边就简称 DCE） - wikipedia

In compiler theory, dead code elimination
(also known as DCE, dead code removal, dead code stripping, or dead code strip)
is a compiler optimization to remove code which does not affect the program results.
*/

/*
作用：减小程序体积，程序运行过程中避免执行无用的指令，缩短运行时间
1、常量替换变量：这样就能在编译时期，将肯定不会执行到的代码给去掉（如果这部分代码还是 hot path，性能提升会更加明显）
2、Go 的编译器也会堆局部变量进行 DCE
3、包级变量，不会进行 DCE
    包初始化函数 init() 中，init() 函数可能有多个，且可能位于不同的 .go 源文件
    包内的其他函数
    如果是 public 变量（首字母大写），其他包引用时可修改
*/

/*
既然 Go 的编译器有 DCE 的操作，我们就可以利用起来，比如定义一个常量 debug bool，然后程序中各处判断该常量，
进行日志输出之类的，反正修改一下值，最终打包的程序，并不含这些代码

当前，修改源码感觉并不优雅，Go 提供了一种方式，结果就是，在 build 的时候，通过参数来指定

debug.go
```
/// +build debug
package main
const debug = true
```

release.go
```
/// +build !debug
package main
const debug = false
```

`// +build debug`  表示 build tags 中包含 debug 时，该源文件参与编译
`// +build !debug` 表示 build tags 中不包含 debug 时，该源文件参与编译

编译 debug   版本：go build -tags debug -o debug .
编译 release 版本：go build -o release .
（-o 参数是指定输出的目录、最后的点表示编译的源）
*/
