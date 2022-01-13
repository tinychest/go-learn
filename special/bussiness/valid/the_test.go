package valid

import (
	"github.com/go-playground/validator/v10"
	"testing"
)

func TestValidate(t *testing.T) {
	f := Form{
		Num:             "07",
		Neq1:            "1",
		Neq2:            "1",
		OneOf:           "a",
		Ne:              2,
		NeS:             "2",
		Eq:              10,
		EqS:             "10",
		Name:            "小明",
		Url:             "",
		Password:        "123456",
		ConfirmPassword: "456789",
		Inner: Inner{
			Phone: "",
		},
		Inners: []Inner{
			{
				Phone: "",
			},
		},
	}

	if err := v.Struct(f); err != nil {
		validErrs := err.(validator.ValidationErrors)
		for _, e := range validErrs {
			// t.Logf("Field: %s\n", e.Field())
			// t.Logf("Tag: %s\n", e.Tag())
			// t.Logf("Namespace: %s\n", e.Namespace())
			// t.Logf("StructNamespace: %s\n", e.StructNamespace())
			// t.Logf("StructField: %s\n", e.StructField())
			// t.Logf("Param: %s\n", e.Param())
			// t.Logf("ActualTag: %s\n", e.ActualTag())
			// t.Logf("Translate: %s\n", e.Translate(tr))
			// t.Log("-----------")
			t.Logf("Error: %s\n", e.Error())
		}
	}
}
