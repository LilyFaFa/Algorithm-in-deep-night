package main

import (
	"fmt"
	"strconv"
)

func main() {
	str := "123"
	i, _ := strconv.Atoi(str)
	fmt.Println(str, "字符串转整型", i)
	i64, _ := strconv.ParseInt(str, 10, 64)
	fmt.Println(str, "字符串转int64", i64)

	i = 123
	j := int64(123)
	s := strconv.Itoa(i)
	fmt.Println(i, "整型转字符串", s)
	s = strconv.FormatInt(j, 10)
	fmt.Println(j, "int64转字符串", s)

}
