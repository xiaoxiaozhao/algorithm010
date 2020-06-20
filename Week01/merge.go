package Week01

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		for idx, item := range nums2 {
			nums1[idx] = item
		}
	} else if n == 0 {
		return
	}
	p1 := m - 1
	p2 := n - 1

	p := m + n - 1

	for {
		if p1 < 0 || p2 < 0 {
			break
		}
		if nums1[p1] < nums2[p2] {
			nums1[p] = nums2[p2]
			p2--
		} else {
			nums1[p] = nums1[p1]
			p1--
		}
		p--
	}
	if p1 >= 0 && p2 < 0 {
		return
	}
	for p2 >= 0 {
		nums1[p2] = nums2[p2]
		p2--
	}
}
