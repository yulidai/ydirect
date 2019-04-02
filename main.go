package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/PuerkitoBio/goquery"
	"github.com/fatih/color"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("missing parameters")
		return
	}

	url := fmt.Sprintf("http://dict.youdao.com/w/%s", os.Args[1])

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
		color.Green(s.Text())
		//fmt.Printf("%s\n", s.Text())
	})
}