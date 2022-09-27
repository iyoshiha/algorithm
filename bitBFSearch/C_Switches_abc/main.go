package main

import(
	. "fmt"
	"math"
)

func main() {

	var n, m int
	Scan()
	var ans []int = make([]int, q)
	var total int = 0

	pattern:= int(math.Pow(2,float64(n)))
	for i := 0; i < pattern ; i++ {
		for j := 0; j < n; j++ {
		// Println(i>>j&1 == 1)
			if i>>j&1 == 1 {
				total+=A[j]
			}
		}
		for k := 0; k < q; k++ {
			if total==M[k] {
				ans = add(k, ans)
			}
		}
		total = 0
	}
	for _, v := range ans {
		if v == 1 { Println("yes")
		}else {
			Println("no")
		}
	}
}
