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
	for i:=0;i<m;i++{
		Scan($tmp)
		p = append(p, tmp)
	}

	for i:=0;i<1<<n;i++{
		for j:=0;j<m;j++{
			for l,totalSwitch:=0,0;l<k[j];l++{

				if 1 == 1&i>>s[j][l]{
					totalSwitch++
				}
			}
			if totalSwitch%2!=p[j]{
				notFlag = true
				break
			}
		}
		if !notFlag {
			ans++
		}
	}
	Println(ans)
}
