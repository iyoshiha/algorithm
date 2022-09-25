package main

import(
	. "fmt"
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


}
