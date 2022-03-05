package tcp

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"testing"
	"time"
)

// 演示通过自定义编解码方式解决粘包问题

const headLen = 4

func TestPasteHandle_Server(t *testing.T) {
	c, err := server()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// NOTE reader 需要放在这里
	//  可以通过在循环中打印 reader.Buffered() 了解到原因（和客户端的 Nagle 算法有关）
	reader := bufio.NewReader(c)

	// 读取客户端消息
	for {
		t.Log("等待客户端消息...")

		bs, err := decode(reader)
		if err = ErrWrap(err); err != nil {
			t.Fatal(err)
		}
		if len(bs) == 0 {
			t.Log("一条数据未读取完毕")
			continue
		}
		t.Log("读取到客户端消息：", string(bs))
	}
}

func decode(reader *bufio.Reader) ([]byte, error) {
	// Peek 不会 advancing the reader
	lengthBs, err := reader.Peek(headLen)
	if err != nil {
		return nil, err
	}

	// 读取到的 4 个字节数据，作为表示主体数据的长度
	var length int32
	err = binary.Read(bytes.NewBuffer(lengthBs), binary.LittleEndian, &length)
	if err != nil {
		return nil, err
	}

	// 检查缓存中是否是约定的一条完整的数据
	if reader.Buffered() < headLen+int(length) {
		// 没有一条完整的数据可以读取
		return nil, err
	}

	// Buffered returns the number of bytes that can be read from the current buffer.
	bs := make([]byte, headLen+length)
	_, err = reader.Read(bs)
	if err != nil {
		return nil, err
	}
	return bs[headLen:], nil
}

func TestPasteHandle_Client(t *testing.T) {
	c, err := client()
	if err != nil {
		t.Fatal(err)
	}
	defer MustClose(c)

	// 向服务端发送消息
	msg, err := encode("I'm Client!")
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < 3; i++ {
		_, err = c.Write(msg)
		if err != nil {
			t.Fatal(err)
		}
		t.Log("成功向服务端发送：", string(msg))
	}

	time.Sleep(time.Second)
}

func encode(msg string) ([]byte, error) {
	var length = int32(len(msg)) // 将数据的长度转成 固定 的 4 个字节存储

	var buffer = new(bytes.Buffer)
	buffer.Grow(int(headLen + length))

	// 将固定字节的长度，写入 []byte 中
	err := binary.Write(buffer, binary.LittleEndian, length)
	if err != nil {
		return nil, err
	}
	// 将消息写入 []byte
	_, err = buffer.WriteString(msg)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
