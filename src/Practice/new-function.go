/*
用new创建变量和普通变量声明语句方式创建变量没有什么区别，除了不需要声明一个临时变
量的名字外，我们还可以在表达式中使用new(T)。换言之，new函数类似是一种语法糖，而
不是一个新的基础概念。

*/
package main

import "fmt"

func newInt() *int {
	return new(int)
}

func newInt2() *int {
	var dummy int
	return &dummy
}

func main() {
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)

	fmt.Println(newInt())
	fmt.Println(newInt2())
}
