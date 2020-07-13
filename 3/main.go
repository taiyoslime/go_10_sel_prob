package main

import (
	"fmt"
)

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func div2(x int) (ans int) {
	for x%2 == 0 {
		x /= 2
		ans++
	}
	return
}

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	var ans int
	ans = 1e9
	for _, e := range arr {
		ans = min(ans, div2(e))
	}
	fmt.Println(ans)
}
