package main

import (
	"fmt"
)

func main() {
	var number int
	for {
		n, _ := fmt.Scan(&number)
		if n == 0 {
			break
		} else {
			odds := []int{}
			evens := []int{}
			for i := 0; i < number; i++ {
				var a int
				fmt.Scan(&a)
				if a%2 == 0 {
					evens = append(evens, a)
				} else {
					odds = append(odds, a)
				}
			}
			for _, o := range odds {
				fmt.Printf("%d ", o)
			}
			for _, e := range evens {
				fmt.Printf("%d ", e)
			}
			fmt.Println("")

		}
	}
}





package main

import (
	"fmt"
)

func main() {
	var first int
	for {
		n, _ := fmt.Scan(&first)
		if n == 0 {
			break
		} else {
			var second int
			fmt.Scan(&second)

			firsts := []int{}
			seconds := []int{}

			for i := 0; i < first; i++ {
				var a int
				fmt.Scan(&a)
				firsts = append(firsts, a)
			}

			for i := 0; i < second; i++ {
				var a int
				fmt.Scan(&a)
				seconds = append(seconds, a)
			}

			i := 0
			j := 0
			result := []int{}

			for i < len(firsts) && j < len(seconds) {
				if firsts[i] < seconds[j] {
					result = append(result, firsts[i])
					i++
				} else {
					result = append(result, seconds[j])
					j++
				}
			}

			for a := j; a < len(seconds); a++ {
				result = append(result, seconds[a])
			}
			for a := i; a < len(firsts); a++ {
				result = append(result, firsts[a])
			}
			fmt.Printf("%d%s", result[0], " -> ")
			for a := 1; a < len(result)-1; a++ {
				if result[a] == result[a-1] {
					continue
				}
				fmt.Printf("%d%s", result[a], " -> ")
			}
			fmt.Println(result[len(result)-1])
		}
	}
}
