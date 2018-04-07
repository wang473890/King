package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"
)

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	e := session.Save()
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登出失败",
			"code": 30001,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "登出成功",
		"code": 0,
	})
	return
}
