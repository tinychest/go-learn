package file

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/* <os.OpenFile(name string, flag int, perm FileMode)>

（源码简译）
os.OpenFile 一般使用会选择通过 os.Create 或 os.Open 去替代
os.Create(name) = OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
os.Open(name) = OpenFile(name, O_RDONLY, 0)
*/

/* <参数 flag：>

（源码复制）
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
（读）                    O_RDONLY int = syscall.O_RDONLY // open the file read-only.
（写）                    O_WRONLY int = syscall.O_WRONLY // open the file write-only.
（读写）                  O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	// The remaining values may be or'ed in to control behavior.
（写时，默认从文件末尾开始写） O_APPEND int = syscall.O_APPEND // append data to the file when writing.
（创建，直接覆盖）           O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
（配合上面使用，不能覆盖文件） O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
（？）                     O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
（打开文件时清空文件内容？）   O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
)

常用的组合：
只写；目标文件不存在，则创建，存在，则报错         os.O_WRONLY|os.O_CREATE|os.O_EXCL
只写；目标文件不存在，则创建，写入时从结尾开始写    os.O_WRONLY|os.O_CREATE|os.APPEND
*/

/* <参数 perm>（这个单词翻译为烫发）

在 linux 中使用 9 个 2 进制来表示文件权限，于是 go 中就约定用一个 3 位 8 进制数来表示创建的文件权限
在 windows 上这个参数是无意义的

- 常见的权限位
（可以通过 os.FileMode(0xxx).String() 来查看 8 进制位对应的绝体权限）
0777=所有人拥有所有的读、写、执行权限
0666=创建了一个普通文件，所有人拥有对该文件的读、写权限，但是都不可执行
0644=创建了一个普通文件，文件所有者对该文件有读写权限，用户组和其他人只有读权限，都没有执行权限
os.Chmod 方法改变权限
os.Chown 改变所有者
os.Symlink 创建软连接（Windows 中无法使用）
os.Lstat 返回文件信息，如果文件是软链接则返回软链接的信息，而不是软链接的信息
*/

func TestOpenFile(t *testing.T) {
	if err := openFileTest(); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("操作成功")
}

func openFileTest() error {
	const (
		// 现象：文件存在 - 报错，文件不存在 - 下面的校验报错
		flag1 = os.O_WRONLY | os.O_CREATE | os.O_EXCL
		// 现象：覆盖了文件开头的内容
		flag2 = os.O_WRONLY | os.O_CREATE
		// 现象：向文件末尾追加了内容
		flag3 = os.O_WRONLY | os.O_CREATE | os.O_APPEND
	)
	var (
		f   *os.File
		err error
	)

	if f, err = os.OpenFile("D:/text.txt", flag3, 0); err != nil {
		// 是否是权限错误
		if os.IsPermission(err) {
			return errors.New(fmt.Sprintf("权限错误：%s", err))
		}
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("不存在错误：%s", err))
		}
		return errors.New(fmt.Sprintf("打开文件错误：%s", err))
	}

	if _, err = f.WriteString("a"); err != nil {
		return errors.New(fmt.Sprintf("写入文件错误：%s", err))
	}
	_ = f
	return nil
}
