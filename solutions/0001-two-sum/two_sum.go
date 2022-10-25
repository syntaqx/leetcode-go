package twosum

// Time complexity: O(n)
// Space complexity: O(n)
func twoSum(nums []int, target int) []int {
	record := make(map[int]int)

	for i, j := range nums {
		complement := target - j
		if res, ok := record[complement]; ok {
			return []int{res, i}
		}
		record[j] = i
	}
	return []int{}
}

// brute force
// Time complexity: O(n^2)
// Space complexity: O(1)
func twoSum1(nums []int, target int) []int {
	for i, j := range nums {
		for k := i + 1; k < len(nums); k++ {
			if nums[k]+j == target {
				return []int{i, k}
			}
		}
	}
	return []int{}
}
