package numeral

import (
	"fmt"
	"testing"
)

func TestRomanNumerals(t *testing.T) {
	cases := []struct {
		// Description string
		Arabic int
		Want   string
	}{
		{1, "I"},
		{2, "II"},
		{3, "III"},
		{4, "IV"},
		{5, "V"},
		{6, "VI"},
		{7, "VII"},
		{8, "VIII"},
		{9, "IX"},
		{10, "X"},
		{14, "XIV"},
		{18, "XVIII"},
		{20, "XX"},
		{39, "XXXIX"},
		{40, "XL"},
		{47, "XLVII"},
		{49, "XLIX"},
		{50, "L"},
		{100, "C"},
		{90, "XC"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
		{1984, "MCMLXXXIV"},
		{3999, "MMMCMXCIX"},
		{2014, "MMXIV"},
		{1006, "MVI"},
		{798, "DCCXCVIII"},
	}

	for _, test := range cases {
		t.Run(fmt.Sprintf("convert %d to %s", test.Arabic, test.Want), func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)
			assertStrings(t, got, test.Want)
		})
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("expected %q but got %q", want, got)
	}
}
