package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	s := strings.Split(string(data), " ")
	fmt.Println(len(s[len(s)-1]))
}
