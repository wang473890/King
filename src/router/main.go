package main

import (
	"github.com/gin-gonic/gin"
	"king"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	//gin.SetMode(gin.ReleaseMode)
	king.Main()
}
