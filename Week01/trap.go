package Week01

func trap(height []int) int {
	if len(height) == 0{
		return 0
	}
	var result int
	left := 0
	for {
		if left < len(height)-1 && height[left+1]>=height[left]{
			left++
		} else {
			break
		}
	}
	right := len(height)-1
	for {
		if right > 0 && height[right-1] >= height[right]{
			right--
		}else {
			break
		}

	}
	for  {
		if left>=right{
			break
		}
		lefthight:= height[left]
		righthight:=height[right]
		if lefthight <= righthight {
			for {
				if left >= right{
					break
				}
				left++
				if height[left]<lefthight{
					result = result + lefthight - height[left]
				}else {
					break
				}
			}
		}else {
			for{
				if left >=right{
					break
				}
				right--
				if height[right] < righthight {
					result = result + righthight - height[right]
				}else {
					break
				}
			}
		}

	}
	return result
}