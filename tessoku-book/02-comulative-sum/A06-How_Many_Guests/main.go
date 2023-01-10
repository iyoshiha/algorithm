package main

import (
	"fmt"
)

func main() {
	var n, q int
	// A -> days
	var days// , l, r  []int
	fmt.Println("enter num: ")
	getInput(&n)
	getInput(&q)
	fmt.Println("the number you entered is ", n)
}

func getInput(v *int) {
	fmt.Scan(v)
}

func getInputToArray(v *[]int, size int) {



}
