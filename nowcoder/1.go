package nowcoder

import (
	"fmt"
	"sync/atomic"
)

func DoTest() {
	a := int64(0)
	JudgeStr("nowcoder", a)
	JudgeStr("NowCoder", a)
	JudgeStr("haha", a)
	fmt.Println(a)
}

func JudgeStr(s string, a int64) {
	switch s {
	case "nowcoder", "haha":
		atomic.AddInt64(&a, 1)
	default:
		atomic.AddInt64(&a, -1)
	}
	fmt.Printf("%s:%d \n", s, a)
}
