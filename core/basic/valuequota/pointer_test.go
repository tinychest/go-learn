package valuequota

import (
	"fmt"
	"testing"
)

type Person struct {
	age    int
	name   string
	school *School
}

type School struct {
	name    string
	address string
}

func TestPointer(t *testing.T) {
	var school = School{name: "清华大学", address: "山顶茅厕"}
	var person = Person{age: 1, name: "小明", school: &school}

	person.print("main方法中")

	// person.fakeChangeName("小红1")
	// person.print("假改名后：") // 没有改
	// person.realChangeName("小红2")
	// person.print("真改名后：") // 改了

	/* 第一轮测试：Person 中的 School不是 * 类型 */
	// person.fakeChangeSchoolName()
	// person.print("假改学校名后：") // 没有改
	// person.realChangeSchoolName()
	// person.print("真改学校名后：") // 改了

	// person.fakeChangeSchool()
	// person.print("假改学校后：") // 没有改
	// person.realChangeSchool()
	// person.print("真改学校后：") // 改了

	/* 第二轮测试：Person 中的 School是 * 类型 */
	// person.fakeChangeSchoolName()
	// person.print("假改学校名后：") // 改了
	// person.realChangeSchoolName()
	// person.print("真改学校名后：") // 改了

	// person.fakeChangeSchool()
	// person.print("假改学校后：") // 没有改
	// person.realChangeSchool()
	// person.print("真改学校后：") // 改了

	// 结论（早就熟知的）：抓住值传递（副本）和指针类型两个关键就不会有问题
	// 方法传参，如果参数（参数如果是结构体，结构体属性有指针也同理 - 嵌套也适用）是指针类型，就能改指针指向的内存中的数据
	// 方法传参，如果参数不是指针类型，那么就是值传递，实际方法的形参对应的变量只是一个副本

	/* 第三轮测试 */
	(&person).fakeChangeName("小红")
	person.print("模拟指针类型的方法调用：")
}

func (p Person) fakeChangeSchool() {
	p.school = &School{name: "XX大学", address: "oh no"}
}

func (p *Person) realChangeSchool() {
	p.school = &School{name: "XX大学", address: "oh no"}
}

func (p Person) fakeChangeSchoolName() {
	p.school.name = "XX大学"
	p.school.address = "oh no"
}

func (p *Person) realChangeSchoolName() {
	p.school.name = "XX大学"
	p.school.address = "oh no"
}

// 基础的数据类型，是会复制一份的
func (p Person) fakeChangeName(name string) {
	p.print("假改名中")
	p.name = name
}

func (p *Person) realChangeName(name string) {
	// 下面这样打印的是指针类型的内存地址值
	// fmt.Printf("realChangeName 收到的变量的内存地址值: %p\n", &p)
	p.print("真改名中")
	p.name = name
}

// 要打印地址，所以必须加 *
func (p *Person) print(prefix string) {
	fmt.Printf("%s: 内存地址：%p, {age: %d, name:%s, school: 内存地址:%p, {name: %s, address: %s}}\n", prefix, p, p.age, p.name, &p.school, p.school.name, p.school.address)
}
