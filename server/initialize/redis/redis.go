package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"nexwind.net/admin/server/initialize/config"
)

// InitRedisSet init redis set.
func InitRedisSet(conf *config.Conf, l *zap.Logger) map[string]*redis.Client {
	var redisMap = make(map[string]*redis.Client)
	for _, v := range conf.RedisSet {
		client := redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", v.Host, v.Port), // Redis 服务器地址
			Password: v.Pwd,                                // Redis 密码，如果没有密码则为空字符串
			DB:       0,                                    // 使用默认的 DB
		})
		redisMap[v.Key] = client
		_, err := client.Ping(context.TODO()).Result()
		if err != nil {
			l.Error("redis instance init failed.", zap.Any("conf", v), zap.Any("err", err))
		}

	}
	return redisMap
}
