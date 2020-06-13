package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	q := "goquery"

	url := "https://godoc.org/?q=" + q

	// 发起http请求获取响应并创建Document结构体指针
	document, err := goquery.NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	// 在docuemnt中通过选择器去查找元素
	// <tagname>
	// 标签选择器
	// 获取所有的a标签
	selection := document.Find("a")
	selection.Each(func(index int, tag *goquery.Selection) {
		href, exists := tag.Attr("href")
		// tag.Html()

		fmt.Println(tag.Text(), href, exists)
	})

	fmt.Println("==============class==============")
	// class选择器
	// .className
	// table
	// 在table下获取所有的超链接
	document.Find(".table-condensed").Find("a").Each(func(index int, tag *goquery.Selection) {
		href, exists := tag.Attr("href")
		// tag.Html()

		fmt.Println(tag.Text(), href, exists)
	})

	// id选择器
	// #id
	fmt.Println(document.Find("#x-search").Attr("class"))
	fmt.Println(document.Find("#x-search").Html())
	fmt.Println(document.Find("#x-search").Text())

	// 符合选择器
	// tag + class
	// <div></div><div class="nav"></div><span class="nav"></span>
	// tag.class

	// 子孙选择器
	// selector1 selector2 selector3 ...
	fmt.Println("=========子孙选择器============")
	document.Find(".table-condensed a").Each(func(index int, tag *goquery.Selection) {
		href, exists := tag.Attr("href")
		// tag.Html()

		fmt.Println(tag.Text(), href, exists)
	})

	// 子选择器
	// selector1 > selector2
	// document.Find(selector1).ChildrenFiltered(selector2)
}
