package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	n := getBufferedIntInput(r)
	q := getBufferedIntInput(r)
	days := getInputToArray(n, r)

	dayMaps := make([]map[string]int, q)
	for i := 0; i < q; i++ {
		dayMaps[i] = getDayMap(r)
	}
	daysComulativeSum := makeComulativeSum(days)

	for _, v := range dayMaps{
		fmt.Fprintln(w, daysComulativeSum[v["right"]] - daysComulativeSum[v["left"] - 1])
	}
}

func getBufferedIntInput(r *bufio.Reader) int{
	var v int
	fmt.Scan(&v)
	return v
}

func getInputToArray(size int, r *bufio.Reader) []int{
	val := make([]int, size)
	for i := 0; i < size; i++{
		tmp := getBufferedIntInput(r)
		val[i] = tmp
	}
	return val
}

func getDayMap(r *bufio.Reader) map[string]int{
	daymap := map[string]int{
		"left": getBufferedIntInput(r),
		"right": getBufferedIntInput(r),
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
