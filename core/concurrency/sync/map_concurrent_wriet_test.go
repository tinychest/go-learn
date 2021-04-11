package sync

import (
	"sync"
	"testing"
)

const currentWriteSum = 10000

func TestMapCurrentRWTest(t *testing.T) {
	// fatal error: concurrent map writes
	// normalMapCurrentWrite()

	safeMapCurrentWrite()
}

// 普通 map 铺垫
var normalMap = make(map[string]string)

func normalMapWrite() {
	normalMap["name"] = "小明"
}

func normalMapCurrentWrite() {
	for i := 0; i < currentWriteSum; i++ {
		go normalMapWrite()
	}
}

// 并发安全 map 铺垫
var safeMap = sync.Map{}

func safeMapWrite() {
	safeMap.Store("name", "小明")
}

func safeMapCurrentWrite() {
	for i := 0; i < currentWriteSum; i++ {
		go safeMapWrite()
	}
}
