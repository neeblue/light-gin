package router

import (
	"github.com/gin-gonic/gin"
	"light-gin/app/api"
)

func InitRouter(app *gin.Engine) {

	v1 := app.Group("/v1")
	{
		book := v1.Group("/book")
		{
			book.GET("/", api.Query)
			book.GET(":id", api.Get)
		}
	}
}
