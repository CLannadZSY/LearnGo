/*
函数值
函数也是值。它们可以像其它值一样传递。

函数值可以用作函数的参数或返回值。
*/
package main

import (
	"fmt"
	"math"
)

func computer(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func TriangleTheorem(x, y float64) float64 {
	/*
		勾股定理
	*/
	return math.Sqrt(x*x + y*y)
}

func main() {
	hypot := TriangleTheorem
	fmt.Println(hypot(5, 12))
	//fmt.Println(hypot)

	fmt.Println(computer(hypot))

	// n ^ n 次方
	fmt.Println(computer(math.Pow))
}
