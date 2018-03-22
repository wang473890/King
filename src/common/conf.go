package common

import (
	"os"
	"io/ioutil"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"db"
	"config"
	"redis"
)

//项目环境初始化
func Init() error {
	//加载配置文件
	LoadConf()
	//数据库加载
	InitDb()
	return nil
}

//加载配置文件
func LoadConf() error {
	var configPath string
	switch gin.Mode() {
	case "debug":
		configPath = "/Users/wanggang/Documents/doc/学习/test/go/king/src/config/debug.conf"
		break
	case "release":
		configPath = "/Users/wanggang/Documents/doc/学习/test/go/king/src/config/debug.conf"
		break
	case "test":
		configPath = "/Users/wanggang/Documents/doc/学习/test/go/king/src/config/debug.conf"
		break
	}
	conf, e := os.Open(configPath)
	if e != nil {
		return e
	}
	defer conf.Close()
	file, e := ioutil.ReadAll(conf)
	if e != nil {
		return e
	}
	e = json.Unmarshal(file, &config.DevCtx)
	if e != nil {
		return e
	}
	return nil
}
func InitDb() {
	//mysql
	db.InitDb()
	//redis
	redis.InitRedisConn()
}
