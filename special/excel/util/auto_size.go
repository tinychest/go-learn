package util

import "unicode/utf8"

func AutoWidth(fontSize float64, content string) float64 {
	return fontSize / 4 * float64(utf8.RuneCountInString(content))
}

func AutoHeight(fontSize float64) float64 {
	return fontSize + 10
}