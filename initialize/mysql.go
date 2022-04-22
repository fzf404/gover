/*
 * @Author: fzf404
 * @Date: 2022-01-22 13:56:31
 * @LastEditTime: 2022-02-27 12:59:44
 * @Description: Mysql 初始化
 */

package initialize

import (
	"gover/global"
	"gover/model/article"
	"gover/model/category"
	"gover/model/tag"
	"gover/model/user"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/**
 * @description: 初始化数据库
 */
func Mysql() {

	// 获取 mysql 配置
	mCfg := global.CONFIG.Mysql.Dsn()

	// 连接
	if db, err := gorm.Open(mysql.Open(mCfg)); err != nil {
		log.Fatal("Connect mysql failed: ", err)
	} else {
		// 自动建表
		db.AutoMigrate(&user.User{})
		db.AutoMigrate(&category.Category{})
		db.AutoMigrate(&tag.Tag{})
		db.AutoMigrate(&article.Article{})
		global.DB = db
		log.Print("Connect mysql success: ", mCfg)

	}
}
