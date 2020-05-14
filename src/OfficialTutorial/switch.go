package main

import (
	"fmt"
	//"runtime"
	"time"
)

//switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。
func main() {
	//fmt.Println("Go runs on")
	//switch os := runtime.GOOS; os {
	//case "darwin":
	//	fmt.Println("OS X.")
	//case "linux":
	//	fmt.Println("Linux")
	//case "windows":
	//	fmt.Printf("%s.\n", os)
	//default:
	//	fmt.Printf("%s.\n", os)
	//}
	main3()
}

func main2() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Tuesday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

//没有条件的 switch 同 switch true 一样。
func main3() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
