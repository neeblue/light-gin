package main

import (
	"github.com/gin-gonic/gin"
	"light-gin/router"
)

func main()  {
	app := gin.Default()

	router.InitRouter(app)

	app.Run()
}
