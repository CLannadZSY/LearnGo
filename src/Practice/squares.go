/*
拥有函数名的函数只能在包级语法块中被声明，通过函数字面量（function literal），我们可
绕过这一限制，在任何表达式中表示一个函数值。函数字面量的语法和函数声明相似，区别
在于func关键字后没有函数名。函数值字面量是一种表达式，它的值被成为匿名函数

*/
package main

import "fmt"

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func main() {
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"

}
