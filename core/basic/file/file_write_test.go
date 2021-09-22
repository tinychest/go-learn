package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

// 程序代码中的文件写操作，相当于普通文本编辑的 Insert 模式下的写
// 好好理解其中原理：是在内存中的读写，最终写回到硬盘中的；写的要素是向指定的内存地址（文件指针）替换指定位置的内容
// 很容易就能够明白，硬要实现删除一行的方法是很不合理的，所以高级语言的方法类库都不会提供这样的方法
func writeFile(f *os.File, data io.ReadCloser) error {
	var err error

	// 方式一（io.copy）
	// if _, err = io.Copy(f, resp.Body); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 方式二（写文件）
	// var picData []byte
	// if picData, err = ioutil.ReadAll(resp.Body); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println("读取到响应的图片大小：", len(picData), "byte")
	// if _, err = f.Write(picData); err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// 方式三（原生 手写）
	var buf = make([]byte, 1024)
	var n = 1
	for i := 1; n > 0; i++ {
		if n, err = data.Read(buf); err != nil {
			if errors.Is(err, io.EOF) {
				return fmt.Errorf("文件读取完毕！")
			}
			return fmt.Errorf("读取响应数据失败：%w", err)
		}
		fmt.Printf("第 %d 次，读取到的字节数 %d ", i, n)
		if n, err = f.Write(buf[:n]); err != nil {
			return fmt.Errorf("数据写入文件失败：%w", err)
		}
		fmt.Printf("写到文件字节数 %d\n", n)
	}
	return nil
}

func TestDownloadPic(t *testing.T) {
	var (
		picUrl = "https://www.baidu.com/img/PCfb_5bf082d29588c07f842ccde3f97243ea.png"
		resp   *http.Response
		f      *os.File
		err    error
	)

	if resp, err = http.Get(picUrl); err != nil {
		fmt.Println(err)
		return
	}

	if f, err = os.Create("D:/1.png"); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if err = writeFile(f, resp.Body); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("下载图片成功")
}
