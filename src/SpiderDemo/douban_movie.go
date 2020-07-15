package main

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	"strings"
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

		// 384.9441ms
		t0 := time.Now()
		htmlNodes, err := htmlquery.Parse(strings.NewReader(string(response.Body)))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 1000; i++ {
			parseHtml(htmlNodes)
		}
		fmt.Println(time.Since(t0))

		//357.0443ms
		t1 := time.Now()
		doc, err := goquery.NewDocumentFromReader(bytes.NewReader(response.Body))
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < 1000; i++ {
			parseHtmlCss(doc)
		}
		fmt.Println(time.Since(t1))

	})

	//c.OnHTML("a[href]", func(element *colly.HTMLElement) {
	//	movieInfo := element.Attr("href")
	//	fmt.Println(movieInfo)
	//})

	c.OnError(func(response *colly.Response, err error) {
		fmt.Println(err)
	})

	startURL := "https://movie.douban.com/top250"
	c.Visit(startURL)

	c.Wait()
}

// 使用 htmlquery 模块的 xpath 语法进行解析数据
func parseHtml(html *html.Node) {

	nodes := htmlquery.Find(html, "//div[@class='info']")
	for _, node := range nodes {
		url := htmlquery.FindOne(node, ".//a/@href")
		title := htmlquery.FindOne(node, ".//span[@class='title']/text()")
		urlText := htmlquery.InnerText(url)
		titleText := htmlquery.InnerText(title)
		fmt.Println(urlText)
		fmt.Println(titleText)
	}
}

func parseHtmlCss(doc *goquery.Document) {

	doc.Find(".info").Each(func(i int, s *goquery.Selection) {
		url := s.Find("a").AttrOr("href", "")
		title := s.Find(".title").Text()
		fmt.Println(url)
		fmt.Println(title)
	})
}
