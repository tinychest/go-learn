package file

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"testing"
)

// 首先澄清一个容易掉入的误区，当你向文本内容 13 中插入一个 2，是不是 3 向后移动了，然后 2 才能放下去；
// 就像队列一样，当你希望做插入操作时，第一步不是直接将元素放到指定位置就好了，而是要将指定位置后边的元素全部向后移动
// 不然，会造成原来指定位置元素的丢失
// （为了更好理解，可以将程序代码中的文件写操作，理解为文本编辑中 Insert 模式下的写）

// 所以高级语言的方法类库都不会提供类似，删除一行、向指定文本为止插入内容 方法
// 那文本内容插入、删除的操作该怎么实现呢？
// 答：通过中间文件，将目标文件全部读出来（太大了可以读一些写一些）
// 将插入位置前面的内容写入临时文件，将目标内容写入临时文件，将剩下的内容写入临时文件

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

	// 方式三（手写）
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
		t.Log(err)
		return
	}

	if f, err = os.Create("D:/1.png"); err != nil {
		t.Log(err)
		return
	}
	defer f.Close()

	if err = writeFile(f, resp.Body); err != nil {
		t.Log(err)
		return
	}
	t.Log("下载图片成功")
}
