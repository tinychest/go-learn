package _struct

import (
	"sort"
	"testing"
)

// 从 core/basic/typ/unsafe_type_cast_test.go 中了解到
// 我们无法规避手动将一个自定义结构体切片转成指定接口切片类型的过程，这是一个方面，先搁置一下
//
// 该如何复用，以合理、优雅的方式来复用：
// 既然，限定了方法的参数类型，那就注定不同的实现细节会体现在不同的接口体中，即，包装模式 或者 装饰器模式
//
// 这里定义了一个接口 和 一个实现方法，用于实现为一组数据进行排名，相同的分数排名相同
// （这里特指实际业务中比较特殊且复杂的逻辑，希望做到复用）

type IntComparable interface {
	Getter() int
	Setter(ranking int)
}

func Ranking(list []IntComparable) {
	sort.Slice(list, func(i, j int) bool {
		return list[i].Getter() > list[j].Getter()
	})

	var (
		gearIdx = 0
		pre     = -1
		cur     int
	)
	for i, v := range list {
		cur = v.Getter()
		if pre != cur {
			gearIdx = i + 1
			pre = cur
		}
		v.Setter(gearIdx)
	}
}

// 使用场景，得出学生数学科目的成绩排名

type Student struct {
	Name        string
	MathScore   int
	MathRanking int
}

func (s Student) Getter() int {
	return s.MathScore
}

func (s *Student) Setter(ranking int) {
	s.MathRanking = ranking
}

func TestMathRanking(t *testing.T) {
	students := []*Student{
		{Name: "小聪", MathScore: 100},
		{Name: "小红", MathScore: 99},
		{Name: "小明", MathScore: 100},
		{Name: "小光", MathScore: 60},
	}

	cs := make([]IntComparable, 0, len(students))
	for _, v := range students {
		cs = append(cs, v)
	}

	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(cs)
	t.Log(*students[0], *students[1], *students[2], *students[3])
}

// 目前为止，没有任何问题；但是，假如学生还有 语文成绩、英语成绩呢？
// 实际上，不需要过多思考，基本上会得到这样的答案

type Student2 struct {
	Name           string
	MathScore      int
	MathRanking    int
	EnglishScore   int
	EnglishRanking int
}

type MathWrap struct {
	*Student2
}

func (w MathWrap) Getter() int {
	return w.MathScore
}

func (w MathWrap) Setter(ranking int) {
	w.MathRanking = ranking
}

type EnglishWrap struct {
	*Student2
}

func (w EnglishWrap) Getter() int {
	return w.EnglishScore
}

func (w EnglishWrap) Setter(ranking int) {
	w.EnglishRanking = ranking
}

func TestMultiRanking2(t *testing.T) {
	students := []*Student2{
		{Name: "小聪", MathScore: 100, EnglishScore: 100},
		{Name: "小红", MathScore: 99, EnglishScore: 59},
		{Name: "小明", MathScore: 100, EnglishScore: 89},
		{Name: "小光", MathScore: 60, EnglishScore: 79},
	}

	// 数学
	ews := make([]IntComparable, len(students))
	for i, v := range students {
		ews[i] = MathWrap{v}
	}
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])

	// 英语
	for i, v := range students {
		ews[i] = EnglishWrap{v}
	}
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])
}

// 接口内嵌

type Student3 struct {
	IntComparable
	Name           string
	MathScore      int
	MathRanking    int
	EnglishScore   int
	EnglishRanking int
}

func (s *Student3) MathComparable() IntComparable {
	return &mathComparable{s}
}

func (s *Student3) EnglishComparable() IntComparable {
	return &englishComparable{s}
}

type mathComparable struct {
	*Student3
}

func (c mathComparable) Getter() int {
	return c.MathScore
}

func (c *mathComparable) Setter(ranking int) {
	c.MathRanking = ranking
}

type englishComparable struct {
	*Student3
}

func (c englishComparable) Getter() int {
	return c.EnglishScore
}

func (c *englishComparable) Setter(ranking int) {
	c.EnglishRanking = ranking
}

func TestMultiRanking3(t *testing.T) {
	students := []*Student3{
		{Name: "小聪", MathScore: 100, EnglishScore: 100},
		{Name: "小红", MathScore: 99, EnglishScore: 59},
		{Name: "小明", MathScore: 100, EnglishScore: 89},
		{Name: "小光", MathScore: 60, EnglishScore: 79},
	}

	// 数学
	ews := make([]IntComparable, len(students))
	for i, v := range students {
		ews[i] = v
		v.IntComparable = v.MathComparable()
	}
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])

	// 英语
	for _, v := range students {
		v.IntComparable = v.EnglishComparable()
	}
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])
}

// 最后，再给出一种拓展思路，程序代码从来都是千变万化

type Mode string

const (
	Math    Mode = "Math"
	English Mode = "English"
)

var mode Mode

type Student4 struct {
	Name           string
	MathScore      int
	MathRanking    int
	EnglishScore   int
	EnglishRanking int
}

func (s Student4) Getter() int {
	if mode == Math {
		return s.MathScore
	}
	if mode == English {
		return s.EnglishScore
	}
	return -1
}

func (s *Student4) Setter(ranking int) {
	if mode == Math {
		s.MathRanking = ranking
	}
	if mode == English {
		s.EnglishRanking = ranking
	}
}

func TestMultiRanking4(t *testing.T) {
	students := []*Student4{
		{Name: "小聪", MathScore: 100, EnglishScore: 100},
		{Name: "小红", MathScore: 99, EnglishScore: 59},
		{Name: "小明", MathScore: 100, EnglishScore: 89},
		{Name: "小光", MathScore: 60, EnglishScore: 79},
	}

	// 数学
	ews := make([]IntComparable, len(students))
	for i, v := range students {
		ews[i] = v
	}
	mode = Math
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])

	// 英语
	mode = English
	t.Log(*students[0], *students[1], *students[2], *students[3])
	Ranking(ews)
	t.Log(*students[0], *students[1], *students[2], *students[3])
}
