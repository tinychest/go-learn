package test

// 测试方法签名要求：func TestXxx(t *testing.T) {}
//     Test 和 Xxx 之间还可以用若干个 _ 隔开
//
// 测试文件名要求：xxx_test.go（以 _test.go 结尾）
//
// 命令语法：（忘记了怎么办，直接在 Golang UI 中运行一个测试方法，然后查看控制台最上面的详细命令）
// go test [-v] [-run[=]<regex>]
// -run <regex>：通过正则匹配要执行的方法名，从而决定要执行的方法
//     例1：go test -run ^\QTestXxx\E$：执行当前目录下所有 test.go 中所有名为 TestXxx 的测试方法
//     例2：go test -run Xxx：执行当前目录下所有 test.go 中方法名包含 Xxx 的方测试方法
//     例3：go test -run TestXxx：执行当前目录下所有 test.go 中方法名以 TestXxx 开头的测试方法
//     例4：go test xxx1_text.go xxx2_test.go：指定指定测试文件中的所有测试方法
// -v：是否在 console 中打印 t.Log 或 t.Logf 中的内容
//
// 在 Goland 中，可以通过 Ctrl + Shift + T 来快速生成指定 方法 或者 文件 的测试用例
// 生成的测试方法的结构其实是表格形式，表格形式能够更加清晰的将多组测试数据展现出来的方式
