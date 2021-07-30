package valid

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	// "github.com/go-playground/validator"
	zhTrans "github.com/go-playground/validator/translations/zh"
	"strings"
	"testing"
)

// 官方文档有，但实际不行 required_id json datetime
// min max gt gte lt lte（对 string number 类型都有效）
// dive（只有添加标签 才会对 map slice struct 的内容进行校验）
// excludesall=0x2C（不包含逗号，不直接写逗号是因为官网推荐 Use the UTF-8 hex representation）
// 题外：对零值不是 nil 的类型添加校验标签 required 是白加的
type Form struct {
	Name            string  `validate:"required,lt=10"`
	Email           string  `validate:"email"`
	Url             string  `validate:"required,base64url,lt=50"`
	Inner           Inner   `validate:"dive"`
	Inners          []Inner `validate:"required,dive"` // required 只校验该字段是否为 nil
	Password        string  `validate:"required,base64,gt=6"`
	ConfirmPassword string  `validate:"eqfield=Password"`
}

type Inner struct {
	Phone string `validate:"numeric,len=11"`
}

func TestValidate(t *testing.T) {
	f := Form{
		Name: "小明",
		Url:  "",
		Password: "123456",
		ConfirmPassword: "456789",
		Inners: []Inner{
			{
				Phone: "",
			},
		},
	}

	v := validator.New()
	// 注册错误提示的翻译器（中文）
	// 1. GetTranslator 方法很蠢，首先第二个返回值就是返回 false，第一个值居然也正常返回内容
	// 2. RegisterDefaultTranslations 方法很蠢，这里并不是说默认生效，而是将其注册到允许使用的范围内（实际获取中文错误提示，还需要在获取校验错误时，手动调用）
	// 3. validator 包有两个，为什么是 v9 因为 zhTrans 形参类型要求
	tr, _ := ut.New(zh.New(), zh.New()).GetTranslator("zh")
	if err := zhTrans.RegisterDefaultTranslations(v, tr); err != nil {
		t.Fatal(err)
	}

	// 自定义错误提示（字段提示首字母小写）
	// v.RegisterTagNameFunc(func(field reflect.StructField) string {
	// 	// 1、先是自定义的 label
	// 	if label := field.Tag.Get("label"); len(label) != 0 {
	// 		return label
	// 	}
	// 	// 2、然后时 beego gin 都支持的表单参数反序列化标签（实际的参数名）
	// 	if form := field.Tag.Get("form"); len(form) != 0 {
	// 		return form
	// 	}
	// 	// 3、再就是 json
	// 	if j := field.Tag.Get("json"); len(j) != 0 {
	// 		return j
	// 	}
	// 	// 4、最后是实际变量名的首字母小写（json 时的实际参数名）
	// 	b := strings.Builder{}
	// 	b.Grow(len(field.Name))
	// 	b.WriteString(strings.ToLower(string(field.Name[0])))
	// 	b.WriteString(field.Name[1:])
	// 	return b.String()
	// })

	// 开始校验
	result := make(map[string]string)
	if err := v.Struct(f); err != nil {
		validErr := err.(validator.ValidationErrors)
		for _, e := range validErr {
			// e.Tag() 校验标签
			// e.Field() 校验提示的字段名
			field := e.Field()
			// 将具体错误提示的重复含义的开头去掉
			result[field] = strings.Replace(e.Translate(tr), field, "", 1)
		}
	}

	fmt.Println(result)
}
