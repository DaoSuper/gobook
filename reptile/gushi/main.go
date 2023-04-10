package main

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

var (
	baseUrl = "https://so.gushiwen.cn"
)

//	{
//	    "content": "轶事典故\n\n姓名由来\nxxx",
//	    "desc": "李白（701年－762年），字太白，号青莲居士，唐朝浪漫主义诗人，被后人誉为“诗仙”。xxx",
//	    "dynasty": "唐代",
//	    "id": 247,
//	    "name": "李白"
//	}
type Poet struct {
	Id      string
	Name    string
	Dynasty string
	Desc    string
}

func main() {
	getAuthors()
}

func getAuthors() {

	res, err := http.Get(baseUrl + "/authors")
	if err != nil {
		fmt.Println("Get error")
	}
	defer res.Body.Close()

	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("goquery error")
		// handler error
	}

	doc.Find(".sright a").Each(func(i int, s *goquery.Selection) {
		content := s.Text()
		href, _ := s.Attr("href")
		getPoet(href)
		fmt.Printf("%d: %s\n", i, content)
	})
}

func getPoet(url string) {
	res, err := http.Get(baseUrl + url)
	if err != nil {
		fmt.Println("poet get error")
	}
	defer res.Body.Close()
	fmt.Printf("%s\n", baseUrl)
	// use utfBody using goquery
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("poet goquery error")
		// handler error
	}

	doc.Find("#leftZhankai .sonspic").Find(".cont").Each(func(i int, s *goquery.Selection) {
		name := s.Find("p").Eq(0).Find("a:first-child b").Text()

		s.Find("p").Eq(1).Eq(0).Find("a").Each(func(i int, a *goquery.Selection) {
			hef, _ := a.Attr("href")
			fmt.Printf("herf: %s\n", hef)
		})
		desc := s.Find("p").Eq(1).Eq(0).Find("a").Remove().End().Text()
		fmt.Printf("%d: name: %s divimg: %s\n", i, name, desc)
	})
}
