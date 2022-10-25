// 7. Reverse Integer
//
// Given a signed 32-bit integer x, return x with its digits reversed. If
// reversing x causes the value to go outside the signed 32-bit integer range
// [-231, 231 - 1], then return 0.
//
// Assume the environment does not allow you to store 64-bit integers
// (signed or unsigned).
//
// Example 1:
//
// Input: x = 123
// Output: 321
// Example 2:
//
// Input: x = -123
// Output: -321
// Example 3:
//
// Input: x = 120
// Output: 21
//
// Constraints:
//
// -231 <= x <= 231 - 1
//
// https://leetcode.com/problems/reverse-integer/
package reverseinteger

import "math"

func reverse(x int) int {
	k := 1
	if x < 0 {
		k = -1
		x *= k
	}

	reverse := 0
	for x > 0 {
		reverse = (reverse * 10) + x%10
		x /= 10
	}

	if reverse > math.MaxInt32 || reverse < math.MinInt32 {
		return 0
	}

	return reverse * k
}
