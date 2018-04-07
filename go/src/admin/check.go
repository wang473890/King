package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"
	"encoding/base64"
)

func Check(c *gin.Context) {
	session := sessions.Default(c)
	inter := session.Get("admin")
	if inter == nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "管理员尚未登陆",
			"code": 40001,
		})
		return
	}
	info := inter.(string)
	decode, _ := base64.StdEncoding.DecodeString(info)
	c.JSON(http.StatusOK, gin.H{
		"msg":  "管理员 " + string(decode) + " 已登陆",
		"code": 0,
	})
}
