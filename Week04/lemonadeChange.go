package Week04
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 || bills[0] != 5{
		return false
	}
	count_5 := 0
	count_10:=0
	for _,value := range bills{
		if value == 5{
			count_5++
		}
		if value == 10{
			count_10++
			if count_5 <= 0{
				return false
			}
			count_5--
		}
		if value==20{
			if count_5 <= 0{
				return false
			}
			if count_10 > 0{
				count_10--
				count_5--
			}else {
				if count_5 >=3{
					count_5 = count_5-3
				}else {
					return false
				}
			}
		}
	}
	return true
}