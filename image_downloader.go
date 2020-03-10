package main

import (
	_"archive/zip"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var netClient = &http.Client{}

func DownloadImage(title string, chapName string, fdName string, link []string) (bool, error){
	
	for i, val := range link {
		fileName := fmt.Sprintf("%d.jpg", i+1)
		_, err := CreateDirIfNotExist(filepath.Join("folder_manhwa",fdName,chapName))
		if err != nil {
			log.Print(err)
			return false, nil
		}
		flName := filepath.Join("folder_manhwa",
				fdName,
				chapName,
				fileName,
		)
		imgMe, err := os.Create(flName)
		if err != nil {
			log.Print(err)
			return false, nil
		}
		defer imgMe.Close()

		resp, err := http.Get(val)
		if err != nil {
			return false, err
		}
		log.Println(val)
		defer resp.Body.Close()
		
		size, err := io.Copy(imgMe, resp.Body)
		if err != nil {
			log.Fatal(err)
			return false, err
		}
		log.Printf("%s || %s || Page > %d || File Size > %d\n", title, chapName, i, size/1000)

	}
	return true, nil
}

func CreateDirIfNotExist(dir string) (bool, error){
	if _, err := os.Stat(dir); os.IsNotExist(err) {
			err = os.MkdirAll(dir, 0755)
			if err != nil {
				log.Panic(err)
				return false, nil
			}
	}
	return true, nil
}

func init()  {
	tr := &http.Transport{
		MaxIdleConns: 20,
		MaxIdleConnsPerHost: 20,
	}
	netClient = &http.Client{Transport: tr}
}