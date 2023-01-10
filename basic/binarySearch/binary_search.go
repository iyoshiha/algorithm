package main

import (
	"fmt"
)

func main() {
	var n int
	var S []int
	var q int
	var T []int
	var C int
	var mid int

	getInput := func(n *int, S *[]int) {
		var tmp int
		fmt.Scan(n)
		*S = make([]int, 0)
		for i := 0; i < *n; i++ {
			fmt.Scan(&tmp)
			*S = append(*S, tmp)
		}
	}
	getInput(&n, &S)
	getInput(&q, &T)


	for _, t := range T{
		for left, right := 0, n; left < right; {
			mid = (left+right)/2
			if t == S[mid]{
				C++
				break 
			} else if S[mid] < t {
				left = mid + 1

			} else if t < S[mid] {
				right = mid
			}
		}
	}
	fmt.Println(C)
}
