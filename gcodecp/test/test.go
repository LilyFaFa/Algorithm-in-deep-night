package main

import (
	"bufio"
	//"fmt"
	"os"
)

func main() {
	buf := bufio.NewWriter(os.Stdout)
	//fmt.Println(buf.Available(), buf.Buffered()) // 4096 0

	buf.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	//fmt.Println(buf.Available(), buf.Buffered()) // 4070 26

	// 缓存后统一输出，避免终端频繁刷新，影响速度
	buf.Flush() // ABCDEFGHIJKLMNOPQRSTUVWXYZ

	res := "huhuhu"
	buf.WriteString(res)

}
