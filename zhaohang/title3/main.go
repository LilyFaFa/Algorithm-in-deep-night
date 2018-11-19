package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var input int
	for {
		n, _ := fmt.Scan(&input)
		if n == 0 {
			break
		} else {
			result := 0
			for i := 1; i <= input; i++ {
				str := strconv.Itoa(i)
				if (strings.Contains(str, "2") || strings.Contains(str, "5") ||
					strings.Contains(str, "6") || strings.Contains(str, "9")) &&
					!strings.Contains(str, "3") && !strings.Contains(str, "4") && !strings.Contains(str, "7") {
					result++
				}

			}
			fmt.Println(result)
		}
	}

}
