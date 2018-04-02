package train

import (
	"github.com/gin-gonic/gin"
	"db"
)

func Main(){
	router := gin.Default()
	db.InitSessions()
	defer db.Mgo.MgoSession.Close()
	router.GET("/test",GetData)
	router.POST("/post_row",PostRow)
	router.Run(":8001")
}
