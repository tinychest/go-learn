package case_delay

import (
	"fmt"
	"go-learn/core"
	"testing"
)

var ps = []core.Person{
	{Name: "小明"}, {Name: "小红"}, {Name: "小光"},
}

func TestDelayFunc(*testing.T) {
	mistake()
}

func mistake() {
	var pfs []func()
	for _, item := range ps {
		pfs = append(pfs, func() {
			fmt.Println(item.Name)
		})
	}

	for _, pf := range pfs {
		pf()
	}
}

// 解决方法完全可以参照 case1_defer，同时也可以反向衬托出这里，使得更好理解