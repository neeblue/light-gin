package api

import (
	"github.com/gin-gonic/gin"
	"light-gin/app/model"
	"light-gin/core/errors"
	"strconv"
)

func Get(c *gin.Context) {
	id := c.Param("id")

	i, err := strconv.Atoi(id)
	if err != nil {
		errors.HandleErr(1000, c)
		return
	}

	c.JSON(200, gin.H{
		"msg": i,
	})

}

func Query(c *gin.Context) {
	c.JSON(200, generateBooks())
}

func generateBooks() []model.Book {
	return []model.Book{
		{
			Title:  "The Art Of Loving",
			Author: "Erich Fromm",
		},
	}
}
