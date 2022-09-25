package main

import(
	. "fmt"
)

func setInput(n int) []int{
		var input int
		var A []int
		for i := 0; i < n ; i++{
			Scan(&input)
			A = append(A, input)
		}
		return A
}



func main() {
	var n int
	var A []int
	var q int
	var M []int

	Scan(&n)
	A = setInput(n)
	Scan(&q)
	M = setInput(q)

	Println(A)


}
