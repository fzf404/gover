/*
 * @Author: fzf404
 * @Date: 2022-02-10 10:52:05
 * @LastEditTime: 2022-02-28 12:36:29
 * @Description: 全局变量
 */
package global

import (
	"gover/config"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"gorm.io/gorm"
)

var (
	DB    *gorm.DB
	REDIS *redis.Client

	LOG    *zap.Logger
	CONFIG config.Config
)

const (
	ADMIN     = 0
	CONSUMEER = 1

	SORT_BY_TIME = 0
	SORT_BY_HOT  = 1
)
