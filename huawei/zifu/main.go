package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()

	data2, _, _ := reader.ReadLine()
	var dream byte
	var flag bool
	if len(data2) != 0 {
		dream = data2[0]
		flag = true
	} else {
		dream = data[len(data)-1]
		flag = false
	}
	result := 0
	if 'a' <= dream && dream <= 'z' {
		for _, f := range data {
			if dream == f || dream == f-32 {
				result++
			}
		}
	} else if 'A' <= dream && dream <= 'Z' {
		for _, f := range data {
			if dream == f || dream == f+32 {
				result++
			}
		}
	} else {
		for _, f := range data {
			if dream == f {
				result++
			}
		}
	}
	if flag == false {
		result = result - 1
	}
	if result < 0 {
		result = 0
	}
	fmt.Println(result)
}
