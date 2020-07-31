package Week03

func letterCombinations(digits string) []string {
	mem := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	result := &[]string{}
	search("", digits, 0, mem, result)
	return *result
}

func search(s string, digits string, i int, mem map[string]string, result *[]string) {
	if i == len(digits) {
		*result = append(*result, s)
		return
	}
	//process
	stringTmp := mem[string(digits[i])]
	for _, value := range stringTmp {
		search(s+string(value), digits, i+1, mem, result)
	}
}
