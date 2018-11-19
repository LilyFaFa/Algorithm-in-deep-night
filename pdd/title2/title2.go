package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scanf("%s", &input)
	result := 0
	head := input[0] == '0'
	tail := input[len(input)-1] == '0'

	for i := 1; i < len(input); i++ {
		tail1 := input[i-1] == '0'
		head1 := input[i] == '0'

		a := 0
		b := 0

		if head && tail1 {
			continue
		}

		if head1 && tail {
			continue
		}

		if head {
			a = 1
		} else {
			if input[i-1] == '0' {
				a = 1
			} else {
				a = i
			}
		}

		if tail {
			b = 1
		} else {
			if input[i] == '0' {
				b = 1
			} else {
				b = len(input) - i
			}
		}
		result += a * b
	}

	if head {
		if input[1] != '0' {
			if !tail {
				result += (len(input) - 1)
			} else {
				result += 1
			}
		} else {
			if !tail {
				result += 1
			}
		}
	}

	if tail {
		if input[len(input)-2] != '0' {
			if !head {
				result += (len(input) - 1)
			}
		} else {
			if !head {
				result += 1
			}
		}
	}

	fmt.Println(result)
}
