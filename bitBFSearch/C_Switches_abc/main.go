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
			tmp--
			s[i] = append(s[i], tmp)
		}
	}
	for i:=0;i<m;i++;{
		Scan($tmp)
		p = append(p, tmp)
	}

	for i:=0;i<1<<n;i++{
		for j:=0;j<
		1&i>>s[j][k]

	}

}
