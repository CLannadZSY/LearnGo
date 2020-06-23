package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

func main() {

	fileName := "quote.md"
	file, errFile := os.Create(fileName)
	if errFile != nil {
		println("operating system create file error :%s", errFile.Error())
		panic(errFile)
	}
	defer func() {
		err := file.Close()
		if err != nil {
			println("file close error")
		}
	}()

	c := colly.NewCollector()
	errProxy := c.SetProxy("http://127.0.0.1:23333/")
	if errProxy != nil {
		println("colly set proxy error :%s", errProxy.Error())
		panic(errProxy)
	}
	// c.AllowedDomains  = []string{"https://www.goodreads.com"}
	c.AllowURLRevisit = true
	extensions.RandomUserAgent(c)

	c.OnHTML(".quoteText ",
		func(e *colly.HTMLElement) {
			text := strings.TrimSpace(strings.Split(e.Text, "―")[0])
			author := TrimSpaceNewlineInString(strings.TrimSpace(e.ChildText(".authorOrTitle")))

			fileWriteForMarkdown(file, text, author)
		})

	errVisit := c.Visit("https://www.goodreads.com/quotes/tag/philosophy")
	if errVisit != nil {
		panic(errVisit)
	}

}

func TrimSpaceNewlineInString(s string) string {
	re := regexp.MustCompile(`\n`)
	return re.ReplaceAllString(s, " ")
}

func fileWriteForMarkdown(file *os.File, lines ...string) {
	var admotionBot = `
\{\{% /admonition %\}\}
`
	head := fmt.Sprintf(`
\{\{%% admonition quote "%s" %%\}\}
`, lines[1])
	_, err := (*file).Write([]byte(head))
	if err != nil {
		println("file write error ", err.Error())
	}
	_, err = (*file).Write([]byte(lines[0]))
	if err != nil {
		println("file write error ", err.Error())
	}
	_, err = (*file).Write([]byte(admotionBot))
	if err != nil {
		println("file write error ", err.Error())
	}
}
