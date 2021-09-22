package valid

type Form struct {
	Num             string `validate:"numeric"`
	Neq1            string
	Neq2            string `validate:"nefield=Neq2"`
	OneOf           string `validate:"oneof=a b"`
	Ne              int    `validate:"ne=1"`
	NeS             string `validate:"ne=1"`
	Eq              int    `validate:"eq=10"`
	EqS             string `validate:"eq=10"`
	Name            string `validate:"required,lt=10"`
	Email           string `validate:"email"`
	Url             string `validate:"required,lt=50"`
	Inner           Inner
	Inners          []Inner `validate:"min=1,dive"` // required 只校验该字段是否为 nil，gin 中不传相关字段或者相关字段为 nil，required 可以拦截下来，但是传了 [] 就拦不住了，因为 != nil（len = 0, cap = 0）
	Password        string  `validate:"required,base64,gt=6"`
	ConfirmPassword string  `validate:"eqfield=Password"`
}

type Inner struct {
	Phone string `validate:"numeric,len=11"`
}
