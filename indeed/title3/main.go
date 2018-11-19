package main

import (
	"fmt"
	"io"
	"math"
)

var p int

func main() {
	for {
		var number int
		_, err := fmt.Scanf("%d", &number)
		if err == io.EOF {
			break
		} else {
			p = int(math.Pow(10, 9))
			p = p + 7

			var column int
			fmt.Scanf("%d", &column)

			var inputs [][]byte
			for i := 0; i < number; i++ {
				var s string
				fmt.Scanf("%s", &s)
				input := []byte(s)
				inputs = append(inputs, input)

			}

			var casenum int
			fmt.Scanf("%d", &casenum)
			for i := 0; i < casenum; i++ {
				var m, n int
				fmt.Scanf("%d", &m)
				fmt.Scanf("%d", &n)
				a := 0
				b := 0
				searchBottom(inputs, &a, m-1, n-1)
				searchBegin(inputs, &b, m-1, n-1)
				result := a * b
				result = result % p

				fmt.Println(result)
			}

		}
	}
}

func searchBottom(inputs [][]byte, result *int, curx int, cury int) {
	if curx == len(inputs)-1 && cury == len(inputs[0])-1 && inputs[curx][cury] == '.' {
		*result = *result % p
		*result = *result + 1
		return
	}

	if curx >= len(inputs) || cury >= len(inputs[0]) {
		return
	}

	if inputs[curx][cury] == '.' {
		searchBottom(inputs, result, curx+1, cury)
		searchBottom(inputs, result, curx, cury+1)
	}

}

func searchBegin(inputs [][]byte, result *int, curx int, cury int) {
	if curx == 0 && cury == 0 && inputs[curx][cury] == '.' {
		*result = *result % p
		*result = *result + 1
		return
	}

	if curx < 0 || cury < 0 {
		return
	}

	if inputs[curx][cury] == '.' {
		searchBegin(inputs, result, curx-1, cury)
		searchBegin(inputs, result, curx, cury-1)
	}
}
