package king

import (
	"github.com/gin-gonic/gin"
	"db"
	"redis"
	"city"
)

func Main() {
	//初始化DB资源
	db.InitDb()
	defer db.DevContext.Db.Close()

	//初始化Redis资源
	redis.InitRedisConn()
	defer redis.DevRedeisConn.Conn.Close()
	router := gin.Default()
	router.GET("/king/add", Add)
	router.GET("/king/delete", Delete)
	router.GET("/king/find", Find)
	router.GET("/king/find_all", FindAll)
	router.GET("/king/update", Update)
	router.GET("/king/set", Set)
	router.GET("/king/get", Get)
	router.GET("/city/list", city.CityList)
	router.Run(":8010")
}
