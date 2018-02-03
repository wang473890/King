package king

import (
	"github.com/gin-gonic/gin"
	"db"
)

func Main() {
	//初始化DB资源
	db.InitDb()
	defer db.DevContext.Db.Close()
	router := gin.Default()
	router.GET("/king/add", Add)
	router.GET("/king/delete", Delete)
	router.GET("/king/find", Find)
	router.GET("/king/find_all", FindAll)
	router.GET("/king/update", Update)
	router.Run(":8010")
}
