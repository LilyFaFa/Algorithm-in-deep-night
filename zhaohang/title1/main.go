package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	a := 0
	for {
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			reader := bufio.NewReader(os.Stdin)
			in1, _, _ := reader.ReadLine()
			inputs1 := strings.Split(string(in1), " ")

			children := []int{}
			children = append(children, a)
			for _, n := range inputs1 {
				i, _ := strconv.Atoi(n)
				children = append(children, i)
			}

			in2, _, _ := reader.ReadLine()
			inputs2 := strings.Split(string(in2), " ")
			candy := []int{}
			for _, n := range inputs2 {
				i, _ := strconv.Atoi(n)
				candy = append(candy, i)
			}

			sort.Ints(children)
			sort.Ints(candy)

			result := 0
			i := 0
			j := 0

			for i < len(children) && j < len(candy) {
				if candy[j] >= children[i] {
					result++
					i++
					j++
				} else {
					j++
				}
			}
			fmt.Println(result)
		}
	}
}
