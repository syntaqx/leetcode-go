// https://leetcode.com/problems/longest-palindromic-substring/
package longestpalindromicsubstring

func longestPalindrome(s string) string {
	res, maxL := "", 0
	for i := range s {
		l, r := i, i // odd
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxL {
				res = s[l : r+1]
				maxL = r - l + 1
			}
			l--
			r++
		}
		l, r = i, i+1 // even
		for l >= 0 && r < len(s) && s[l] == s[r] {
			if r-l+1 > maxL {
				res = s[l : r+1]
				maxL = r - l + 1
			}
			l--
			r++
		}
	}
	return res
}
