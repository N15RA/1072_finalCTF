package main

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

func main() {

	for index := 0; index < 1344; index++ {

		_, body, _ := gorequest.New().Get("http://localhost:8080/id?n=" + fmt.Sprintf("%d", index)).End()

		fmt.Print(body)

	}

}
