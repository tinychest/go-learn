package main

import (
	"go-learn/micro/distribute"
	"testing"
)

var snow = distribute.NewSnowFlake()

func TestLockAndUnlock(t *testing.T) {
	id, err := snow.NextID()
	if err != nil {
		t.Fatal(err)
	}

	// lock
	// 锁被占用；可以自旋，每隔一段时间自旋一次，自旋一定次数后，认定失败
	// 注：你业务不能说，这个怎么还能失败。请思考这个问题，开发人员从技术上尽可能减少失败的可能性，产品业务给出失败该如何处理
	if err := lockPre(lockKey, id); err != nil {
		t.Fatal(err)
	}

	// 业务操作...

	// unlock
	if err := unlock(lockKey, id); err != nil {
		t.Fatal(err)
	}
}

func TestUnlockFail(t *testing.T) {
	// 没有锁，解锁会报错（可以改）
	if err := unlock(lockKey, 1); err != nil {
		t.Fatal(err)
	}
}

func TestLock(t *testing.T) {
	if err := lockPre(lockKey, 1); err != nil {
		t.Fatal(err)
	}
}
