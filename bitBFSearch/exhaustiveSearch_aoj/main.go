package main

import(
	. "fmt"
)

func main() {
	var n byte
	var A []int
	var q byte
	var M []int
	var input

	Scan($n)

	for i, v := range n {
		Scan($input)
		A = append(A, input)
	}
	Println(A)
}
