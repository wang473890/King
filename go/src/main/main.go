package main

import (
	"github.com/gin-gonic/gin"
	"train"
	"time"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	goSyncTasks()
	train.Main()
}
func goSyncTasks() {
	//自动设置告警信息已读
	go func() {
		//ticker := time.NewTicker(time.Duration(24) * time.Hour)
		ticker := time.NewTicker(time.Duration(60) * time.Second)
		for {
			select {
			case <-ticker.C:
				train.TimeTask()
			}
		}
	}()
}
