package king

import (
	"github.com/gin-gonic/gin"
	"redis"
	"log"
	"net/http"
)

func Set(c *gin.Context) {
	id := c.Query("id")
	_, err := redis.DevRedeisConn.RedisSet("king", id, 10)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{
			"ret": "fail",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ret": "ok",
	})
}
