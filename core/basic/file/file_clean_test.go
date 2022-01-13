package file

import (
	"os"
	"testing"
)

// 详解 Truncate 方法

// 1、changes the size of the file（修改文件的大小，这里就涉及到操作系统的一些原理了）
// 2、设置成 0，就是清空文件
func TestClean(t *testing.T) {
	// clean1(t)
	// clean2(t)
}

// 简单直接的清空文件内容
func clean1(t *testing.T) {
	if err := os.Truncate("D:/text.txt", 0); err != nil {
		t.Log(err)
		return
	}
}

// 这里讲的是更贴近实际场景的，你想清空的文件，很有可能是在被并发操作的；又或者是清空完后，还需要继续读写操作的
// 不是说上面那样不用考虑这个问题，而是说你在知道没有并发操作以及继续操作的情况下，可以简单的直接调用上面的方法，否则也还是像下边这样操作
func clean2(t *testing.T) {
	var (
		f   *os.File
		err error
	)

	// 以 读写 模式打开文件
	if f, err = os.OpenFile("D:/text.txt", os.O_RDWR, 0); err != nil {
		t.Log(err)
		return
	}
	// 向文件中写入 123
	if _, err = f.WriteString("123"); err != nil {
		t.Log(err)
		return
	}
	// 删除文件所有内容
	if err = f.Truncate(0); err != nil {
		t.Log(err)
		return
	}
	// 向文件写入 456
	if _, err = f.WriteString("456"); err != nil {
		t.Log(err)
		return
	}

	// 此时，文件内容应该是 456，而实际发现是 "   456"（有3个空格）
	// 但是说，如果你有正在且已经操作过的文件指针，go 内存中的文件变量，记录着指针的位置，你希望向指定位置写时，发现文件中没有内容，只能填充空格，以达到从指定位置写的效果
	// 所以比较妥当的清空文件做法是，在清空文件后，也将文件操作指针也重置
	// （还可参见：moveAndWriteTest 方法）
	// if _, err = f.Seek(0, 0); err != nil {
	// 	t.Log(err)
	// 	return
	// }
}
