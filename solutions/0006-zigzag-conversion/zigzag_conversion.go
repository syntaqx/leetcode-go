// https://leetcode.com/problems/zigzag-conversion/
package zigzagconversion

func convert(s string, numRows int) string {
	if numRows <= 1 {
		return s
	}

	rows := make([]string, numRows)

	goingDown := false
	row := 0

	for _, char := range s {
		rows[row] += string(char)

		if (row == 0) || (row == numRows-1) {
			goingDown = !goingDown
		}

		if goingDown {
			row++
		} else {
			row--
		}
	}

	result := ""
	for _, row := range rows {
		result += row
	}

	return result
}
