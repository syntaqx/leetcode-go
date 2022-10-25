package motsa

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l := len(nums1) + len(nums2)
	if l == 0 {
		panic("nums1 and nums2 cannot be both empty.")
	}
	if l%2 == 1 {
		return findKth(nums1, nums2, l/2)
	}
	return (findKth(nums1, nums2, l/2-1) + findKth(nums1, nums2, l/2)) / 2
}

func findKth(a, b []int, k int) float64 {
	if len(a) > len(b) {
		a, b = b, a
	}

	if len(a) == 0 {
		return float64(b[k])
	}

	switch k {
	case len(a) + len(b) - 1:
		return float64(MaxOfTwo(a[len(a)-1], b[len(b)-1]))
	}

	pa := MinOfTwo(len(a)-1, k/2)  // k/2
	pb := MinOfTwo(len(b)-1, k-pa) // k - k/2

	if a[pa] < b[pb] {
		return findKth(a[pa:], b[:pb], pb)
	}
	return findKth(a[:pa], b[pb:], pa)
}

func MaxOfTwo(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func MinOfTwo(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
