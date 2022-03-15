package unit

// 概念：go test 命令可以通过添加一个参数，让其在执行测试用例代码的时候，为执行到的代码添加标记，收集好标记信息后，输出到文件中
// 再通过 go tool cover 来分析这个文件信息，通过各种方式将覆盖率信息展示出来
//
// [第一步] 执行一次单元测试，并将相关信息写到指定的文件中
// 语法：go test -coverprofile <filename> <package name>
// 实例：go test -coverprofile xxx.out
// （得到当前包下所有测试用例的覆盖率信息）
// （out 只是一种习惯，命名成什么都行）
//
// 如果想得到当前目录下某个包下的所有测试用例覆盖率信息，就要指定 <package name> 参数了
// 实例：go test -coverprofile xxx.out ./xxx
// （不会递归扫描子子包）
//
// [第二步]
// 根据测试用例信息，分析得出测试覆盖率（粒度到方法的每行代码）
// 帮助：go tool cover -help
// 将覆盖率信息以非常形象的方式，用页面展示出来
// 语法：go tool cover -html <xxx.out>
// 实例：go tool cover -html="cover.out" 或 go tool cover -html cover.out
// （这里只是皮毛，实际上还支持很多参数，如：-func -mode -o -var）
