package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	fmt.Println(zhuankuai(input))
}

func zhuankuai(counts string) int {
	hashmap := make(map[rune]int)
	for _, s := range counts {
		_, ok := hashmap[s]
		if !ok {
			hashmap[s] = 1
			if len(hashmap) > 2 {
				return 0
			}
		}
	}
	return len(hashmap)
}
