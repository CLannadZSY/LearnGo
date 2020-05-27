/*
练习：斐波纳契闭包
让我们用函数做些好玩的事情。

实现一个 fibonacci 函数，它返回一个函数（闭包），该闭包返回一个斐波纳契数列 `(0, 1, 1, 2, 3, 5, ...)`。
*/
package main

import "fmt"

// 返回一个“返回int的函数”
func fibonacci() func() int {

	back1, back2 := 0, 1 // 预先定义好前两个值

	return func() int {

		//记录（back1）的值
		temp := back1

		// 重新赋值
		back1, back2 = back2, back1+back2

		return temp
	}

}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
