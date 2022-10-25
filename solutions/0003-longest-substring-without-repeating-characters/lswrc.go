// Given a string, find the length of the longest substring without repeating characters.
//
// Example 1:
// Input: s = "abcabcbb"
// Output: 3
// Explanation: The answer is "abc", with the length of 3.
//
// Example 2:
// Input: s = "bbbbb"
// Output: 1
// Explanation: The answer is "b", with the length of 1.
//
// Example 3:
// Input: s = "pwwkew"
// Output: 3
// Explanation: The answer is "wke", with the length of 3.
// Notice that the answer must be a substring, "pwke" is a subsequence and not a substring.
//
// Constraints:
// 0 <= s.length <= 5 * 104
// Consists of English letters, digits, symbols and spaces.
package lswrc

func lengthOfLongestSubstring(s string) int {
	start := -1
	lenMax := 0
	dict := [256]int{}
	for i := range dict {
		dict[i] = -1
	}
	for i, r := range s {
		if v := dict[r]; v > start {
			start = v
		}
		length := i - start
		if length > lenMax {
			lenMax = length
		}
		dict[r] = i
	}
	return lenMax
}
