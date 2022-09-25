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

	for i, v := range n {
		Scan(&input)
		A = append(A, input)
	}
	Println(A)
}
