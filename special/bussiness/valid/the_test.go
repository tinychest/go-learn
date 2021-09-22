package valid

import (
	"fmt"
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
			// fmt.Printf("Field: %s\n", e.Field())
			// fmt.Printf("Tag: %s\n", e.Tag())
			// fmt.Printf("Namespace: %s\n", e.Namespace())
			// fmt.Printf("StructNamespace: %s\n", e.StructNamespace())
			// fmt.Printf("StructField: %s\n", e.StructField())
			// fmt.Printf("Param: %s\n", e.Param())
			// fmt.Printf("ActualTag: %s\n", e.ActualTag())
			// fmt.Printf("Translate: %s\n", e.Translate(tr))
			// fmt.Println("-----------")
			fmt.Printf("Error: %s\n", e.Error())
		}
	}
}
