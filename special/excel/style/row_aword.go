package style

import "github.com/xuri/excelize/v2"

var DefaultRowAWord = &excelize.Style{
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
