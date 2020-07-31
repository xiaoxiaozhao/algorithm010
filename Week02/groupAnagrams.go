package Week02

import (
	"sort"
	"strings"
)

//暴力法
//在数据量大时，会超过时间限制
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return [][]string{}
	}
	result := make([][]string, 0)
	result[0] = make([]string, 0)
	//append the first
	result[0] = append(result[0], strs[0])
	//base compare
	base := make([]string, 0)
	base = append(base, strs[0])
	for _, value := range strs {
		isExist := false
		idx := 0
		for _, baseItem := range base {
			if isAnagram(value, baseItem) {
				isExist = true
				break
			}
		}
		if isExist {
			result[idx] = append(result[idx], value)
		} else {
			tmpStrList := make([]string, 0)
			tmpStrList = append(tmpStrList, value)
			result = append(result, tmpStrList)
			base = append(base, value)
		}
	}
	return result
}

//分组法
func groupAnagrams1(strs []string) [][]string {
	var result [][]string
	base := make(map[string]int)
	if len(strs) == 0 {
		return result
	}
	for _, item := range strs {
		tmp := strings.Split(item, "")
		sort.Strings(tmp)
		baseStr := strings.Join(tmp, "")
		if findidx, ok := base[baseStr]; ok {
			result[findidx] = append(result[findidx], item)
		} else {
			tmpList := []string{item}
			result = append(result, tmpList)
			base[baseStr] = len(result) - 1
		}
	}
	return result
}
