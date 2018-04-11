package admin

import (
	"github.com/gin-gonic/gin"
	"time"
	"db"
	"log"
	"net/http"
	"common"
)

func Sign(c *gin.Context) {
	var admin TplAdmin
	var l int
	name := c.PostForm("name")
	l = len(name)
	if 0 < l && 18 < l {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请输6字以内的用户名",
			"code": 10001,
		})
		return
	}
	pass := c.PostForm("pass")
	l = len(pass)
	if 6 > l && 12 < l {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请输入6到12位密码",
			"code": 10002,
		})
		return
	}
	pass2 := c.PostForm("pass2")
	if pass2 == "" {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "请输入确认密码",
			"code": 10003,
		})
		return
	}
	if pass != pass2 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "两次密码不一致",
			"code": 10004,
		})
		return
	}
	var count int
	e := db.My.QueryRow("select count(*) from tbl_admin where FName = ?", name).Scan(&count)
	if e == nil && count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "管理员名称已存在",
			"code": 10005,
		})
		return
	}
	admin.FName = name
	admin.FSale = common.GetRandomSalt()
	md5Data := common.TakeMd5(common.TakeMd5(pass) + admin.FSale)
	admin.FPass = md5Data
	admin.FCreateTime = time.Now().Unix()
	admin.FUpdateTime = time.Now().Unix()
	admin.FLevel = 4
	_, e = db.My.Exec("INSERT INTO tbl_admin(FName, FPass,FSale,FCreateTime,FUpdateTime,FLevel) VALUES (?,?,?,?,?,?)", admin.FName, admin.FPass, admin.FSale, admin.FCreateTime, admin.FUpdateTime, admin.FLevel)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":  "管理员创建失败",
			"code": 10006,
		})
		return
	}
	if e != nil {
		log.Fatalln(e)
	}
	c.JSON(http.StatusOK, gin.H{
		"msg":  "管理员创建成功",
		"code": 0,
	})
}
