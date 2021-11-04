package style

import "github.com/xuri/excelize/v2"

var DefaultTableBodyRow = &excelize.Style{
	Border: DefaultBorder,
	Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#d9d9d9"},
	},
	Font: &excelize.Font{
		Bold:      true,
		Italic:    false,
		Underline: "",
		Family:    DefaultFontFamily,
		Size:      DefaultTableBodyFontSize,
		Strike:    false,
		Color:     "#000000",
	},
	Alignment:  DefaultAlignment,
	Protection: DefaultProtection,
}
