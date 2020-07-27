package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
)

func main() {
	url := "https://httpbin.org/delay/1"

	// Instantiate default collector
	c := colly.NewCollector()

	// create a request queue with 2 consumer threads
	q, _ := queue.New(
		2, // Number of consumer threads
		&queue.InMemoryQueueStorage{MaxSize: 10000}, // Use default queue storage
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
	})

	// 链接不能重复
	m, n := 0, 5
	for {
		fmt.Println("执行")
		for i := m; i < n; i++ {
			// Add URLs to the queue
			q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
		}
		// Consume URLs
		q.Run(c)
		fmt.Println("执行完成")
		m = n
		n += 10
	}

}
