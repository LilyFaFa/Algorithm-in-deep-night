package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var point int
		_, err := fmt.Scanf("%d", &point)
		if err == io.EOF {
			break
		} else {
			if point == 1 {
				fmt.Printf("%d %d", 0, 0)
			}
			if point == 2 {
				fmt.Printf("%d %d", 1, 0)
			}
			if point == 3 {
				fmt.Printf("%d %d", 3, 1)
			}
			if point > 3 {
				fmt.Printf("%d %d", 3*(point-2), (point-1)*(point-2)/2)
			}
		}
	}

}
