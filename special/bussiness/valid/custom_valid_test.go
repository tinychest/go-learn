package valid

import (
	"fmt"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strconv"
	"testing"
)

func TestCustomValid(t *testing.T) {
	if err := v.RegisterValidation("checkMobile", checkMobile); err != nil {
		t.Fatal(fmt.Errorf("注册自定义校验规则失败: %w", err))
	}

	// 参见：github.com/go-playground/validator/v10@v10.9.0/_examples/translations/main.go:103
	err := v.RegisterTranslation("checkMobile", tr, func(ut ut.Translator) error {
		return ut.Add("checkMobile", "{0} checkMobile failed!", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("checkMobile", fe.Field())

		return t
	})
	if err != nil {
		t.Fatal(fmt.Errorf("注册自定义翻译规则失败: %w", err))
	}
}

/* 校验结果中的信息仅是描述了，字段x 违反了 规则x 这样，具体的规则参数都没有提示出来，所以说自定义规则需要实现的接口方法很简单 */
func checkMobile(fl validator.FieldLevel) bool {
	mobile := strconv.Itoa(int(fl.Field().Uint()))
	re := `^1[3456789]\d{9}$`
	r := regexp.MustCompile(re)
	return r.MatchString(mobile)
}