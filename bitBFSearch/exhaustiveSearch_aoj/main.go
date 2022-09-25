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

	pattern := math.Pow(2,n);
	for i := 0; i < pattern ; i++ {
		for j := 0; j < n; j++ {
			pattern >> j & 1
		}
	}
}
