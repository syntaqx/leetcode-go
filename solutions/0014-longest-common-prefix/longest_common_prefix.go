// https://leetcode.com/problems/longest-common-prefix/
package longestcommonprefix

import "sort"

func longestCommonPrefix(strs []string) string {
	var longestPrefix string = ""
	endPrefix := false

	if len(strs) > 0 {
		sort.Strings(strs)
		first := string(strs[0])
		last := string(strs[len(strs)-1])

		for i := 0; i < len(first); i++ {
			if !endPrefix && string(last[i]) == string(first[i]) {
				longestPrefix += string(last[i])
			} else {
				endPrefix = true
			}
		}
	}
	return longestPrefix
}
