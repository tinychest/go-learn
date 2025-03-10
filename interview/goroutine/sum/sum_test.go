package main

import (
	"runtime"
	"testing"
)

func TestNumGoroutine(t *testing.T) {
	t.Log(runtime.NumGoroutine())
}
