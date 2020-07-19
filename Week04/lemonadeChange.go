package Week04
func lemonadeChange(bills []int) bool {
	if len(bills) == 0 || bills[0] != 5{
		return false
	}
	count5 := 0
	count10 :=0
	for _,value := range bills{
		if value == 5{
			count5++
		}
		if value == 10{
			count10++
			if count5 <= 0{
				return false
			}
			count5--
		}
		if value==20{
			if count5 <= 0{
				return false
			}
			if count10 > 0{
				count10--
				count5--
			}else {
				if count5 >=3{
					count5 = count5 -3
				}else {
					return false
				}
			}
		}
	}
	return true
}