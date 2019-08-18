package dao

import (
	"projects/Services/common/properties"

	"github.com/garyburd/redigo/redis"
)

func GetConnectionRedis() (redis.Conn, error) {
	//Connect to Redis

	prop := properties.GetProp()

	redisHost := prop.Redis.RedisHost
	redisPort := prop.Redis.RedisPort
	redisPrtc := prop.Redis.RedisPrtc

	rconn, err := redis.Dial(redisPrtc, redisHost+":"+redisPort)
	return rconn, err
}

func SetRedis(key string, value string, rconn redis.Conn) {
	rconn.Do("SET", key, value)
}

func GetRedis(key string, rconn redis.Conn) (string, error) {
	s, err := redis.String(rconn.Do("GET", key))
	return s, err
}

func ExistsRedis(key string, rconn redis.Conn) (int, error) {
	i, err := redis.Int(rconn.Do("EXISTS", key))
	return i, err
}

func ExpireRedis(key string, sec int, rconn redis.Conn) {
	rconn.Do("EXPIRE", key, sec)
}

func CloseConnectionRedis(rconn redis.Conn) {
	defer rconn.Close()
}
