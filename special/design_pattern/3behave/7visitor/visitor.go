package _visitor

import "fmt"


/*
给的样例，只是说明模式，没有和实际场景关联上（可以以不同文件的压缩策略为背景）

并且由于采用了类型断言，所以如果需要操作的对象比较多的话，这个函数其实也会膨胀的比较厉害
后续可以考虑按照命名约定使用 generate 自动生成代码 或者是使用 反射 简化
*/

// Visitor 访问器
type Visitor interface {
	Visit(IResource) error
}

// IResource 资源（一组对象的抽象）
type IResource interface {
	Accept(Visitor) error
}

// 资源实现

type resource1 struct{
	Key string
}

type resource2 struct {
	Value string
}

func (r *resource1) Accept(v Visitor) error {
	return v.Visit(r)
}

func (r *resource2) Accept(v Visitor) error {
	return v.Visit(r)
}

// 访问器实现

type Visitor1 struct {}

func (v *Visitor1) Visit(i IResource) error {
	switch t := i.(type) {
	case *resource1:
		return v.VisitResource1(t)
	case *resource2:
		return v.VisitResource2(t)
	default:
		return fmt.Errorf("not found resource typr: %#v", i)
	}
}

func (v Visitor1) VisitResource1(i *resource1) error {
	fmt.Println(i.Key)
	return nil
}

func (v Visitor1) VisitResource2(i *resource2) error {
	fmt.Println(i.Value)
	return nil
}