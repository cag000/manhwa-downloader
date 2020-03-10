package main

import (
	"fmt"
	"log"
	"path/filepath"
)

func main()  {
	// GetLinkImage("https://manhwa18.com/read-a-wonderful-new-world-raw-chapter-1.html") //chromedp	
	// fmt.Println(storeMe[len(storeMe)-1]) //print last index
	title := "youre-not-that-special"
	linkUrl := "https://manhwa18.com/read-youre-not-that-special-chapter-"
	chap := 50
	for i := 36; i <=chap; i++ {
		linkMe := fmt.Sprintf("%s%d.html", linkUrl, i)
		storeMe := GetImage(linkMe)
		chapName := fmt.Sprintf("chap_%d", i)
		_, err := CreateDirIfNotExist(filepath.Join("folder_manhwa", title, chapName))
		if err != nil {
			log.Fatal(err)
		}
		_, serr := DownloadImage(title, chapName, title, storeMe)
		if serr != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%s || COMPLETE\n", title)
}

func Manhwa18()  {
	title := "youre-not-that-special"
	linkUrl := "https://manhwa18.com/read-youre-not-that-special-chapter-"
	chap := 50
	for i := 36; i <=chap; i++ {
		linkMe := fmt.Sprintf("%s%d.html", linkUrl, i)
		storeMe := GetImage(linkMe)
		chapName := fmt.Sprintf("chap_%d", i)
		_, err := CreateDirIfNotExist(filepath.Join("folder_manhwa", title, chapName))
		if err != nil {
			log.Fatal(err)
		}
		_, serr := DownloadImage(title, chapName, title, storeMe)
		if serr != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%s || COMPLETE\n", title)
}

func Toonily() {
	title := "youre-not-that-special"
	linkUrl := "https://toonily.com/webtoon/not-that-special/chapter-"
	chap := 50
	for i := 36; i <=chap; i++ {
		linkMe := fmt.Sprintf("%s%d", linkUrl, i)
		storeMe := GetImage(linkMe)
		chapName := fmt.Sprintf("chap_%d", i)
		_, err := CreateDirIfNotExist(filepath.Join("folder_manhwa", title, chapName))
		if err != nil {
			log.Fatal(err)
		}
		_, serr := DownloadImage(title, chapName, title, storeMe)
		if serr != nil {
			log.Fatal(err)
		}
	}
	log.Printf("%s || COMPLETE\n", title)
}