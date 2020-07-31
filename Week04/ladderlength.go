package Week04

import "strings"

type levelInfo struct {
	word  string
	level int
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	length := len(beginWord)
	//get the memcash
	combineDic := getAllCombineDic(wordList)
	//BFS
	Q := []levelInfo{{beginWord, 1}}
	visited := make(map[string]bool)
	for len(Q) != 0 {
		//queue
		tmp := Q[0]
		Q = Q[1:]
		word := tmp.word
		level := tmp.level
		for i := 0; i < length; i++ {
			newWord := word[0:i] + "*" + word[i+1:]
			if _, ok := combineDic[newWord]; ok {
				for _, value := range combineDic[newWord] {
					if strings.EqualFold(value, endWord) {
						return level + 1
					}
					if _, ok := visited[value]; !ok {
						visited[value] = true
						Q = append(Q, levelInfo{value, level + 1})
					}
				}
			}
		}

	}
	return 0
}

func getAllCombineDic(wordList []string) map[string][]string {
	//遍历每个单词
	res := make(map[string][]string)
	for _, value := range wordList {
		//构造词典
		for i := 0; i < len(value); i++ {
			newWord := value[0:i] + "*" + value[i+1:]
			if _, ok := res[newWord]; ok {
				res[newWord] = append(res[newWord], value)
			} else {
				res[newWord] = make([]string, 0)
				res[newWord] = append(res[newWord], value)
			}
		}
	}
	return res
}
