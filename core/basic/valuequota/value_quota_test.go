package valuequota

import (
    "fmt"
    "go-learn/core"
    "testing"
)

func TestReturn(t *testing.T) {
    p := getPerson()
    // 0xc00003c570
    t.Logf("%p\n", &p)

    // 结论：地址值是不同的，即在函数返回值中也是值传递拷贝的这个道理
}

func getPerson() core.Person {
    p := core.Person{Age: 11, Name: "小明"}

    // 0xc00003c570
    fmt.Printf("%p\n", &p)

    return p
}
