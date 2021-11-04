package style

import "github.com/xuri/excelize/v2"

var DefaultTableHead = &excelize.Style{
	Border: DefaultBorder,
	Fill: excelize.Fill{
		Type:    "pattern",
		Pattern: 1,
		Color:   []string{"#ee7a00"},
	},
	Font: &excelize.Font{
		Bold:      true,
		Italic:    false,
		Underline: "",
		Family:    DefaultFontFamily,
		Size:      DefaultTableHeadFontSize,
		Strike:    false,
		Color:     "#6cea97",
	},
	Alignment:  DefaultAlignment,
	Protection: DefaultProtection,
}
