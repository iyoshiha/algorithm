package main

import(
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	var input int
	var a []int
	for i := 0; i < n; i++{
		fmt.Scan(&input)
		a = append(a, input)
	}

	var dp []int = []int{0}
	for i := 0; i < n; i++{
		dp = append(dp, int(math.Max(float64(dp[i]), float64(dp[i]+a[i]))))
	} 

	fmt.Println(dp[n])	
}
