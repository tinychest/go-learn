package cache

import (
	"go-learn/todo/cache/cache_f"
	"go-learn/todo/cache/cache_i"
	"testing"
)

type Data struct{}

type DataSource struct{}

func (s DataSource) Query() ([]*Data, error) {
	return nil, nil
}

func Query() ([]*Data, error) {
	return nil, nil
}

func TestSlice(t *testing.T) {
	var _ = cache_f.New("slice1", Query)
	// Goland 无提示，运行时：type DataSource of DataSource{} does not match cache_i.SliceCacheable[T] (cannot infer T)
	// var _ = cache_i.New("data1", DataSource{})
	var _ = cache_i.NewSlice[*Data]("slice2", DataSource{})
	// Goland 无提示，运行时，type cache_i.SliceWrapper[*cache_i.Data] of cache_i.SliceWrapper[*cache_i.Data]{…} does not match cache_i.SliceCacheable[T] (cannot infer T)
	// var _ = cache_i.New("data1", cache_i.SliceWrapper[*cache_i.Data]{Func: Query})
	var _ = cache_i.NewSlice[*Data]("slice3", cache_i.SliceWrapper[*Data]{Func: Query})

	// cache_f 和 cache_i 都是泛型函数，但是 cache_f 调用时，并不需要显示指定实际泛型类型，而 cache_i 则必须指定
}

func TestAny(t *testing.T) {
	var c = cache_i.New[[]*Data]("data", DataSource{})
	list := c.Load()
	for _, v := range list {
		t.Log(v)
	}
}
