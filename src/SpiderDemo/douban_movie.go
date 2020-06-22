package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"time"
)

type item struct {
	storyURL  string
	Source    string
	comments  string
	CrawledAt time.Time
	Comments  string
	Title     string
}

func main() {
	//stories := []item{}

	c := colly.NewCollector(
		//colly.AllowedDomains("*.douban.com"),
		// 异步??  还需要设置一下并发量
		colly.Async(true),
	)
	c.Limit(&colly.LimitRule{
		// 并发数
		Parallelism: 2,
		// 随机延迟
		RandomDelay: 5 * time.Second,
	})
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Request: ", request.URL)
	})

	c.OnResponse(func(response *colly.Response) {
		fmt.Println("OnResponse", response.StatusCode)
	})

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		movieInfo := element.Attr("href")
		fmt.Println(movieInfo)
	})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	startURL := "https://movie.douban.com/top250"
	c.Visit(startURL)

	c.Wait()
}
