package greedy

import "fmt"

//给你一个字符串 s 。我们要把这个字符串划分为尽可能多的片段，同一字母最多出现在一个片段中。
//注意，划分结果需要满足：将所有划分结果按顺序连接，得到的字符串仍然是 s 。
//返回一个表示每个字符串片段的长度的列表。
// 示例 1： 输入：s = "ababcbacadefegdehijhklij" 输出：[9,7,8]
//示例 2： 输入：s = "eccbbbbdec" 输出：[10]

func partitionLabels(s string) []int {
	var res []int
	position := make([]int, 26)
	l := len(s)
	for i := 0; i < l; i++ {
		position[s[i]-'a'] = i
	}
	var left int
	var right int
	for i := 0; i < l; i++ {
		if right < position[s[i]-'a'] {
			right = position[s[i]-'a']
		}
		if i == right {
			res = append(res, right-left+1)
			left = i + 1
		}
	}

	return res
}

func Handle19() {
	fmt.Println(partitionLabels("eccbbbbde"))
}
