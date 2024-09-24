package backtrack

import (
	"fmt"
	"strconv"
)

//93.复原IP地址
//给定一个只包含数字的字符串，复原它并返回所有可能的 IP 地址格式。
//输入:s = "25525511135" 输出:["255.255.11.135","255.255.111.35"]
//输入:s = "0000" 输出:["0.0.0.0"]
//输入:s = "1111" 输出:["1.1.1.1"]
//输入:s = "010010" 输出:["0.10.0.10","0.100.1.0"]

func restoreIpAddresses(s string) []string {
	var res []string
	var path []string
	l := len(s)
	var backTrack func(start int)
	backTrack = func(start int) {
		if len(path) == 3 {
			last := s[start:]
			if isValid(last) {
				res = append(res, getStr(path, last))
			}
			return
		}

		for i := start; i < l; i++ {
			current := s[start : i+1]
			if !isValid(current) {
				return
			}
			path = append(path, current)
			backTrack(i + 1)
			path = path[:len(path)-1]
		}
	}
	backTrack(0)
	return res
}

func isValid(s string) bool {
	l := len(s)
	if l == 0 {
		return false
	}
	if l > 1 && s[0] == '0' {
		return false
	}
	n, _ := strconv.Atoi(s)
	if n > 255 {
		return false
	}
	return true
}

func getStr(s []string, last string) string {
	var res string
	for _, v := range s {
		res += v + "."
	}
	return res + last
}

func Handle10() {
	fmt.Println(restoreIpAddresses("010010"))
}
