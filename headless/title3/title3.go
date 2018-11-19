package main

import (
	"fmt"
)

func main() {

	var m, n int
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &n)
	var input []string
	var dic []string
	var result []int
	var tmp string
	for i := 0; i < m; i++ {
		var instr string
		fmt.Scanf("%s", &instr)
		dic = append(dic, instr)
	}

	fmt.Scanf("%s", &tmp)

	for i := 0; i < n; i++ {
		var instr string
		fmt.Scanf("%s", &instr)
		input = append(input, instr)
	}

	for i := 0; i < len(input); i++ {
		item := input[i]
		flag := 0
		for j := 0; j < len(dic); j++ {
			dicitem := dic[j]
			k := 0
			for k = 0; k < len(dicitem); k++ {
				if item[k] != dicitem[k] {
					break
				}

			}
			if k == len(dicitem) {
				flag = 1
				break
			}
		}

		if flag == 1 {
			result = append(result, 1)
		} else {
			result = append(result, -1)
		}

	}

	for i := 0; i < len(result); i++ {
		fmt.Println(result[i])
	}

}
