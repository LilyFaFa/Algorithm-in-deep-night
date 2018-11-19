package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var result []string
	searchResult(&result, s, 0, 0, "", 0)
	return result
}

func searchResult(result *([]string), s string, cur int, curi int, tmp string, i int) {
	if len(s) == i {
		if cur <= 255 && curi == 3 {
			*result = append(*result, tmp)
		}
		return
	}

	if curi >= 4 {
		return
	}

	if cur > 255 {
		return
	}

	a, _ := strconv.Atoi(string(s[i]))
	if len(tmp) == 0 || string(tmp[len(tmp)-1]) == "." {
		searchResult(result, s, cur*10+a, curi, tmp+string(s[i]), i+1)
	} else if string(tmp[len(tmp)-1]) != "." && cur == 0 {
		searchResult(result, s, 0, curi+1, tmp+".", i)
	} else {
		searchResult(result, s, 0, curi+1, tmp+".", i)
		searchResult(result, s, cur*10+a, curi, tmp+string(s[i]), i+1)
	}

}

func main() {
	result := restoreIpAddresses("25525511135")
	fmt.Println(result)
}
