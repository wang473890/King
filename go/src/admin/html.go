package admin

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func Html(c *gin.Context) {

	t, _ := template.ParseFiles("/Users/wanggang/Documents/doc/学习/test/go/king/go/src/web/index.html")
	t.Execute()
}
