package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"os"
	"regexp"
)

func main() {

	// 引数定義
	twid := flag.String("i", "", "-i TwitterID")
	flag.Parse()

	twUrl := "https://twitter.com/" + *twid + "/"

	picUrlList, flg := getPic(twUrl)

	if flg == false {
		os.Exit(1)
	}

	for _, u := range picUrlList {
		fmt.Println(u)
	}
}

// 与えられたTwitterIDで投稿されている画像をスクレイピングする
func getPic(tid string) ([]string, bool) {

	picUrls := []string{}

	doc, err := goquery.NewDocument(tid)
	if err != nil {
		return nil, false
	}

	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		imgUrl, _ := s.Attr("src")
		if ok, _ := regexp.MatchString("media", imgUrl); ok {
			picUrls = append(picUrls, imgUrl)
		}
	})

	return picUrls, true
}
