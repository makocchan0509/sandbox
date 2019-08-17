package dao

import (
	"log"
	"projects/Services/common/properties"

	"github.com/garyburd/redigo/redis"
)

func GetConnectionRedis() redis.Conn {
	//Connect to Redis

	prop := properties.GetProp()

	redisHost := prop.Redis.RedisHost
	redisPort := prop.Redis.RedisPort
	redisPrtc := prop.Redis.RedisPrtc

	rconn, err := redis.Dial(redisPrtc, redisHost+":"+redisPort)
	if err != nil {
		log.Println("error: ", err.Error())
	}
	return rconn
}

func SetRedis(key string, value string, rconn redis.Conn) {
	rconn.Do("SET", key, value)
}

func GetRedis(key string, rconn redis.Conn) string {
	s, err := redis.String(rconn.Do("GET", key))
	if err != nil {
		log.Println("error: ", err.Error())
	}
	return s
}

func ExistsRedis(key string, rconn redis.Conn) int {
	i, err := redis.Int(rconn.Do("EXISTS", key))
	if err != nil {
		log.Println("error: ", err.Error())
	}
	return i
}

func ExpireRedis(key string, sec int, rconn redis.Conn) {
	rconn.Do("EXPIRE", key, sec)
}

func CloseConnectionRedis(rconn redis.Conn) {
	defer rconn.Close()
}
