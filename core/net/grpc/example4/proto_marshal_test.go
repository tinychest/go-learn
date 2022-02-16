package example4

import (
	"go-learn/core/net/grpc/example4/proto/hello"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/encoding/prototext"
	"google.golang.org/protobuf/proto"
	"testing"
)

// Marshal 方法的参数类型，就是 proto 文件生成的实体类型
// protobuf/encoding 下还有个 protowire 包（比较硬核的作用）

func TestProtoMarshal(t *testing.T) {
	// protoMarshal(t)
	// protoJSONMarshal(t)
	protoTextMarshal(t)
}

func protoMarshal(t *testing.T) {
	args := &hello.HelloArgs{ Value: "ping" }

	// 序列化
	bs, err := proto.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("【", string(bs), "】")

	// 反序列化
	newArgs := new(hello.HelloArgs)
	if err = proto.Unmarshal(bs, newArgs); err != nil {
		t.Fatal(err)
	}
	t.Log(newArgs.Value)
}

func protoJSONMarshal(t *testing.T) {
	args := &hello.HelloArgs{ Value: "ping" }

	// 序列化
	bs, err := protojson.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("【" + string(bs) + "】")

	// 反序列化
	newArgs := new(hello.HelloArgs)
	err = protojson.Unmarshal(bs, newArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newArgs.Value)
}

func protoTextMarshal(t *testing.T) {
	args := &hello.HelloArgs{ Value: "ping" }

	// 序列化
	bs, err := prototext.Marshal(args)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("【" + string(bs) + "】")

	// 反序列化
	newArgs := new(hello.HelloArgs)

	err = prototext.Unmarshal(bs, newArgs)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(newArgs.Value)
}
