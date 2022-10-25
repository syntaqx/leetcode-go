// https://leetcode.com/problems/container-with-most-water/
package containerwithmostwater

import "math"

func maxArea(height []int) int {
	start := 0
	end := len(height) - 1

	var maxArea float64 = 0
	width := end

	for start < end {
		maxArea = math.Max(maxArea, math.Min(float64(height[start]), float64(height[end]))*float64(width))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
		width--
	}

	return int(maxArea)
}
