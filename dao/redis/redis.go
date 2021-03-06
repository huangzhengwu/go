package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"web_app/settings"
)

// 声明一个全局的rdb变量
var client *redis.Client

// 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	client = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.Db,       // use default DB
		PoolSize: cfg.PoolSize, // use default DB
	})

	_, err = client.Ping().Result()
	if err != nil {
		zap.L().Info("connect redis success")

	}
	return
}

func Close() {
	_ = client.Close()
}
