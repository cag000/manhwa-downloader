package main

import (
	"context"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
)

func main() {
	GetLinkImage("https://toonily.com/webtoon/not-that-special/chapter-36/") //chromedp
}

func GetLinkImage(url string) {
	//create context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	var link string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.InnerHTML("div.reading-content", &link, chromedp.ByQuery),
	)
	if err != nil {
		log.Fatal(err)
	}

	// log.Print(link)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(link))
	if err != nil {
		log.Fatal(err)
	}

	var linkStore []string
	// log.Print(doc)
	doc.Find(".page-break no-gaps").Each(func(i int, s *goquery.Selection)  {
		// log.Print(s)
		linkUrl, _ := s.Find(".wp-manga-chapter-img img-responsive lazyload effect-fade").Attr("data-src")
		linkStore = append(linkStore, linkUrl)
	})

	log.Print(doc.Find("wp-manga-chapter-img img-responsive lazyload effect-fade"))

	log.Print(linkStore)
}
