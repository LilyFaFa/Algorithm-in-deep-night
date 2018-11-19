package main

import (
	"fmt"
	"io"
)

func main() {
	day := 0
	for {
		_, err := fmt.Scanf("%d", &day)
		if err == io.EOF {
			break
		} else {
			result := stock(day)
			fmt.Printf("%d\n", result)
		}
	}
}

func stock(day int) int {
	d := 1
	circle := 1
	result := day
	for d <= day {
		if d == ((1+circle)*circle)/2+circle+1 {
			result = result - 2
			circle++
		}
		d++

	}
	return result
}
