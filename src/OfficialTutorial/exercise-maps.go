/*
练习：映射
实现 WordCount。它应当返回一个映射，其中包含字符串 s 中每个“单词”的个数。
函数 wc.Test 会对此函数执行一系列测试用例，并输出成功还是失败。
你会发现 strings.Fields 很有帮助。
*/
package main

import (
	"fmt"
	"strings"
	//"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {

	m := make(map[string]int)
	str2 := strings.Fields(s)
	for _, v := range str2 {
		m[v]++
	}
	return m
}

func main() {
	s := "Hello World hello"
	ret := WordCount(s)
	fmt.Println(ret)
}
