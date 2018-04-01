package main

import (
	"task"
	"test"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	task.Main()
	test.Main()
}
