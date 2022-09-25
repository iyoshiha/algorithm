package main

import(
	. "fmt"
)

func an(n int, A []int) []int{
		var input int
		for i := 0; i < n ; i++{
			Scan(&input)
			A = append(A, input)
		}
		return A
}



func main() {
	var n int 
	var A []int
//	var q int
//	var M []int

	Scan(&n)

	// for i, v := range n {
	var bn func(int, []int)[]int = an
	Println(an)
	Println(bn)
	Println(A)


}
