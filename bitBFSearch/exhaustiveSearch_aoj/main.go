package main

import(
	. "fmt"
)

func main() {
	var n int 
	var A []int
//	var q int
//	var M []int
	var input int

	Scan(&n)

	// for i, v := range n {
	A = func (n int, A []int) []int{
		for i := 0; i < n ; i++{
			Scan(&input)
			A = append(A, input)
		}
		return A
	}(n, A)
		Println(A)


}
