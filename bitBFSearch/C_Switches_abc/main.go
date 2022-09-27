package main

import(
	. "fmt"
	"math"
)

func main() {

	var n, m int
	var k []int
	var s [][]int
	var p []int
	var tmp int

	Scan(&n)
	Scan(&m)

	for i:=0;i<m;i++ {
		Scan(&tmp)
		k = append(k,tmp)
		for j:=0;k[i]<j;j++{
			Scan(&tmp)
			s[i] = append(s[i], tmp)
		}
	}
	for i:=0;i<m;i++;{
		Scan($tmp)
		p = append(p, tmp)
	}



	pattern:= int(math.Pow(2,float64(n)))
	for i := 0; i < pattern ; i++ {
		for j := 0; j < n; j++ {
		// Println(i>>j&1 == 1)
			if i>>j&1 == 1 {
				total+=A[j]
			}
		}
		for k := 0; k < q; k++ {
			if total==M[k] {
				ans = add(k, ans)
			}
		}
		total = 0
	}
	for _, v := range ans {
		if v == 1 { Println("yes")
		}else {
			Println("no")
		}
	}
}
