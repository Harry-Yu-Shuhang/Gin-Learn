package cache

import (
	"context"
	"gin-learn/config"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb  *redis.Client
	Rctx context.Context
)

func init() { //当go执行的时候会默认执行init
	Rdb = redis.NewClient(&redis.Options{
		Addr:     config.RedisAddress,
		Password: config.RedisPassword, // 没有密码，默认值
		DB:       config.RedisDb,       // 默认DB 0
	})
	Rctx = context.Background() //保存上下文
} //redis客户端初始化

func Zscore(id int, score int) redis.Z {
	return redis.Z{Score: float64(score), Member: id} //每一个用户得了多少分
}
