package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result1 := 0
	result2 := 0
	result := 0
	for {

		var inputReader *bufio.Reader
		inputReader = bufio.NewReader(os.Stdin)
		code, _ := inputReader.ReadString('\n')

		if len(code) == 2 && code[0] == '}' {
			break
		} else {
			for i := 0; i < len(code); i++ {
				if code[0] != '/' {
					break
				}
				if len(code) > 1 && code[0] == '/' && code[1] == '/' {
					result1++
					break
				}

				if code[i] == '"' {
					for j := i + 1; j < len(code); j++ {
						if code[j] == '"' {
							i++
							break
						} else {
							i++
						}
					}
				} else if len(code) > i+1 && code[i] == '/' && code[i+1] == '*' {
					i++
					result2++
				} else if len(code) > i+1 && code[i] == '*' && code[i+1] == '/' {
					i++
					result2++
				}

			}
		}
		result = result1 + result2/2

	}

	fmt.Println(result)
}
