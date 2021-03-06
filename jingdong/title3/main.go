package main

import (
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func main() {
	for {
		var inputs string
		_, err := fmt.Scanf("%s", &inputs)
		if err == io.EOF {
			break
		} else {
			hh := make(map[string]bool)
			ss := strings.Split(inputs, ",")
			if len(ss) == 1 {
				fmt.Println(ss[0])
				continue
			}

			numbers := []int{}
			for _, s := range ss {
				if _, ok := hh[s]; ok {
					continue
				}
				hh[s] = true
				num, _ := strconv.Atoi(s)
				numbers = append(numbers, num)
			}

			sort.Ints(numbers)
			fmt.Print(numbers[0])
			for i := 1; i < len(numbers); i++ {
				if numbers[i] == numbers[i-1]+1 {
					continue
				} else if i > 1 && numbers[i-1] == numbers[i-2]+1 {
					fmt.Print("-")
					fmt.Print(numbers[i-1])
				}
				fmt.Print(",")
				fmt.Print(numbers[i])
			}
			if numbers[len(numbers)-1] == numbers[len(numbers)-2]+1 {
				fmt.Print("-")
				fmt.Print(numbers[len(numbers)-1])
			}
			fmt.Println("")
		}
	}
}
