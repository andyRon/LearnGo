package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

func main() {
	demo2()
}

func demo1() {
	c := colly.NewCollector(
		colly.AllowedDomains("www.baidu.com"),
	)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
		c.Visit(e.Request.AbsoluteURL(link))
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnResponse(func(r *colly.Response) {
		fmt.Printf("Response %s: %d bytes\n", r.Request.URL, len(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Printf("Error %s: %v\n", r.Request.URL, err)
	})

	c.Visit("http://www.baidu.com/")
}

// 这段代码会让爬虫从编程狮（W3Cschool.cn）网站的首页开始，抓取页面上的所有链接，并访问这些链接所指向的页面。MaxDepth 参数限制了爬虫的最大爬取深度，避免它陷入无限爬取的循环。
func demo2() {
	// 创建默认收集器
	c := colly.NewCollector()
	// 当访问页面时，输出页面标题
	c.OnHTML("title", func(e *colly.HTMLElement) {
		fmt.Println("网页标题：", e.Text)
	})
	// 当找到链接时，继续访问链接
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Println("发现链接：", link)
		e.Request.Visit(link)
	})
	// 限制最大深度，避免无限爬取
	c.MaxDepth = 2
	// 访问起始页面
	c.Visit("https://www.w3cschool.cn/")
}
