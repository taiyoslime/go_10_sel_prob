package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	words := []string{"dream", "dreamer", "erase", "eraser"}
	for s != "" {
		op := false
		for _, v := range words {
			if strings.HasSuffix(s, v) {
				s = s[:len(s)-len(v)]
				op = true
			}
		}
		if !op {
			fmt.Println("NO")
			return
		}
	}
	fmt.Println("YES")
}
