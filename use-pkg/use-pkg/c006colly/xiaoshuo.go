package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"strings"
)

// 将一本小说从头下载到尾
var superEvolutionUrl = "https://www.biqiuge.com/book/2753"

//流程：
//1.获取小说的首页，并解析出章节列表
//2.解析出每一章的url和章节名
//3.继续访问每一张的url

// 将最新的章节和全本保存在两个不同的文件夹中
func main() {
	//1.创建collector收集器
	c := colly.NewCollector()

	//2.设置gbk编码，可重复访问
	c.DetectCharset = true
	c.AllowURLRevisit = true

	//3.clone collector用于内容解析
	contentCollector := c.Clone() //拷贝
	beginRevist := false

	//4.div[class]筛选出Element为div并且有class这个属性的
	catalogSelector := "div[class=listmain]"
	c.OnHTML(catalogSelector, func(elemCatalog *colly.HTMLElement) {
		//5.筛选出dd元素下元素为a的内容
		elemCatalog.ForEach("dd>a", func(i int, elemHref *colly.HTMLElement) {
			tmpUrl := elemHref.Attr("href ")

			//6.1忽略前面的内容，从第一张开始
			if strings.Index(elemHref.Text, "第一章") != -1 {
				beginRevist = true
			}

			//6.2 拼装成全路径url
			if beginRevist {
				chapterUrl := elemHref.Request.AbsoluteURL(tmpUrl)
				//继续访问章节url
				contentCollector.Visit(chapterUrl)
			}
		})
	})

	c.OnRequest(func(request *colly.Request) {
		fmt.Println("visiting", request.URL.String())
	})

	//设置Onhtml回调函数
	contentSelector := "div[class=showtxt]"
	contentCollector.OnHTML(contentSelector, func(eleContent *colly.HTMLElement) {
		fmt.Printf("%s\n", eleContent.Text)
	})

	c.Visit(superEvolutionUrl)
}
