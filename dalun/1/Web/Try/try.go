package main

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tily/gostrcrypt"
)

func aes256toint(enstr string, key string) int {

	sc := gostrcrypt.StrCrypt{key}

	decrypted, err := sc.Decrypt(enstr)
	if err != nil {
		return 0
	}

	id, err := strconv.Atoi(decrypted)

	if err != nil {
		return 0
	}

	return id
}

func inttoaes256(id int, key string) string {

	sc := gostrcrypt.StrCrypt{key}

	encrypted, _ := sc.Encrypt(fmt.Sprintf("%d", id))

	return encrypted
}

func main() {

	key := "hands0meda1un-secret-key"
	flag := `NISRA{NICE_BOT_BRO}`

	r := gin.Default()
	r.LoadHTMLGlob("tmpl/try.tmpl")

	r.GET("/next", func(c *gin.Context) {

		idstr := c.Query("id")

		if len(idstr) < 30 {
			c.Redirect(302, "/next?id="+inttoaes256(1, key))
		}

		id := 0
		id = aes256toint(idstr, key)

		if id == 0 {
			c.Redirect(302, "/next?id="+inttoaes256(1, key))
		} else {

			f := ""

			if id == 13337 {
				f = flag
				fmt.Println("id:", id)
			}

			c.HTML(200, "try.tmpl", gin.H{
				"now":    id,
				"nextid": inttoaes256(id+1, key),
				"flag":   f,
			})
		}

	})

	r.GET("/", func(c *gin.Context) {

		c.Redirect(302, "/next?id="+inttoaes256(1, key))

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
