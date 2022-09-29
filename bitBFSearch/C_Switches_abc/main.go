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

	for i:=0;i<1<<n;i++{
			sc := 0
		for j:=0;j<m;j++{
			tc := 0
			for l:=0;l<k[j];l++{
				if 1 == i&1<<s[j][l]{
					tc++
				}
			}
			if tc&1 == p[j]&1{
				sc++;
			}
		}
		if sc == m{
			ans++
		}
	}
	Println(ans)
}

