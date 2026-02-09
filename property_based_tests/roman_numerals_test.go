package roman_numerals

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct{
		Arabic int
		Roman string
	}{
		{ 1,  "I"},
		{ 2,  "II"},
		{ 3,  "III"},
		{ 4,  "IV"},
		{ 5,  "V"},
		{ 6,  "VI"},
		{ 7,  "VII"},
		{ 8,  "VIII"},
		{ 9,  "IX"},
		{ 10,  "X"},
		{ 11,  "XI"},
		{ 13,  "XIII"},
		{ 15,  "XV"},
		{ 16,  "XVI"},
		{ 20,  "XX"},
		{ 21,  "XXI"},
		{ 25,  "XXV"},
		{ 30,  "XXX"},
		{ 31,  "XXXI"},
		{ 40,  "XL"},
		{ 44,  "XLIV"},
		{ 45,  "XLV"},
		{ 50,  "L"},
		{ 90,  "XC"},
		{ 100,  "C"},
		{ 400, "CD"},
		{ 500, "D"},
		{ 900, "CM"},
		{ 1000, "M"},
		{ 1984, "MCMLXXXIV"},
		{ 3999, "MMMCMXCIX"},
		{ 2014, "MMXIV"},
		{ 1006, "MVI"},
		{ 798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("Converts %d to %s", test.Arabic, test.Roman), func(t *testing.T) {
			got := ArabicToRoman(test.Arabic)
			want := test.Roman

			if got != want {
				t.Errorf("got %s want %s", got, want)
			}
		})
	}
}
