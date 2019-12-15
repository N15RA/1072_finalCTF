package main

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/parnurzeal/gorequest"
)

func main() {

	k := ""

	for index := 0; index < 13337; index++ {

		resp, _, _ := gorequest.New().Get("http://localhost:8080/next?id=" + k).End()

		doc, err := goquery.NewDocumentFromReader(resp.Body)
		if err != nil {
			log.Fatal(err)
		}

		k = doc.Find("data").Text()
		fmt.Println(k)

	}

}
