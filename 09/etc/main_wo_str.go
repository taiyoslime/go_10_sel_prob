package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)
	pt := len(s)
	words := []string{"dream", "dreamer", "erase", "eraser"}
	for pt > 3 {
		_pt := pt
		for _, v := range words {
			if s[pt-3:pt] == v[len(v)-3:len(v)] && s[pt-len(v):pt] == v {
				pt -= len(v)
				break
			}
		}
		if _pt == pt {
			pt = -1
		}
	}
	if pt == 0 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
