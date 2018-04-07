package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"db"
	"common"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"encoding/base64"
)

func Login(c *gin.Context) {
	var admin TplAdmin
	var l int
	name := c.PostForm("name")
	l = len(name)
	if 0 < l && 18 < l {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请输6字以内的用户名",
			"code": 20001,
		})
		return
	}
	pass := c.PostForm("pass")
	l = len(pass)
	if 6 < l && 12 > l {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请输入6到12位密码",
			"code": 20002,
		})
		return
	}

	e := db.My.QueryRow("select FPass,FSale from tbl_admin where FName = ?", name).Scan(&admin.FPass, &admin.FSale)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "管理员不存在",
			"code": 20003,
		})
		return
	}
	md5Data := common.TakeMd5(common.TakeMd5(pass) + admin.FSale)
	if md5Data != admin.FPass {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "密码错误",
			"code": 20004,
		})
		return
	}
	adminInfo := base64.StdEncoding.EncodeToString([]byte(name))
	fmt.Println(adminInfo)
	session := sessions.Default(c)
	session.Set("admin", adminInfo)
	e = session.Save()
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "登陆失败",
			"code": 20005,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":  "登陆成功",
		"code": 0,
	})

}
