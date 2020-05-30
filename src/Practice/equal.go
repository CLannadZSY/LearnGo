/*
slice之间不能比较，因此我们不能使用==操作符来判断两个slice是否含有
全部相等元素。不过标准库提供了高度优化的bytes.Equal函数来判断两个字节型slice是否相
等（[]byte），但是对于其他类型的slice，


为何slice不直接支持比较运算符呢？
1. 一个slice的元素是间接引用的，
2. 一个slice甚至可以包含自身
*/
package main

import "fmt"

func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func main() {
	s := [...]string{"a", "a", "a", "a"}
	s1 := s[:2]
	s2 := s[2:]
	ret := equal(s1, s2)
	fmt.Println(ret)

	// slice唯一合法的比较操作是和nil比较，
	//if summer == nil { /* ... */ }
	var ss []int    // len(s) == 0, s == nil
	ss = nil        // len(s) == 0, s == nil
	ss = []int(nil) // len(s) == 0, s == nil
	ss = []int{}    // len(s) == 0, s != nil

}
