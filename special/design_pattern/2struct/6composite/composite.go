package _composite

// IOrganization 组织抽象接口
type IOrganization interface {
	Count() int
}

type Organization struct {
	Name     string
	Children []IOrganization
}

type Employee struct {
	Name string
}

func (o *Organization) Count() int {
	var c int
	for _, v := range o.Children {
		c += v.Count()
	}
	return c
}

func (e *Employee) Count() int {
	return 1
}
