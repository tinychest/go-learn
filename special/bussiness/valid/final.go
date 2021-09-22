package valid

import (
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

func Validate(param interface{}) (map[string]string, bool) {
	res := map[string]string{}

	paramName := reflect.TypeOf(param).String()
	uselessPre := paramName[strings.LastIndex(paramName, "."):]

	if err := v.Struct(param); err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			rawField := e.Namespace()
			rawMsg := e.Translate(tr)

			field := strings.TrimLeft(rawField, uselessPre)
			msg := strings.TrimLeft(rawMsg, e.StructField())

			res[field] = msg
		}
		return res, false
	}
	return nil, true
}