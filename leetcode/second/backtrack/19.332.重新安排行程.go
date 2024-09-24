package backtrack

import (
	"fmt"
	"sort"
)

// 给你一份航线列表 tickets ，其中 tickets[i] = [fromi, toi] 表示飞机出发和降落的机场地点。请你对该行程进行重新规划排序。
// 所有这些机票都属于一个从 JFK（肯尼迪国际机场）出发的先生，所以该行程必须从 JFK 开始。如果存在多种有效的行程，请你按字典排序返回最小的行程组合。
// 例如，行程 ["JFK", "LGA"] 与 ["JFK", "LGB"] 相比就更小，排序更靠前。
// 假定所有机票至少存在一种合理的行程。且所有的机票 必须都用一次 且 只能用一次。
func findItinerary(tickets [][]string) []string {
	var res []string
	l := len(tickets)
	var backTrack func(tickets [][]string, from string) bool

	backTrack = func(tickets [][]string, from string) bool {
		if len(tickets) == 0 {
			if len(res) == l+1 {
				res = append(res, from)
				return true
			}
			return false
		}
		indexS := getIndex(tickets, from)
		if len(indexS) == 0 {
			return false
		}
		for i := 0; i < len(indexS); i++ {
			index := indexS[i]
			temp := make([][]string, len(tickets))
			copy(temp, tickets)
			to := tickets[index][1]
			newTickets := append(temp[:index], temp[index+1:]...)
			res = append(res, from)
			if !backTrack(newTickets, to) {
				res = res[:len(res)-1]
				continue
			}
			return true
		}
		return true
	}
	backTrack(tickets, "JFK")
	return res
}

func getIndex(tickets [][]string, from string) []int {
	var res []int
	for i, v := range tickets {
		if v[0] == from {
			res = append(res, i)
		}
	}
	// 根据每个票的第二个元素对索引进行排序
	sort.Slice(res, func(i, j int) bool {
		return tickets[res[i]][1] < tickets[res[j]][1]
	})
	return res
}

func Handle19() {
	//{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}
	//{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}   --[JFK ATL JFK SFO ATL SFO]
	//{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}   --[JFK NRT JFK KUL]
	//tickets =
	//[["EZE","TIA"],["EZE","HBA"],["AXA","TIA"],["JFK","AXA"],["ANU","JFK"],["ADL","ANU"],["TIA","AUA"],["ANU","AUA"],["ADL","EZE"],["ADL","EZE"],["EZE","ADL"],["AXA","EZE"],["AUA","AXA"],["JFK","AXA"],["AXA","AUA"],["AUA","ADL"],["ANU","EZE"],["TIA","ADL"],["EZE","ANU"],["AUA","ANU"]]
	//
	//添加到测试用例
	//输出
	//["JFK","AXA","AUA","ADL","ANU","AUA","ANU","EZE","ADL","EZE","ANU","JFK","AXA","EZE","TIA","ADL"]
	//预期结果
	//["JFK","AXA","AUA","ADL","ANU","AUA","ANU","EZE","ADL","EZE","ANU","JFK","AXA","EZE","TIA","AUA","AXA","TIA","ADL","EZE","HBA"]
	fmt.Println(findItinerary([][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}))
}

//type pair struct {
//	target  string
//	visited bool
//}
//type pairs []*pair
//
//func (p pairs) Len() int {
//	return len(p)
//}
//func (p pairs) Swap(i, j int) {
//	p[i], p[j] = p[j], p[i]
//}
//func (p pairs) Less(i, j int) bool {
//	return p[i].target < p[j].target
//}
//
//func findItinerary(tickets [][]string) []string {
//	result := []string{}
//	// map[出发机场] pair{目的地,是否被访问过}
//	targets := make(map[string]pairs)
//	for _, ticket := range tickets {
//		if targets[ticket[0]] == nil {
//			targets[ticket[0]] = make(pairs, 0)
//		}
//		targets[ticket[0]] = append(targets[ticket[0]], &pair{target: ticket[1], visited: false})
//	}
//	for k, _ := range targets {
//		sort.Sort(targets[k])
//	}
//	result = append(result, "JFK")
//	var backtracking func() bool
//	backtracking = func() bool {
//		if len(tickets)+1 == len(result) {
//			return true
//		}
//		// 取出起飞航班对应的目的地
//		for _, pair := range targets[result[len(result)-1]] {
//			if pair.visited == false {
//				result = append(result, pair.target)
//				pair.visited = true
//				if backtracking() {
//					return true
//				}
//				result = result[:len(result)-1]
//				pair.visited = false
//			}
//		}
//		return false
//	}
//
//	backtracking()
//
//	return result
//}
