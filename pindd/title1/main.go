package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	var head string
	for {
		n, _ := fmt.Scan(&head)
		if n == 0 {
			break
		} else {
			reader := bufio.NewReader(os.Stdin)
			in1, _, _ := reader.ReadLine()
			splits := strings.Split(string(in1), " ")

			splits = append(splits, head)
			tmp := make(map[string]int)

			for _, str := range splits {
				str := strings.ToLower(str)
				if 'a' > str[0] || str[0] > 'z' {
					str = str[1:]
				}

				if len(str) == 0 {
					continue
				}

				if 'a' > str[len(str)-1] || str[len(str)-1] > 'z' {
					str = str[:len(str)-1]
				}

				if len(str) == 0 {
					continue
				}

				flag := false
				for i := 0; i < len(str); i++ {
					if 'a' > str[i] || str[i] > 'z' {
						flag = true
						break
					}
				}

				if flag {
					continue
				}
				if _, ok := tmp[str]; ok {
					tmp[str]++
				} else {
					tmp[str] = 1
				}
			}

			strMax := []string{}
			freMax := 0
			for key, value := range tmp {
				if value > freMax {
					freMax = value
					strMax = []string{key}
				} else if freMax == value {
					strMax = append(strMax, key)
				}
			}

			sort.Strings(strMax)
			for i := 0; i < len(strMax); i++ {
				fmt.Printf("%s ", strMax[i])
			}
			fmt.Println()
		}
	}
}
