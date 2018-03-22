package config

//配置文件结构体
type DevContext struct {
	RedisIp      string
	RedisPass    string
	MysqlIp      string
	MysqlAdmin   string
	MysqlPass    string
	MongodbIp    string
	MongodbAdmin string
	MongodbPass  string
}

//配置文件
var DevCtx = DevContext{}
