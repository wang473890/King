package admin

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Test(c *gin.Context) {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
		"is": "ok",
	})
}
