package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

// 程序代码中的文件写操作和现实逻辑类比起来，相当于 Insert 模式下的写
// 好好理解其中原理：是在内存中的读写，最终写回到硬盘中的；写的要素是向指定的内存地址（文件指针）替换指定位置的内容
// 很容易就能够明白，硬要实现删除一行的方法是很不合理的，所以高级语言的方法类库都不会提供这样的方法
func TestFileWrite(t *testing.T) {

}

func TestDownloadPic(t *testing.T) {
	var (
		picUrl = "http://thirdqq.qlogo.cn/g?b=oidb&k=hLWLicG3Fibrp2coiahSg4oKA&s=640&t=1555148173"
		resp   *http.Response
		f      *os.File
		err    error
	)

	if resp, err = http.Get(picUrl); err != nil {
		fmt.Println(err)
		return
	}
	_ = resp

	// 将图片保存到指定路径
	if f, err = os.Create("D:/1.png"); err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

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
		if n, err = resp.Body.Read(buf); err != nil {
			if errors.Is(err, io.EOF) {
				fmt.Println("文件读取完毕！")
				return
			}
			fmt.Printf("读取响应数据失败：%s\n", err)
			return
		}
		fmt.Printf("第 %d 次，读取到的字节数 %d ", i, n)
		if n, err = f.Write(buf[:n]); err != nil {
			fmt.Printf("数据写入文件失败：%s\n", err)
			return
		}
		fmt.Printf("写到文件字节数 %d\n", n)
	}
	fmt.Println("下载图片成功")
}
