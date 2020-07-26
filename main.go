package main

import "github.com/gin-gonic/gin"

func main()  {
	app := gin.Default()

	app.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"msg": "pong",
		})
	})

	app.Run()
}
