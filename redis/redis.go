package redis

import (
	"github.com/garyburd/redigo/redis"
)

var (
	/*
	  This variable provide global 
          connection pool to redis server
	*/
	RedisPool *redis.Pool
)
