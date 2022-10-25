package integertoroman

import (
	"strconv"
	"strings"
)

var conversions = []struct {
	value int
	digit string
}{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func intToRoman(num int) string {
	maxRomanNumber := 3999
	if num > maxRomanNumber {
		return strconv.Itoa(num)
	}

	var roman strings.Builder
	for _, conversion := range conversions {
		for num >= conversion.value {
			roman.WriteString(conversion.digit)
			num -= conversion.value
		}
	}

	return roman.String()
}
