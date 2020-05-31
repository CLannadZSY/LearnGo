package main

import "fmt"

func main() {
	//ages := make(map[string]int)
	//ages["bob"]++
	//fmt.Println(ages)

	fmt.Println(EqualMap(map[string]int{"A": 0}, map[string]int{"B": 42}))
	fmt.Println(EqualMap(map[string]int{"A": 0}, map[string]int{"A": 0}))

}

func EqualMap(x, y map[string]int) bool {
	//!ok来区分元素不存在，与元素存在但为0的。我们不能简单地用
	//xv != y[k]判断，那样会导致在判断下面两个map时产生错误的结果：
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true

}
