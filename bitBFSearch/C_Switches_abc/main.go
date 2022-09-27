package main

import(
	. "fmt"
)

func main() {

	var n, m int
	var k []int
	var s [][]int
	var p []int
	var tmp int
	var ans int
	var notFlag bool

	Scan(&n)
	Scan(&m)
	s = make([][]int, m)
	for i:=0;i<m;i++{
		s[i] = make([]int, n)
	}

	for i:=0;i<m;i++ {
		Scan(&tmp)
		k = append(k,tmp)
		for j:=0;j<k[i];j++{
			Scan(&tmp)
			tmp--
			s[i][j] = tmp
		}
	}
	for i:=0;i<m;i++{
		Scan(&tmp)
		p = append(p, tmp)
	}
	Println(n,m,k,s,p)

	for i:=0;i<1<<n;i++{
		for j,totalSwitch:=0,0;j<m;j++{
			for l,totalSwitch:=0,0;l<k[j];l++{

				if 1 == 1&i>>s[j][l]{
					totalSwitch++
				}
			}
			if totalSwitch%2!=p[j]{
				notFlag = true
				break
			}
			totalSwitch=0
		}
		if !notFlag {
			ans++
		}
		notFlag=false
	}
	Println(ans)
}
