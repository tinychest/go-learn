package valid

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhTrans "github.com/go-playground/validator/v10/translations/zh"
)

var (
	v  = validator.New()
	tr ut.Translator
)

func init() {
	// 注册翻译
	tr, _ = ut.New(zh.New(), zh.New()).GetTranslator("zh")
	if err := zhTrans.RegisterDefaultTranslations(v, tr); err != nil {
		panic(err)
	}
}
