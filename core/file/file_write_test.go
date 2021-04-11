package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

// go 中（可能 c 也是）关于文件写，就是说站在程序角度上操作文件，就不会像平时操作的那样，什么修改，删除一行（甚至没有行的概念，都是认为规定分隔符为一行）
//  文章论坛关于删除和替换一行的做法就是将文件内容全部读取出来进行修改删除，然后写入到一个新的文件
//  文件内容修改里边的内容相当于键盘上显示的 Insert 模式
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
