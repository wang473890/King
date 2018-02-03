package king

import (
	"github.com/gin-gonic/gin"
	"redis"
	"log"
	"net/http"
)

func Get(c *gin.Context) {
	data, err := redis.DevRedeisConn.RedisGet("king")
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ret":  "fail",
			"data": "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ret":  "ok",
		"data": data,
	})
}
