package main

import(
	. "fmt"
	"math"
)

func setInput() (int, []int){
	var n int
		Scan(&n)
		var input int
		var A []int
		for i := 0; i < n ; i++{
			Scan(&input)
			A = append(A, input)
		}
		return n, A
}



func main() {

	n, A := setInput()
	q, M := setInput()
	Println(A, n, q, M)

	for pattern, k := 0,int(math.Pow(2,float64(n))); k < q; k++ {
		for i := 0; i < pattern ; i++ {
			for total, j := 0, 0; j < n; j++ {
				if pattern>>j&1 == 1 {
					total+=A[j]
				}
				if total==M[k] {
					Println("yes")
				} else {
					Println("no")
				}
			}
		}
	}
}
