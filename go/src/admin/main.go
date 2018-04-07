package admin

import (
	"github.com/gin-gonic/gin"
	"db"
	"github.com/gin-gonic/contrib/sessions"
)

func Main() {
	//连接mysql
	router := gin.Default()
	db.InitMy()
	defer db.My.Close()
	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.GET("/index", Test)
	router.POST("/login", Login)
	router.POST("/sign", Sign)
	router.POST("/logout", Logout)
	router.POST("/check", Check)
	router.POST("/test", Test)
	router.Run(":8000")
}
