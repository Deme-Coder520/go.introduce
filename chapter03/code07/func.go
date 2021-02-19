package main

// ...表示一个变长函数
func sum(ns ...int) int {
	ret := 0
	for _, n := range ns {
		ret += n
	}
	return ret
}

func main() {
	println(sum())
	println(sum(1))
	println(sum(2, 3))
	println(sum(4, 5, 6))
}
