package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	sort.Slice(arr, func(i, j int) bool {
		if arr[i] > arr[j] {
			return true
		}
		return false
	})
	alice, bob := 0, 0
	for i, v := range arr {
		if i%2 == 0 {
			alice += v
		} else {
			bob += v
		}
	}
	fmt.Println(alice - bob)
}
