package main

import (
	"bufio"
	"fmt"
	"os"
)

//exit自动退出程序
func main() {
	fmt.Println("Exit自动退出程序")
	f := bufio.NewReader(os.Stdin)
	input := ""
	str := ""
	for {
		fmt.Print(">")
		input, _ = f.ReadString('\n')
		if len(input) == 1 { //空行继续输入
			continue
		}
		// fmt.Sscan可以移除换行换行符
		fmt.Sscan(input, &str)
		if str == "exit" {
			break
		} else {
			fmt.Printf("输入%s\n", str)
		}
	}
}
