package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var point int
		_, err := fmt.Scanf("%s=d", &point)
		if err == io.EOF {
			break
		} else {

		}
	}

}
