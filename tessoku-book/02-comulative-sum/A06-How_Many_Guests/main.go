package main

import (
	"fmt"
)

func main() {
	// , l, r  []int
	
	fmt.Println("enter num: ")
	n := getInput()
	q := getInput()
	days := getInputToArray(n + 1)
	fmt.Println("the numbers you got from stdin are", days)
	fmt.Println(n)
	fmt.Println(q)
	
	dayMaps := make([]map[string]int, q)
	for i := 0; i < q; i++ {
		dayMaps[i] = getDayMap()
	}
	fmt.Println("this is dayMaps", dayMaps)
}

func getInput() int{
	var v int
	fmt.Scan(&v)
	return v
}

func getInputToArray(size int) []int{
	val := make([]int, size)
	for i := 1; i < size; i++{
		tmp := getInput()
		val[i] = tmp
	}
	return val
}

func getDayMap() map[string]int{
	daymap := map[string]int{
		"left": getInput(),
		"right": getInput(),
	}
	return daymap
}
