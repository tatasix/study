package bit

import "fmt"

// 定义状态
const (
	FlagUp           = 1 << iota // 0001
	FlagBroadcast                // 0010
	FlagLoopback                 // 0100
	FlagPointToPoint             // 1000
)

// 设置状态
func setFlag(flags, flag int) int {
	return flags | flag
}

// 清除状态
func clearFlag(flags, flag int) int {
	return flags &^ flag
}

// 检查状态
func hasFlag(flags, flag int) bool {
	return flags&flag != 0
}

func Handle() {
	// 初始化状态
	var flags int
	flags = setFlag(flags, FlagUp)
	fmt.Println(flags)

	flags = setFlag(flags, FlagBroadcast)
	fmt.Println(flags)

	// 检查状态
	fmt.Println(hasFlag(flags, FlagUp))       // true
	fmt.Println(hasFlag(flags, FlagLoopback)) // false
	fmt.Println(flags)

	// 清除状态
	flags = clearFlag(flags, FlagBroadcast)
	fmt.Println(hasFlag(flags, FlagBroadcast)) // false

	fmt.Println(flags)

	flags = setFlag(flags, FlagPointToPoint)
	fmt.Println(flags)
}
