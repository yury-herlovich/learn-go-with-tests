package numeral

import (
	"strings"
)

type RomanNumeral struct {
	Value  int
	Symbol string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumerals = RomanNumerals{
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

func ConvertToRoman(n int) string {
	var result strings.Builder

	for _, numeral := range allRomanNumerals {
		for n >= numeral.Value {
			result.WriteString(numeral.Symbol)
			n -= numeral.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	total := 0

	for i := 0; i < len(roman); i++ {
		s := roman[i]

		// check next symbol
		if i+1 < len(roman) {
			combinationValue := allRomanNumerals.ValueOf(string([]byte{s, roman[i+1]}))

			if combinationValue != 0 {
				total += combinationValue
				i++

				continue
			}
		}

		total += allRomanNumerals.ValueOf(string(s))
	}

	return total
}
