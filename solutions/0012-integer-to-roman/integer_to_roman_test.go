package integertoroman

import "testing"

func TestIntToRoman(t *testing.T) {
	testCases := []int{
		3,
		4,
		9,
		58,
		1994,
		0,
	}
	expected := []string{"III", "IV", "IX", "LVIII", "MCMXCIV", ""}

	for index, s := range testCases {
		if res := intToRoman(s); res != expected[index] {
			t.Errorf("expected %s, got %s", expected[index], res)
		}
	}
}
