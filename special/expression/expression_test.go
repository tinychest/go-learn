package expression

import (
	"github.com/Knetic/govaluate"
	"math"
	"testing"
)

func TestApi(t *testing.T) {
	expr := "a/b + a"

	ee, err := govaluate.NewEvaluableExpression(expr)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(ee.Vars()) // 不会去重
	t.Log(ee.ToSQLQuery())
	t.Log(ee.Tokens()) // 解析详情
}

func TestExpression(t *testing.T) {
	var res interface{}

	// 算术表达式
	res = expressionTest(t, `a/b`, map[string]interface{}{
		"a": 1,
		"b": 2,
	})
	t.Log(res)

	// 逻辑表达式：通过 表达式的 Var 方法可以了解到 nil 会被视为参数
	res = expressionTest(t, `i34==0?0:nil`, map[string]interface{}{
		"i34": 1,
		"nil": nil,
	})
	t.Log(res)

	// 算术表达式：除数为 0（不会 panic）
	res = expressionTest(t, `a/b`, map[string]interface{}{
		"a": 1,
		"b": 0,
	})
	t.Log(res, math.IsInf(res.(float64), 0))

	// 算术表达式：除数为0、相减
	res = expressionTest(t, `a/b - a/b`, map[string]interface{}{
		"a": 1,
		"b": 0,
	})
	t.Log(res, math.IsNaN(res.(float64)))

	// 算术表达式：除数为0、相加
	res = expressionTest(t, `a/b + a/b`, map[string]interface{}{
		"a": 1,
		"b": 0,
	})
	t.Log(res, math.IsInf(res.(float64), 0))

	// 算术表达式：参数值为 nil（除号左右应该是 number，否则 Evaluate 会返回错误）
	// res = expressionTest(t, `a/b`, map[string]interface{}{
	// 	"a": nil,
	// 	"b": 1,
	// })
	// t.Log(res)

}

func expressionTest(t *testing.T, formula string, parameters map[string]interface{}) interface{} {
	expr, err := govaluate.NewEvaluableExpression(formula)
	if err != nil {
		t.Fatal(err)
	}

	res, err := expr.Evaluate(parameters)
	if err != nil {
		t.Fatal(err)
	}
	return res
}
