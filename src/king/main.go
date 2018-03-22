package king

import (
	"github.com/gin-gonic/gin"
	"common"
	"db"
	"redis"
	"time"
)

func Main() {
	//初始化DB资源
	common.Init()
	defer db.DevContext.Db.Close()
	//初始化Redis资源
	redis.InitRedisConn()
	defer redis.DevRedeisConn.Conn.Close()
	router := gin.Default()
	router.GET("/king/get", AnalogGet)
	router.GET("/king/post", AnalogPost)
	router.Run(":8010")
}

func goSyncTasks() {
	go func() {
		ticker := time.NewTicker(time.Duration(10) * time.Second)
		for {
			select {
			case <-ticker.C:
				Weather()
			}
		}
	}()
}
