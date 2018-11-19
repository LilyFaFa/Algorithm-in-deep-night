package main

import (
	"fmt"
	//	"io"
)

func main() {

	var f float64
	fmt.Scanf("%f", &f)

	i := int(f)
	f2 := float64(i) + 0.5
	if f >= f2 {
		fmt.Println(i + 1)
	} else {
		fmt.Println(i)
	}
}
