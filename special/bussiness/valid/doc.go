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

package valid

/*
校验概念：每个字段要是有多个校验规则标签，则按照顺序依次校验，直到遇到一个违法的标签
校验结果：针对 字段 - 校验标签 给出结果

《大全》
https://pkg.go.dev/gopkg.in/go-playground/validator.v10#hdr-Validation_Functions_Return_Type_error
详见：github.com/go-playground/validator/v10@v10.9.0/baked_in.go:70


- string
required（不能为空）
eq=<xxx>（等于）
ne=<xxx>（不等于）

numeric（string 只包含数字）
datetime=2006-01-02 15:04:05（string - datetime）
email（邮箱）
json（json）

- number
required（不能为 0）
eq=<n>（等于）
ne=<n>（不等于）

- complicated
startswith、endswith、base64 等等

- special
gt、gte、lt、lte、min、max（number、string、array、slice、map）
oneof=<xxx1> <xxx2>...（枚举）
length=<n>（string、array、slice、map）
required（必填 - 基础类型非空、指针类型不能为 nil）
eqfield=<xxx>（必须等同于指定字段的值 - 例如，密码、确认密码）
neqfield=<xxx>（不等同于指定字段的值）字段名必须完全匹配
dive（array、slice、map - 默认会校验 struct）
excludesall=0x2C（不包含英文逗号）
*/
