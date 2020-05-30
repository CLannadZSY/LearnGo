/*
go 语言圣经-中文版
4.2.1 章节 append 函数
*/
package main

import (
	"fmt"
)

const HelloWorld = "Hello World"

func appendExample() []rune {
	var runes []rune
	for _, r := range HelloWorld {
		runes = append(runes, r)
	}
	fmt.Printf("%q \n", runes)
	return runes
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

func main() {
	appendExample()

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d %v\n", i, cap(y), y)
		x = y
	}

}
