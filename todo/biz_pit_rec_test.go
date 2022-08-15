package todo

import (
	"fmt"
	"testing"
	"time"
)

/* 实际开发中遇到的打印信息的模棱两可 */
func TestPrintStringSlice(t *testing.T) {
	t.Log([]string{"a", "b"})   // [a b]
	t.Log([]string{`"a b"`})    // ["a b"]
	t.Log([]string{`"a`, `b"`}) // ["a b"]

	t.Log(fmt.Sprintf("%+v", []string{"1", "2", "3"}))
	t.Log(fmt.Sprintf("%#v", []string{"1", "2", "3"}))
}

/* 实际开发中遇到不少的问题；很容易会被忽视，被 go 关键字迷糊的问题 */
func TestGo(t *testing.T) {
	ch1 := make(chan int)
	go fmt.Println(<-ch1) // 确认参数值时会发生什么呢？应该通过闭包函数去达到目标效果
	ch1 <- 5
	time.Sleep(1 * time.Second)
}

/*
  反引号
  Go string 的 backtick 对于 Go 编译器来说就是不进行任何转移的语义。
  所以不支持 backtick 内嵌 backtick 的写法，这点在正则表达式的案例中，就有启示过。
  实际开发中会遇到相关的需求，通常写一段多行的 SQL 语句，肯定会用反引号来表示这个字符串，
  但是 MySQL 的 SQL 语法，会有将特定关键字使用反引号引起来，来取消其表达含义的场景。

  结论：数据库、数据表 设计，尽量规范一些，不要使用关键字作为相关表名、字段名
*/
func TestPrintBacktick(t *testing.T) {
	q := `\x60`
	println(q)
}
