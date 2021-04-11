package mock

import (
	"errors"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestGetFromDB(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish() // 断言 DB.Get() 方法是否被调用

	m := NewMockDB(ctrl)
	// 定义测试预期
	m.EXPECT().Get(gomock.Eq("Tom")).Return(-1, errors.New("not exist"))
	// m.EXPECT().Get(gomock.Any()).Return(-1, nil)
	// m.EXPECT().Get(gomock.Not("Sam")).Return(200, nil)
	// m.EXPECT().Get(gomock.Nil()).Return(-1, errors.New("nil"))

	// 被调用的顺序、次数都是可以指定的

	// 实际测试（传入 tom，在 GetFromDB 方法的内部会得到上面定义的预期，即 -1 和 error：not exist）
	if v := GetFromDB(m, "Tom"); v != -1 {
		t.Fatal("expected -1, but got", v)
	}
}
