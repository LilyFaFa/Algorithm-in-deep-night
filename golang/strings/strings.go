package main

import (
	"fmt"
	"strings"
)

func TestStrings() {

	str := "AaBbCc"
	fmt.Println(str)

	// 1 将字符串的大写转化为小写
	str2 := strings.ToLower(str)
	fmt.Println("*** convert to lower ***")
	fmt.Println(str2)

	// 2 将字符串的小写转化为大写
	str3 := strings.ToUpper(str)
	fmt.Println("*** convert to upper ***")
	fmt.Println(str3)

	// 3 计算目标串中子串的数目
	count := strings.Count(str, "Aa")
	fmt.Println("*** count “Aa” ***")
	fmt.Println(count)

	// 4 计算字串出现的位置
	position := strings.Index(str, "Bb")
	fmt.Println("*** find position of “Bb” ***")
	fmt.Println(position)

	// 5 字符串比较
	fmt.Println("*** compare ***")
	fmt.Println(strings.Compare(str, str))
	fmt.Println(strings.Compare(str, str2))
	fmt.Println(strings.Compare(str, str3))

	// 6 字符串
	str = strings.Join([]string{"da", "zhu", "zai"}, ",")
	fmt.Println(str)
}

func main() {
	TestStrings()
}
