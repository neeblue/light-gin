package api

import (
	"github.com/gin-gonic/gin"
	"light-gin/app/model"
)

func GetBook(c *gin.Context)  {
	c.JSON(200, generateBooks())
}

func generateBooks() []model.Book {
	return []model.Book{
		{
			Title: "The Art Of Loving",
			Author: "Erich Fromm",
		},
	}
}
