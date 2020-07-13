package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n, a, b int
	fmt.Scan(&n, &a, &b)
	ans := 0
	s := func(x int) (sum int) {
		xStr := strconv.Itoa(x)
		for _, str := range xStr {
			ret, _ := strconv.Atoi(string(str))
			sum += ret
		}
		return
	}
	for i := 0; i <= n; i++ {
		if ret := s(i); a <= ret && ret <= b {
			ans += i
		}
	}
	fmt.Println(ans)
}
