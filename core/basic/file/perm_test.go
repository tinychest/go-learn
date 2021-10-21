package file

import (
	"fmt"
	"os"
	"testing"
)

// 0b 二进制
// 0|0o 八进制
// 0x 十六进制

// [参数 perm]（这个单词翻译为烫发）
// 在 linux 中使用 9 个 2 进制来表示文件权限，但是比起 0b111111111，还是 0777 更加易读，所以一般都会用八进制数来表示权限
// 在 windows 中，没有文件是不可读的，也就意味着我们能控制的，只有文件是否是只读（可写）的，而根据 file_open_test 中的现象
// 只要保证 perm x1x xxx xxx 中指定的位置为 1，就能保证文件是可写的+-

func TestPerm(t *testing.T) {
	printPerm(0777) // 111 111 111 → -rwxrwxrwx
	printPerm(0666) // 110 110 110 → -rw-rw-rw-
	printPerm(0644) // 110 100 100 → -rw-r--r--
}

func printPerm(perm int) {
	permStr := os.FileMode(perm).String()
	fmt.Println(permStr)
}
