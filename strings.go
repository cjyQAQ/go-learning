package main

import (
	"fmt"
	"strings"
)

/*
strings.Contains(s, substr string) bool  包含关系
strings.HasSuffix(s, suffix string) bool   后缀
strings.HasPrefix(s, suffix string) bool   前缀
strings.Index(s, str string) int   判断子字符串或字符在父字符串中出现的位置（索引）
strings.LastIndex(s, str string) int
strings.Replace(str, old, new string, n int) string   字符串替换
strings.Count(s, str string) int    某个字符出现的次数
strings.Repeat(s, count int) string   重复count次s串，生成一个新的字符串
strings.ToLower(s) string  大写转小写
strings.ToUpper(s) string 小写转大写
*/

func main() {
	var str string = "This is an example of a string"
	if strings.HasPrefix(str, "Th") { //前缀
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
	if strings.HasSuffix(str, "in") { //后缀
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	str2 := "hello world 12345"
	str3 := "12345"
	fmt.Println(strings.Replace(str2, str3, "1", 2)) // 将str2中的12345替换成1

}
