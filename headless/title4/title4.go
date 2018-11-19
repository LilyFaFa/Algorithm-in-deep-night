package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var h int
		_, err := fmt.Scanf("%d", &h)
		fmt.Println(h)
		if err == io.EOF {
			fmt.Println("EOF")
			return
		}

	}
}
