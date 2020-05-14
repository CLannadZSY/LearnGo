package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	main2()
}

//for（续）
//初始化语句和后置语句是可选的。
func main2() {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	fmt.Println(sum)
}
