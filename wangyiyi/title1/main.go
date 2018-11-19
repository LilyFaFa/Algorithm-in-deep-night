package main

import (
	"fmt"
)

func main() {
	a := 0
	for {
		result := 0
		n, _ := fmt.Scan(&a)
		if n == 0 {
			break
		} else {
			result += a / 5
			a := a % 5
			if a == 0 {
				fmt.Println(result)
				continue
			}

			result += a / 4
			a := a % 4
			if a == 0 {
				fmt.Println(result)
				continue
			}

			result += a / 3
			a := a % 3
			if a == 0 {
				fmt.Println(result)
				continue
			}

			result += a / 2
			a := a % 2
			if a == 0 {
				fmt.Println(result)
				continue
			}

			result += a / 1
			a := a % 1
			if a == 0 {
				fmt.Println(result)
				continue
			}

		}
	}
}


package main

import (
	"fmt"
)

func main() {
	var input string
	for {
		n, _ := fmt.Scan(&input)
		if n == 0 {
			break
		} else {
			if len(input) == 1 || len(input) == 0 {
				fmt.Println(len(input))
				continue
			}
			dp := make([]int, len(input))
			
			if input[0] == input[len(input)-1] {
				dp[0] = 1
				for i := 1; i < len(input); i++ {
					if input[i] != input[i-1] {
						dp[i] = dp[i-1] + 1
					} else {
						dp[i] = 0
						flag=true
					}
				}
			}

			result := 0
			for i := 0; i < len(dp); i++ {
				if dp[i] > result {
					result = dp[i]
				}
			}
			
			fmt.Println(result)
		}
	}
}
