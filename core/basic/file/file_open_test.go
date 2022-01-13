package file

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

/*
os.OpenFile(name string, flag int, perm FileMode)

[源码简译]
os.OpenFile 一般使用会选择通过 os.Create 或 os.Open 去替代
os.Create(name) = OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	只写（不可读）毫无意义；创建新文件的时候，只会去写，可能这个模式也就这个时候有用
os.Open(name) = OpenFile(name, O_RDONLY, 0)

[param flag]
// Flags to OpenFile wrapping those of the underlying system. Not all
// flags may be implemented on a given system.
const (
	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	O_RDONLY int = syscall.O_RDONLY // open the file read-only.（只读）
	O_WRONLY int = syscall.O_WRONLY // open the file write-only.（只写）
	O_RDWR   int = syscall.O_RDWR   // open the file read-write.（读写）
	// The remaining values may be or'ed in to control behavior.
	O_APPEND int = syscall.O_APPEND // append data to the file when writing.（写时，从文件末尾开始写）
	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.（文件存在时报错，配合上面使用，不能覆盖文件）
	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.（文件存在时清空内容，清空常规可写文件的内容）
)

[原理 - 也是由下面测试的结果]
perm 参数仅在创建文件时起作用，真实决定了，创建文件时，文件的权限

flag 参数决定了，程序中以何种模式操作。
	举例来说，假如一个文件只读，你希望以只写或者读写的方式打开，就直接告诉没有权限了；
	然后，如果以只读方式打开，没有问题，但是进行写操作，就会提示没有权限

[其他]
os.Chmod 方法改变权限
os.Chown 改变所有者
os.Symlink 创建软连接（Windows 中无法使用）
os.Lstat 返回文件信息，如果文件是软链接则返回软链接的信息，而不是软链接的信息
*/

func TestOpenFile(t *testing.T) {
	if err := openFileTest("D:/text.txt", "3"); err != nil {
		t.Log("操作失败-" + err.Error())
		return
	}
	t.Log("操作成功")
}

// os.OpenFile 默认行为
// - 找不到指定路径对应的文件，就返回 error
// - 文件的权限是在创建时，就指定好了的
func openFileTest(file, content string) error {
	const (
		perm1 = 0
		perm2 = 0777
		flag1 = os.O_CREATE | os.O_EXCL
		flag2 = os.O_CREATE | os.O_RDWR
		flag3 = os.O_CREATE | os.O_APPEND

		// flag1 + perm1
		// 第一次运行：文件不存在 - 创建（只读）并写入
		// 第二次运行：文件存在 - error "The file exists"

		// flag2 + perm1
		// 第一次运行：文件不存在 - 创建（只读）并写入
		// 第二次运行：文件存在 - error "Access is denied."

		// flag3 + perm1
		// 第一次运行：文件不存在 - 创建（只读）并写入
		// 第二次运行：文件存在 - error "Access is denied."

		// flag2 + perm2
		// 第一次运行：文件不存在 - 创建并写入
		// 第二次运行：文件存在 - 写（从文件头开始覆盖内容）

		// flag3 + perm2
		// 第一次运行：文件不存在 - 创建并写入
		// 第二次运行：文件存在 - 写（累加）

		// 结论，perm 最小 0200 就能保证不创建只读文件
	)
	var (
		f   *os.File
		err error
	)
	if f, err = os.OpenFile(file, os.O_RDWR, 0); err != nil {
		// 是否是权限错误
		if os.IsPermission(err) {
			return errors.New(fmt.Sprintf("权限错误：%s", err))
		}
		// 是否是没有找到文件错误
		if os.IsNotExist(err) {
			return errors.New(fmt.Sprintf("不存在错误：%s", err))
		}
		return errors.New(fmt.Sprintf("打开文件错误：%s", err))
	}

	if _, err = f.WriteString(content); err != nil {
		return errors.New(fmt.Sprintf("写入文件错误：%s", err))
	}
	_ = f
	return nil
}
