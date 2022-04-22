/*
 * @Author: fzf404
 * @Date: 2022-02-24 19:25:45
 * @LastEditTime: 2022-02-27 10:07:33
 * @Description: description
 */
package main

import (
	"gover/global"
	"gover/initialize"
	"gover/router"
	"gover/service"
)

func init() {
	initialize.Config()
	initialize.Mysql()
	initialize.Redis()
	service.Init()
}

func main() {
	r := router.Router()
	r.Run(global.CONFIG.Common.Host + ":" + global.CONFIG.Common.Port)
}
