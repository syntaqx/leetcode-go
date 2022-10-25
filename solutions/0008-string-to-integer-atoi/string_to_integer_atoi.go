// https://leetcode.com/problems/string-to-integer-atoi/
package stringtointegeratoi

import "math"

func myAtoi(str string) int {
	var res int
	var sign int = 1
	var i, j int
	if len(str) > 0 {
		s := []byte(str)
		for i = 0; i < len(s); i++ {
			// trim spaces
			for s[i] == ' ' {
				i++
				if i >= len(s) {
					return 0
				}
			}
			// sign
			if i < len(s)-1 && !isNumber(s[i]) && isNumber(s[i+1]) {
				if s[i] == '-' {
					sign = -1
				} else if s[i] == '+' {
					sign = 1
				} else {
					return 0
				}
				i++
			}
			// check if need to continue parsing
			if !isNumber(s[i]) {
				return 0
			}
			// check leading zeroes
			for i < len(s) && isNumber(s[i]) && s[i] == '0' {
				i++
			}
			// cut number
			for i < len(s) && isNumber(s[i]) {
				j++
				i++
			}
			// check if the number is too big
			if j > 10 {
				return maxSigned(sign)
			}
			break
		}
		power := 1
		for t := 0; t < j; t++ {
			res = res + int(s[i-t-1]-'0')*power
			if res > math.MaxInt32 {
				return maxSigned(sign)
			}
			power *= 10
		}
	}
	return res * sign
}

func maxSigned(sign int) int {
	if sign > 0 {
		return math.MaxInt32
	} else {
		return math.MinInt32
	}
}

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}
