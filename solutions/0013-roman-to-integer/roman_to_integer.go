package romantointeger

var romanNumerals = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	length := len(s)

	if length == 0 {
		return 0
	} else if length == 1 {
		return romanNumerals[s[0]]
	}

	sum := romanNumerals[s[length-1]]

	for i := length - 2; i >= 0; i-- {
		if romanNumerals[s[i]] < romanNumerals[s[i+1]] {
			sum -= romanNumerals[s[i]]
		} else {
			sum += romanNumerals[s[i]]
		}
	}

	return sum
}
