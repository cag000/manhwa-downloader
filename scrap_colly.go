package main

import (
	"fmt"
	_ "os"
	"strings"

	"github.com/gocolly/colly/v2"
)

func GetImage(url string)  []string{
	var linkStore = []string{}
	c := colly.NewCollector(
		colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36"),
		colly.AllowURLRevisit(),
	)
	
	c.OnRequest(func (r *colly.Request)  {
		fmt.Println("visit", r.URL)
	})

	c.OnHTML(`.chapter-img`, func (e *colly.HTMLElement)  {
		link := e.Attr("src")
		linkStore = append(linkStore, strings.TrimSpace(link))
	})

	c.Visit(url)
	return linkStore

}

