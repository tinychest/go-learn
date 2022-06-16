package distribute

import (
	"fmt"
	"testing"
	"time"
)

func TestSnowFlake_NextID(t *testing.T) {
	// milliSumTest(t) // OK 符合预期，每个毫秒的时间里，最多生成 4096 个
	dupTest(t)
}

// 测试一段时间内生成的 ID 数量的统计结果是否符合预期
func milliSumTest(t *testing.T) {
	sf := NewSnowFlake()
	m := make(map[int64]int)

	go func() {
		for {
			_, err := sf.NextID()
			if err != nil {
				panic(err)
			}
			m[time.Now().UnixMilli()]++
		}
	}()

	time.Sleep(3 * time.Millisecond)
	t.Log(m)
}

func dupTest(t *testing.T) {
	sf := NewSnowFlake()
	m := make(map[uint64]struct{})
	go func() {
		for {
			id, err := sf.NextID()
			if err != nil {
				panic(err)
			}
			if _, ok := m[id]; ok {
				panic(fmt.Errorf("dup id: %d", id))
			}
			t.Log(id)
			m[id] = struct{}{}
		}
	}()

	time.Sleep(3 * time.Millisecond)
	t.Logf("大致 3 毫秒内，共生成 %d 个 id，没有重复", len(m))
}
