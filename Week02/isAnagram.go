package Week02

func isAnagram(s string, t string) bool {
	mapMem := make(map[int32]int)
	if len(s) != len(t) {
		return false
	}
	for _, item := range s {
		if value, ok := mapMem[item]; ok {
			mapMem[item] = value + 1
		} else {
			mapMem[item] = 1
		}
	}
	//--
	for _, item := range t {
		if value, ok := mapMem[item]; ok {
			if value == 1 {
				delete(mapMem, item)
			} else {
				mapMem[item] = value - 1
			}
		}
	}
	if len(mapMem) == 0 {
		return true
	} else {
		return false
	}
}
