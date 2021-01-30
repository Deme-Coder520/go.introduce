package main

import "fmt"

func main() {

	// 原子计算器，只需要知道质子p、中子n和电子的e数目，计算出原子序数、质量和电荷
	// Z=p、A = p+n、c = p - e
	p, n, e := 0, 0, 0 //质子、中子和电子
	fmt.Println("原子计算器V1.0,输入质子、中子和电子的数量,得到原子序数、质量和电荷；")
	fmt.Println("说明：使用空格来分隔3个字符，回车进行计算，如16 16 18")
	fmt.Println("=============================")
	// Scanln遇到换行时才停止扫描
	fmt.Scanln(&p, &n, &e)
	Z := p     // 原子序数
	A := p + n // 原子质量
	c := p - e // 原子电荷
	fmt.Printf("当质子数=%d,中子数=%d,电子=%d时:\n", p, n, e)
	fmt.Printf("原子序数Z=%d,原子质量A=%d,电荷c=%d\n", Z, A, c)
}
