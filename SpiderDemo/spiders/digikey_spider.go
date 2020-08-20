package main

import (
	"SpiderDemo/save_data"
	"fmt"
	"github.com/antchfx/htmlquery"
	"github.com/gocolly/colly"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"strings"
	"time"
)

type DigikeyItem struct {
	kucName string
	kucGame string
}

var item []DigikeyItem

var c = colly.NewCollector()

func init() {
	c.Async = true

	c.Limit(&colly.LimitRule{
		// 并发数
		Parallelism: 3,
		// 随机延迟
		RandomDelay: 3 * time.Second,
	})

	// 设置代理, 会自己随机切换
	// 但是需要一直更新这个代理, 怎么处理呢? 在错误里面进行回调?
	//if p, err := proxy.RoundRobinProxySwitcher("http://183.143.161.201:4231"); err == nil {
	//	c.SetProxyFunc(p)
	//}

	//禁止 keep-live
	//c.WithTransport(&http.Transport{
	//	DisableKeepAlives: true,
	//})

	RedisSelf := save_data.ConnRedis("localhost", "./config_file/redis_config.yaml", "")
	TestDB := save_data.ConnMysql("test", "test", "./config_file/mysql_config.yaml", "")
	defer RedisSelf.Close()
	defer TestDB.Close()
}

func main() {
	t0 := time.Now()

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("Request: ", request.URL)
		//fmt.Println("proxy: ", request.ProxyURL)
	})

	c.OnResponse(func(response *colly.Response) {
		if response.StatusCode == http.StatusOK {
			htmlNodes, err := htmlquery.Parse(strings.NewReader(string(response.Body)))
			if err != nil {
				log.Fatal(err)
			}
			parseList(htmlNodes)

		} else {
			log.Fatalf("url: %s, code: %d", response.Request.URL, response.StatusCode)
		}
	})

	for i := 1; i <= 100; i++ {
		listURL := "https://www.digikey.com/products/en/cable-assemblies/flat-flex-ribbon-jumper-cables/463?FV=-8|463&quantity=0&ColumnSort=0&page=%d&pageSize=25"
		c.Visit(fmt.Sprintf(listURL, i))
	}
	c.Wait()
	fmt.Println(time.Since(t0))
	fmt.Printf("%v\n", item)

}

func parseList(html *html.Node) {

	nodes := htmlquery.Find(html, "//tbody[@id='lnkPart']/tr")
	for _, tr := range nodes {
		var digikeyItem DigikeyItem
		digikeyItem.kucName = htmlquery.FindOne(tr, "./td[@class='tr-mfgPartNumber']/a/span/text()").Data
		kucGame := htmlquery.FindOne(tr, "./td[@class='tr-dkPartNumber nowrap-culture']/a/text()").Data
		digikeyItem.kucGame = strings.TrimSpace(kucGame)
		item = append(item, digikeyItem)
	}
}
