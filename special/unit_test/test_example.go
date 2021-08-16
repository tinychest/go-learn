package unit

import (
	"reflect"
	"testing"
)

// 需要根据测试的具体方法，进行调整更改
type FunctionParam struct {
}

// 就只支持单个返回值的，不整怪异的例子和没用的设计
type Case struct {
	// 要测试方法的方法名（必须 - 逻辑上非必须，错误提示用）
	FunctionName string
	// 要测试方法的参数（必须）
	FunctionParam FunctionParam
	// 要测试方法的返回值（非必须 - 如校验方法）
	ExpectResult interface{}
	// 要测试方法的返回错误（必须 - 用于测试的方法没个返回值类型为 error 的测试什么，但是并不是说方法实际调用一定会发生异常）
	ExpectError bool
}

// 不同的情况还是要根据实际的来，可能会有一些经常用到的功能方法，这里列举一下，实际使用的样例
func _(testcases []Case, t *testing.T) {
	// 这个函数完全是为了语法创建的
	exampleFunction := func(FunctionParam) (interface{}, error) {
		return nil, nil
	}
	for _, testcase := range testcases {
		t.Run(testcase.FunctionName, func(t *testing.T) {
			result, err := exampleFunction(testcase.FunctionParam)
			// 错误
			if (err != nil) != testcase.ExpectError {
				t.Errorf("%s() error = %v, expect %v", testcase.FunctionName, err, testcase.ExpectError)
			}
			// 返回值
			if !reflect.DeepEqual(result, testcase.ExpectResult) {
				t.Errorf("%s() result = %v, expect %v", testcase.FunctionName, result, testcase.ExpectResult)
			}
		})
	}
}
