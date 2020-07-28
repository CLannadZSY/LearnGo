// channel
// channel 是 goroutine 之间只要的通讯方式, 一般会和 select 搭配使用
// 如想了解channel实现原理可参考
// https://github.com/guyan0319/golang_development_notes/blob/master/zh/9.9.md
// 1. 声明一个 stop 的 chan
// 2. 在 goroutine 中, 使用 select 判断 stop 是否可以接收到值, 如果可以, 就表示可以退出了
//    如果没有接收到, 就会执行 default 中的逻辑, 直到收到
// 3. 主程序发送 stop <- true 结束的指令后
// 4. 子 goroutine 接收到结束指令 case <- stop 退出  return

// 这种select+chan是一种比较优雅的并发控制方式，但也有局限性，如多个goroutine 需要结束，以及嵌套goroutine 的场景。

package main

import (
	"fmt"
	"time"
)

var stop = make(chan bool)

func main() {
	go OpenThread()
	time.Sleep(2 * time.Second)
	stop <- true
	time.Sleep(2 * time.Second)
	fmt.Println("main fun exit")
}

func OpenThread() {
	for {
		select {
		case <-stop:
			fmt.Println("goroutine exit")
			return
		default:
			fmt.Println("goroutine running")
			time.Sleep(1 * time.Second)
		}
	}
}
