package _0interpreter

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type IExpression interface {
	Interpreter(stats map[string]float64) bool
}

type AlterRule struct {
	expression IExpression
}

func NewAlterRule(rule string) (*AlterRule, error) {
	exp, err := NewAndExpression(rule)
	return &AlterRule{expression: exp}, err
}

func (r *AlterRule) Interpreter(stats map[string]float64) bool {
	return r.expression.Interpreter(stats)
}

// LessExpression expression 实现
type LessExpression struct {
	key   string
	value float64
}

func (e *LessExpression) Interpreter(stats map[string]float64) bool {
	v, ok := stats[e.key]
	if !ok {
		return false
	}
	return v < e.value
}

func NewLessExpression(exp string) (*LessExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != "<" {
		return nil, fmt.Errorf("less exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("less exp is invalid: %s", exp)
	}

	return &LessExpression{
		key:   data[0],
		value: val,
	}, nil
}

// GreaterExpression expression 实现
type GreaterExpression struct {
	key   string
	value float64
}

func (e *GreaterExpression) Interpreter(stats map[string]float64) bool {
	v, ok := stats[e.key]
	if !ok {
		return false
	}
	return v > e.value
}

func NewGreaterExpression(exp string) (*GreaterExpression, error) {
	data := regexp.MustCompile(`\s+`).Split(strings.TrimSpace(exp), -1)
	if len(data) != 3 || data[1] != ">" {
		return nil, fmt.Errorf("greater exp is invalid: %s", exp)
	}

	val, err := strconv.ParseFloat(data[2], 10)
	if err != nil {
		return nil, fmt.Errorf("greater exp is invalid: %s", exp)
	}

	return &GreaterExpression{
		key:   data[0],
		value: val,
	}, nil
}

// AndExpression expression 实现
type AndExpression struct {
	expressions []IExpression
}

func (e *AndExpression) Interpreter(stats map[string]float64) bool {
	for _, v := range e.expressions {
		if !v.Interpreter(stats) {
			return false
		}
	}
	return true
}

func NewAndExpression(exp string) (*AndExpression, error) {
	exps := strings.Split(exp, "&&")
	expressions := make([]IExpression, len(exps))

	for i, e := range exps {
		var expression IExpression
		var err error

		switch {
		case strings.Contains(e, ">"):
			expression, err = NewGreaterExpression(e)
		case strings.Contains(e, "<"):
			expression, err = NewLessExpression(e)
		default:
			err = fmt.Errorf("and exp is invalid: %s", exp)
		}
		if err != nil {
			return nil, err
		}

		expressions[i] = expression
	}
	return &AndExpression{expressions: expressions}, nil
}
