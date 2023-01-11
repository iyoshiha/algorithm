package main

import (
	"fmt"
)

func main() {
	n := getInput()
	q := getInput()
	days := getInputToArray(n)
	dayMaps := make([]map[string]int, q)
	for i := 0; i < q; i++ {
		dayMaps[i] = getDayMap()
	}
	daysComulativeSum := makeComulativeSum(days)

	for _, v := range dayMaps{
		fmt.Println(daysComulativeSum[v["right"]] - daysComulativeSum[v["left"] - 1])
	}

	// check input 
	fmt.Println("this is dayMaps", dayMaps)
	fmt.Println("the numbers you got from stdin are", days)
	fmt.Println(n)
	fmt.Println(q)
	fmt.Println("this is comulative sum", daysComulativeSum)

}

func getInput() int{
	var v int
	fmt.Scan(&v)
	return v
}

func getInputToArray(size int) []int{
	val := make([]int, size)
	for i := 0; i < size; i++{
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

func makeComulativeSum(days []int) []int{
	comulativeSum := make([]int, len(days) + 1)
	for i:=0; i < len(days); i++{
		comulativeSum[i+1] = comulativeSum[i] + days[i]
	}
	return comulativeSum
}
