package main

import (
	"fmt"
	"os"
	"net/http"
	//"io/ioutil"
	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing parameters")
		return
	}

	url := fmt.Sprintf("http://dict.youdao.com/w/%s", os.Args[1])
	//fmt.Printf("查询单词: %s\n", url)

	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Printf("error: %s\n", err)
		return
	}

	doc.Find("#phrsListTab .trans-container > ul > li").Each(func(i int, s *goquery.Selection) {
		fmt.Printf("%s\n", s.Text())
	})
}