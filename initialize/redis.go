/*
 * @Author: fzf404
 * @Date: 2022-02-10 10:44:20
 * @LastEditTime: 2022-02-27 15:00:37
 * @Description: redis 初始化
 */
package initialize

import (
	"gover/global"
	"log"

	"github.com/go-redis/redis"
)

func Redis() {

	// 获取 redis 配置
	r := global.CONFIG.Redis

	// 建立连接
	client := redis.NewClient(&redis.Options{
		Addr:     r.Addr,
		Password: r.Password,
		DB:       r.DB,
	})

	// 测试连接
	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatal("Connect redis failed: ", err)
	} else {
		global.REDIS = client
		log.Print("Connect redis success: ", pong)
	}
}
