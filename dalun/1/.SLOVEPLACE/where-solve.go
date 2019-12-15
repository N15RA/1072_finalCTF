package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {

	for index := 0; index < 4144; index++ {

		resp, _, _ := gorequest.New().Get("http://localhost:8080/id?n=" + fmt.Sprintf("%d", index)).End()

		for i, v := range resp.Header {
			for _, q := range v {
				if q == "I dont want to eat pineapple pizze QQ!" {
					fmt.Print(i)
				}
			}
		}

	}

}
