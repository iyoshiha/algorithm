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

	d := getBufferedIntInput(r)
	n := getBufferedIntInput(r)

	maps := make([]map[string]int, n)
	for i := 0; i < n; i++ {
		maps[i] = getMap(r)
	}
	comulativeSum := makeComulativeSum(maps, d)

	for _, v := range comulativeSum{
		fmt.Fprintln(w, v)
	}
}

func getBufferedIntInput(r *bufio.Reader) int{
	var v int
	fmt.Fscan(r, &v)
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

func getMap(r *bufio.Reader) map[string]int{
	personMap := map[string]int{
		"left": getBufferedIntInput(r),
		"right": getBufferedIntInput(r),
	}
	return personMap
}

func makeComulativeSum(maps []map[string]int, d int) []int{
	peopleNumForEventdays := make([]int, d)
	comulativeSum := make([]int, d)
	for i:=0; i < len(maps); i++{
		// -1は、eventの日が1日目の場合、対応する配列のインデックスは0のためminusする
		peopleNumForEventdays[maps[i]["left"] - 1]++
		// rightの場合,参加した次の日が前日比が-1になる。
		// よって、５日目まで参加した場合、6日目に-1。
		// 対応する日は、日付-1=インデックス
		// 6-1=5 index 5 が-1になる
		peopleNumForEventdays[maps[i]["right"]]--
	}
	for i, v:=range peopleNumForEventdays{
		if i == 0{
			comulativeSum[0] = v
			continue
		}
		comulativeSum[i] = v + comulativeSum[i-1]
	}
	return comulativeSum
}
