package main

import (
	"github.com/gin-gonic/gin"
	"admin"
	"train"
	"fmt"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	moduleName := "admin"
	switch moduleName {
	case "train":
		train.Main()
		break
	case "admin":
		admin.Main()
		break
	default:
		fmt.Println("invalid module")
	}
}
