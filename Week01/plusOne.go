package Week01

func plusOne(digits []int) []int {
	if len(digits) == 0 {
		return []int{1}
	}
	//从最高位开始递增
	p := len(digits) - 1
	var result []int
	for {
		if digits[p] == 9 {
			if p == 0 {
				result = append(result, 1)
				digits[p] = 0
				result = append(result, digits...)
				break
			} else {
				digits[p] = 0
				p--
			}
		} else {
			digits[p] = digits[p] + 1
			result = append(result, digits...)
			break
		}
	}
	return result
}
