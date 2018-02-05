package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Return(c *gin.Context, code int, msg string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ret":  code,
		"msg":  msg,
		"data": data,
	})
}
