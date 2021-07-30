package marshal

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 对于 json 来说如果是 "+Inf"，json 类库可以处理，但是直接返回 +Inf，是 类库无法处理的，将得到 panic
// json 类库如何返回 err，详见 go\src\encoding\json\encode.go/575
func TestMarshalInf(t *testing.T) {
	f1, f2 := 1.0, 0.0
	f := f1 / f2

	if result, err := json.Marshal(f); err != nil {
		t.Fatal(err)
	} else {
		fmt.Println(string(result))
	}
}
