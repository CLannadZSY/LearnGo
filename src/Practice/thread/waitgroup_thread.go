// WaitGroup
// Go语言提供同步包（sync），源码（src/sync/waitgroup.go）。

// Sync包同步提供基本的同步原语，如互斥锁。除了Once和WaitGroup类型之外，大多数类型都是供低级库例程使用的。通过Channel和沟通可以更好地完成更高级别的同步。并且此包中的值在使用过后不要拷贝。

// Sync.WaitGroup是一种实现并发控制方式，WaitGroup 对象内部有一个计数器，最初从0开始，它有三个方法：Add(), Done(), Wait() 用来控制计数器的数量。

// Add(n) 把计数器设置为n 。
// Done() 每次把计数器-1 。
// wait() 会阻塞代码的运行，直到计数器地值减为0。
// 这种控制并发的方式适用于，好多个goroutine协同做一件事情的时候，因为每个goroutine做的都是这件事情的一部分，只有全部的goroutine都完成，这件事情才算是完成，这是等待的方式。WaitGroup相对于channel并发控制方式比较轻巧。

// 注意：
// 1、计数器不能为负值
// 2、WaitGroup对象不是一个引用类型
package main

import (
	"fmt"
	"sync"
	"time"
)

// 定义一个 WaitGroup
var wg sync.WaitGroup

func main() {
	// 计数器设置为2
	wg.Add(2)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutineA finish")
		//计数器减1
		wg.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("goroutineB finish")
		//计数器减1
		wg.Done()
	}()
	//会阻塞代码的运行，直到计数器地值减为0。
	wg.Wait()
	time.Sleep(2 * time.Second)
	fmt.Println("main fun exit")

}
