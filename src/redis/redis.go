package redis

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

const REDISCONNECTSTRING = "176.122.165.37:6379";

const REIDS_PASSWORD = "meizizi";

type RedisConn struct {
	redis.Conn
}

var DevRedeisConn RedisConn

func InitRedisConn() {
	conn, err := redis.Dial("tcp", REDISCONNECTSTRING)
	//conn , err := redis.DialTimeout("tcp", REDISCONNECTSTRING, 0, 1*time.Second, 1*time.Second)
	if err != nil {
		fmt.Println(err)
	}

	if _, err := conn.Do("AUTH", REIDS_PASSWORD); err != nil {
		conn.Close()
		fmt.Println(err)
	}
	DevRedeisConn.Conn = conn;
}

func (*RedisConn) RedisSet(key, value string, expire int) (bool, error) {
	_, err := DevRedeisConn.Conn.Do("SET", key, value);
	_, err1 := DevRedeisConn.Conn.Do("EXPIRE", key, expire);
	if err != nil {
		fmt.Println(err)
		return false, nil
	}
	if err1 != nil {
		fmt.Println(err1)
		return false, nil
	}
	return true, nil
}

func (*RedisConn) RedisGet(key string) (string, error) {
	data, err := DevRedeisConn.Conn.Do("GET", key);
	if err != nil {
		fmt.Println(err)
	}
	if data == nil {
		return "", nil
	}
	return string(data.([]byte)), nil
}
