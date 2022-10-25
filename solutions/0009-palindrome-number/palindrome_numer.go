// https://leetcode.com/problems/palindrome-number/
package palindromenumber

func isPalindrome(x int) bool {
	orig := x
	var remainder int
	res := 0
	for x > 0 {
		remainder = x % 10
		res = (res * 10) + remainder
		x = x / 10
	}
	if orig == res {
		return true
	}
	return false
}
