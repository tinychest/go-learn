package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// 《json.Marshal》
// nil → "null"
// string → 和原来的值相比较，键的值和值的开头和结尾都多了双引号
func TestUnmarshalBasicType(t *testing.T) {
	j := `{"name":"xiaoming", "age":11}`

	// string
	if r, err := json.Marshal(j); err != nil {
		panic(err)
	} else {
		fmt.Println(string(r))
	}

	// nil
	if s, err := json.Marshal(nil); err != nil {
		panic(err)
	} else {
		fmt.Println(string(s) == "null")
	}
}
