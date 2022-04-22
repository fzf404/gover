/*
 * @Author: fzf404
 * @Date: 2022-02-25 12:34:38
 * @LastEditTime: 2022-02-27 08:47:40
 * @Description: 配置初始化
 */
package initialize

import (
	"gover/global"
	"log"

	"github.com/spf13/viper"
)

func Config() {

	// 文件路径
	viper.AddConfigPath("./config")

	// 文件信息
	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	// 读取配置信息
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Read config failed: ", err)
	}

	// 解析至结构体
	if err := viper.Unmarshal(&global.CONFIG); err != nil {
		log.Fatal("Unmarshal config failed: ", err)
	}
}
