// 全局变量
// 这是并发控制最简单的实现方式
// 1、声明一个全局变量
// 2、所有子goroutine共享这个变量，并不断轮询这个变量检查是否有更新
// 3、在主进程中变更该全局变量
// 4、子goroutine检测到全局变量更新，执行相应的逻辑

// 优点：实现简单。
// 缺点：适用一些逻辑简单的场景，全局变量的信息量比较少，为了防止不同goroutine同时修改变量需要用到加锁来解决

// go run global_thread.go
package main

import (
	"fmt"
	"strconv"
	"time"
)

var open = true

func main() {
	for i := 0; i < 10; i++ {
		go OpenThread(strconv.Itoa(i), 1)
	}

	time.Sleep(2 * time.Second)
	open = false
	time.Sleep(2 * time.Second)
	fmt.Println("main func exit")
}

func OpenThread(ThreadName string, SleepTime time.Duration) {
	for open {
		fmt.Printf("Thread_%s running \n", ThreadName)
		time.Sleep(SleepTime * time.Second)
	}
	fmt.Printf("%s exit \n", ThreadName)
}
