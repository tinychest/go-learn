/*
一、https://github.com/gookit/validate
（国人开发，星星最少）
详见：https://github.com/gookit/validate/blob/master/README.zh-CN.md

二、https://github.com/go-ozzo/ozzo-validation
（不支持 struct 标签，星星倒数第二）

三、https://github.com/go-playground/validator
（gin 默认支持、GitHub 星星最多）

四、https://github.com/asaskevich/govalidator
（中规中矩）

本篇就介绍第三种
*/

package validate

/*
校验概念：每个字段要是有多个校验规则标签，则按照顺序依次校验，直到遇到一个违法的标签
校验结果：针对 字段 - 校验标签 给出结果

《大全》
https://pkg.go.dev/gopkg.in/go-playground/validator.v10#hdr-Validation_Functions_Return_Type_error
https://pkg.go.dev/github.com/go-playground/validator/v10#section-readme
详见：github.com/go-playground/validator/v10@v10.9.0/baked_in.go:70

- datetime=<时间格式>

- required（必填，基础类型非零值、指针类型不能为 nil）

- gt、gte、lt、lte、min、max（number、string、array、slice、map）

- length=<n>（string、array、slice、map）

- oneof=<xxx1> <xxx2>...（枚举）

- excludesall=0x2C（不包含英文逗号）

- unique：数组、切片、map 要求不能有重复的元素

- required_if <FieldName> <FieldValue>：当指定的字段的值为指定值时，当前字段才要求 required

- eqfield=<xxx>（必须等同于指定字段的值 - 例如，密码、确认密码）
- neqfield=<xxx>（不等同于指定字段的值）字段名必须完全匹配

- |：多个条件，满足一个即可（多个条件的错误，翻译插件是无法翻译的）

- dive（array、slice、map - 默认会校验 struct）
	dive required：修饰数组时，数组中每一个元素都要求 required
	required dive required：修饰数组时，要求数组本身 required
*/
