package style

import "github.com/xuri/excelize/v2"

// 默认字体
const (
	DefaultFontFamily        = "微软雅黑"
	DefaultTableHeadFontSize = 20
	DefaultTableBodyFontSize = 10
)

// 默认对齐策略
const (
	DefaultHorizontal      = "center"
	DefaultVertical        = "center"
	DefaultIndent          = 0
	DefaultJustifyLastLine = false
	DefaultReadingOrder    = 0
	DefaultRelativeIndent  = 0
	DefaultShrinkToFit     = false
	DefaultTextRotation    = 0
	DefaultWrapText        = false
)

// 默认保护策略
const (
	DefaultProtectionHidden = false
	DefaultProtectionLocked = true
)

var (
	DefaultBorder = []excelize.Border{
		{
			Type:  "top",
			Color: "#000000",
			Style: 2,
		},
		{
			Type:  "bottom",
			Color: "#000000",
			Style: 2,
		},
		{
			Type:  "left",
			Color: "#000000",
			Style: 2,
		},
		{
			Type:  "right",
			Color: "#000000",
			Style: 2,
		},
	}

	DefaultAlignment = &excelize.Alignment{
		Horizontal:      DefaultHorizontal,
		Indent:          DefaultIndent,
		JustifyLastLine: DefaultJustifyLastLine,
		ReadingOrder:    DefaultReadingOrder,
		RelativeIndent:  DefaultRelativeIndent,
		ShrinkToFit:     DefaultShrinkToFit,
		TextRotation:    DefaultTextRotation,
		Vertical:        DefaultVertical,
		WrapText:        DefaultWrapText,
	}

	DefaultProtection = &excelize.Protection{
		Hidden: DefaultProtectionHidden,
		Locked: DefaultProtectionLocked,
	}
)
