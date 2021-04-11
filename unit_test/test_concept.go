package unit

import (
	"testing"
)

// go 的测试，无论是文件名还是方法名都要要求
// 测试方法要求：func TestXxx(xxx *testing.T) {}
// 测试文件名要求：func xxx_test.go
func TestXxx(t *testing.T) {
	println("I can not run because file name is not xxx_test style...")
}
