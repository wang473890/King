package city

import (
	"github.com/gin-gonic/gin"
	"db"
	"redis"
)

func Main() {
	//初始化DB资源
	db.InitDb()
	defer db.DevContext.Db.Close()

	//初始化Redis资源
	redis.InitRedisConn()
	defer redis.DevRedeisConn.Conn.Close()
	router := gin.Default()
	router.GET("/city/list", CityList)
	router.Run(":8011")
}
