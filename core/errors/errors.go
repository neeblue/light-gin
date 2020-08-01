package errors

import "github.com/gin-gonic/gin"

const (
	URL_ERROR  = 1000
	BOOK_EXIST = 2000
)

type Err struct {
	code int
	msg  string
}

var errMsg = map[int]Err{
	URL_ERROR: {
		code: 404,
		msg:  "url错误",
	},
	BOOK_EXIST: {
		code: 404,
		msg:  "书籍不存在",
	},
}

func HandleErr(errorCode int, c *gin.Context) {
	e := errMsg[errorCode]

	c.JSON(e.code, gin.H{
		"errCode": errorCode,
		"msg":     e.msg,
		"request": c.Request.Method + " " + c.FullPath(),
	})
}
