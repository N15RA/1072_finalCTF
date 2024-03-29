package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	flag := `NISRA{QxDU6pmgd4SQrKyte9pCb7yr9Sh95BMk2wFumQhmxtRz9ELkkHhVPhQuCWgNJFxqMrNumWiUrkvKecDbiXutqz9rtBa2Wv7XKk4dN9MC4gdPjzeuTpXDmqabMkNmZHi8ny7fCReQJgmg5y2NnKcCV97GE7yebaKd2MecxxQcj68NWHZF47qLV8PHj6gyX63xAb3LJzYbzYqDx8dHXxZ3yWr5XMthw9i8XZj3vWtxuZgxGYi74ikHGXUGVARGWRiRVrhd27gWQDjYzBN8kTSaJpKuSNCDxLRTJ6HXycLvav6rQXvcWVQ4d7N8BGkrvuyGMYpLxwpwYXiDwpuceiLCrStU5U2CDx4fC8LBZMRYGWjZwQyADdHZP4UwkxD4JbMyELURtxjD3FUNKvPrP8L6i2iHxxHgKaEDRvfKTEQP8vb4q9i97YaZnjXT2JCf3NS6c6GfPuvWaNgWBXWnanbRzuuUXB7mhaMRTrXmrAJfgJgudDNt97krhT4etjWxEUkmEMJaNbn25pwYmrHtkdziz5D49eqFY4VyfQ2B4UjMvHJhScYaUyjwMYNT9CpzkhUKG4DUR6ecHWWkMCPVVNQHme3kYUTxvVKPAu6Bt4vhXULT7QBMxwSXbH8wVFVmt3YVrntb3aY8rPVQQ5qBeW5CxbiZcKcdBgjE3BcEBu3jtujqWQACMGK8TE2Rk2qmmvVYDvjxYjz7MpeLBirYAci4nY44NdVRf7fAKwkiFjMRYpCurkwzQR8YxxEbZMrkZAiUuyiRADijCvCgHkyuRmPE9hWHRy9aHMaTdGyzRUynqMynDhFPrS3BHXvEMMxaeSPfR5yMxf9i4nGcNrTcgQQEGiM3ECmhUaAhUrK73rHK2eBQgHcCCnCQg898HYmL2ZqEWa6CY4ffgRVHRbyjYbcgtZTirfbJaGCRWVuCnQnH7pKSufSQczw3SJNbXZewqwwL6Act4pPKNxMpEEeGpMriAWnvFfyWPutKGdCifBFwNRYNTv4kDRE7tf2mXSW6UPKBBE2CAR7HJS5PeGKBu9LUaZkkStHqASmHU7wGYRDgV7NuRtGPJQkbGvD9pkB6FjLd7wepN8TAqTuKL74iANzQfMpDNQDbT4q3PCAqFbbweycukUxZkwAdmw2Zu73eJ3H9tpmFjc325mv62A5kkjf9EmcnxrgdvM64EJyzhgByf6XCjBagMvGzwzCSMkMFadCAKqa43G3Z84XDbCXzivkgzV25LmMipxt4YmPvLYrtuvpvjYbd9Zw9qwFrEPJYycRWzAnX5tTRcgCM6YideQuzZdUbAm8h3We45PhgTTKZKe9kv554PPKKXHaLj}`

	r := gin.Default()

	r.GET("/id", func(c *gin.Context) {

		nstr := c.Query("n")

		if len(nstr) == 0 {
			c.Redirect(302, "/next?id=0")
		}

		n, err := strconv.Atoi(nstr)

		if err != nil || n < 0 {
			n = 0
		}

		n = n % len(flag)

		c.String(200, "%s", flag[n:n+1])

	})

	r.GET("/", func(c *gin.Context) {

		c.Redirect(302, "/id?n=0")

	})

	r.Run() // listen and serve on 0.0.0.0:8080
}
