package distribute

import (
	"errors"
	"sync"
	"time"
)

// 雪花算法产生的背景是 twitter 为了应对高并发环境下的分布式系统对唯一 ID 生成的需求
//
// [特点]
// - 算法效率高
// - 生成的 ID 具有 时序性 和 唯一性
// - 不依赖任何三方库或者中间件
//
// [详解]
// 雪花算法产生的 ID 是一个 64 位的整形数字
// 它的位数具有这样的含义：
// 1 位（最高位）不用
// 41 位代表时间戳 → 41 个二进制位能代表的正整数为 0 - 2^41 - 1，假如用来表示毫秒值，转化为年 (2^41-1) / (1000 * 60 * 60 * 24 * 365) = 69 年
// （实际，你可以在代码中定义一个起始的时间的时间戳作为常量，所有实时生成的 ID 的时间戳都减去改值后作为实际的时间戳的值，也就是支持从指定常量的时间开始，向后的 69 年时间）
// 10 位代表机器 id
// 12 位代表序号 → 12 个二进制位能表示 2^12 - 1 = 4095 个 ID
// 意味着雪花算法支持：一个机器在一毫秒内最多产生 4095 个 ID
//
// [指定任意低位为 1，其他为 0 的操作]
// int64(-1) ^ (int64(-1) << n)
// -1：0000 0001 -- 补码 --> 1111 1111（负数以补码形式表示）
// 1111 1111 ^ 1111 0000 → 0000 1111

const (
	WorkerId       = 1
	StartTimestamp = 1654041600000 // 2022-06-01 00:00:00.000

	TimestampBits  = uint64(41)
	WorkerIdBits   = uint64(10) // 在实际分布式生产场景中，10 位会分为 5 位数据中心 ID，5 位节点 ID，目前明显不需要，就直接简化成节点 ID 了
	SequenceIdBits = uint64(12)

	MaxMachineId = int64(-1) ^ (int64(-1) << WorkerIdBits)
	MaxSequence  = int64(-1) ^ (int64(-1) << SequenceIdBits)

	TimestampShift = WorkerIdBits + SequenceIdBits
	MachineIdShift = SequenceIdBits
)

type SnowFlake struct {
	sync.Mutex
	LastStamp int64 // 上一个生成 ID 的时间戳
	WorkerId  int64 // 节点 id
	Sequence  int64 // 当前毫秒内，上一个生成 ID 的序号
}

func NewSnowFlake() *SnowFlake {
	return &SnowFlake{
		LastStamp: 0,
		WorkerId:  WorkerId,
		Sequence:  0,
	}
}

func (w *SnowFlake) NextID() (uint64, error) {
	w.Lock()
	defer w.Unlock()

	return w.nextID()
}

func (w *SnowFlake) nextID() (uint64, error) {
	timeStamp := time.Now().UnixMilli()
	if timeStamp < w.LastStamp {
		return 0, errors.New("time is moving backwards,waiting until")
	}

	if w.LastStamp == timeStamp {
		w.Sequence = (w.Sequence + 1) & MaxSequence

		// 如果已经达到当前毫秒内能生成的 ID 最大数量，就等待至下一秒
		// 实际可以预估并发量，适当腾挪
		if w.Sequence == 0 {
			for timeStamp <= w.LastStamp {
				timeStamp = time.Now().UnixMilli()
			}
		}
	} else {
		w.Sequence = 0
	}

	w.LastStamp = timeStamp
	id := ((timeStamp - StartTimestamp) << TimestampShift) |
		(w.WorkerId << MachineIdShift) |
		w.Sequence

	return uint64(id), nil
}
