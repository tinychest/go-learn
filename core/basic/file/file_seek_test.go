package file

import (
	"os"
	"testing"
)

// 详解 f.Seek 方法

// 一、首先从一个现象说起，你双击打开一个 txt 文档，编辑光标默认在最开头

// 二、原理
// 虽然注释写的很清楚，写下这里注释说明的时候，也看懂了，但是怕以后忘了
// Seek sets the offset for the next Read or Write on file to offset, interpreted
// according to whence:
// 0 means relative to the origin of the file
// 1 means relative to the current offset
// 2 means relative to the end.
// It returns the new offset and an error, if any.
// The behavior of Seek on a file opened with O_APPEND is not specified.

// 三、方法签名
// 参数 offset：向后位移多少
// 参数 whence：表示参考系 0=文件头 1=光标当前所在的位置 2=文件结尾
//  os.SEEK_SET、os.SEEK_CUR、os.SEEK_END（已过时）
//	io.SeekStart、io.SeekCurrent、io.SeekEnd
// 返回值 ret：移动后，参照参文件头的位移数
func TestSeek(t *testing.T) {
	seekTest(t)
	moveAndWriteTest(t)
}

func seekTest(t *testing.T) {
	var (
		f   *os.File
		ret int64
		err error
	)
	if f, err = os.OpenFile("D:/text.txt", os.O_RDWR, 0); err != nil {
		t.Log(err)
		return
	}

	// 假设文件内容是 123456
	// whence 0 返回 3
	// whence 1 返回 3
	// whence 2 返回 9
	if ret, err = f.Seek(3, 2); err != nil {
		t.Log(err)
		return
	}
	t.Log(ret)
	// whence 0 返回 4
	// whence 1 返回 7
	// whence 2 返回 10
	if ret, err = f.Seek(4, 2); err != nil {
		t.Log(err)
		return
	}
	t.Log(ret)
}

// 在空文件中，直接移动文件指针，进行写操作，看下效果
func moveAndWriteTest(t *testing.T) {
	var (
		f   *os.File
		err error
	)

	// 以 读写 模式打开文件
	if f, err = os.OpenFile("D:/text.txt", os.O_RDWR, 0); err != nil {
		t.Log(err)
		return
	}
	if _, err = f.Seek(3, 0); err != nil {
		t.Log(err)
		return
	}
	if _, err = f.WriteString("123"); err != nil {
		t.Log(err)
		return
	}
}
