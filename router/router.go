package router

import (
	"github.com/gin-gonic/gin"
	"light-gin/app/api"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		book := v1.Group("/book")
		{
			book.GET("/", api.GetBook)
		}
	}

	return r
}
