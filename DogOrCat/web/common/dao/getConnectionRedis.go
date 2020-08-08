package dao

import (
	"projects/DogOrCat/web/config"

	"github.com/garyburd/redigo/redis"
)

func GetConnectionRedis() (redis.Conn, error) {
	//Connect to Redis

	redisHost := config.Env.RedisHost
	redisPort := config.Env.RedisPort

	rconn, err := redis.Dial("tcp", redisHost+":"+redisPort)
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
