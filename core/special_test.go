package core

import "testing"

type Master interface {
	GetId() int64
	// 这个参数类型不能是 *[]Slave 因为借助 map，而 map 的值是不可寻址的
	SetSlaves([]Slave)
}

type Slave interface {
	GetMasterId() int64
}

// 将 slaves 分组塞给 masters
func Organize(masters []Master, slaves []Slave) {
	orgMap := make(map[int64][]Slave, len(masters))

	for _, slave := range slaves {
		if orgMap[slave.GetMasterId()] == nil {
			orgMap[slave.GetMasterId()] = make([]Slave, 0)
		}
		orgMap[slave.GetMasterId()] = append(orgMap[slave.GetMasterId()], slave)
	}
	for _, master := range masters {
		master.SetSlaves(orgMap[master.GetId()])
	}
}

// -------------分割线

type Fruit struct {
	Id     int64
	Apples []*Apple
}

type Apple struct {
	Id      int64
	FruitId int64
}

func (f *Fruit) GetId() int64 {
	return f.Id
}

func (f *Fruit) SetSlaves(apples []*Apple) {
	f.Apples = apples
}

func (a *Apple) GetMasterId() int64 {
	return a.FruitId
}

// --------------分割线
func TestSpecial(t *testing.T) {
	fruits := []*Fruit{{Id: 1}, {Id: 2}}
	apples := []*Apple{{Id: 1, FruitId: 1}, {Id: 2, FruitId: 1}, {Id: 3, FruitId: 2}, {Id: 4, FruitId: 2}, {Id: 5, FruitId: 2}}

	// *Apple 类型可以转成 Slave
	_ = Slave(&Apple{})
	// []*Apple 类型不能转成 []Slave
	_ = []Slave([]*Apple{})

	// 所以下边的方法无法调用
	Organize(fruits, apples)
}
