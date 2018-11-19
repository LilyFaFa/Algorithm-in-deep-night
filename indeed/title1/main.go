package main

import (
	"fmt"
	"io"
)

func main() {
	for {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err == io.EOF {
			break
		} else {
			var change int
			fmt.Scanf("%d", &change)
			var numbers []int
			result := 0
			for i := 0; i < number; i++ {
				var a int
				fmt.Scanf("%d", &a)
				if a%2 == 0 {
					result += a
				}
				numbers = append(numbers, a)
			}

			for i := 0; i < change; i++ {
				var m, n int
				fmt.Scanf("%d", &m)
				fmt.Scanf("%d", &n)

				t1 := false
				t2 := false

				if numbers[m-1]%2 == 0 {
					t1 = true
				}
				if n%2 == 0 {
					t2 = true
				}

				if t1 {
					if t2 {
						result = result + n
					} else {
						result = result - numbers[m-1]
					}
				} else {
					if !t2 {
						result = result + n
						result = result + numbers[m-1]
					}
				}
				numbers[m-1] = n + numbers[m-1]
				fmt.Println(result)
			}
		}
	}

}
