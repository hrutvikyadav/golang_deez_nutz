package roman_numerals

import "strings"

type RomanNumeral struct {
	Value int
	Symbol string
}

var rnSymbols = []RomanNumeral{
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

func ArabicToRoman(arabic int) (roman string) {
	var result strings.Builder

	for _, rn := range rnSymbols {
		for arabic >= rn.Value {
			result.WriteString(rn.Symbol)
			arabic -= rn.Value
		}
	}
	// for arabic > 0 {
	// 	switch {
	// 	case arabic > 9:
	// 		result.WriteString("X")
	// 		arabic -= 10
	// 	case arabic > 8:
	// 		result.WriteString("IX")
	// 		arabic -= 9
	// 	case arabic > 4:
	// 		result.WriteString("V")
	// 		arabic -= 5
	// 	case arabic > 3:
	// 		result.WriteString("IV")
	// 		arabic -= 4
	// 	default:
	// 		result.WriteString("I")
	// 		arabic--
	// 	}
	// }
	return result.String()
}
