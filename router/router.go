/*
 * @Author: fzf404
 * @Date: 2022-02-25 12:18:53
 * @LastEditTime: 2022-08-25 20:24:36
 * @Description: 路由初始化
 */
package router

import (
	"gover/api"
	"gover/global"
	"gover/middleware"

	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.Use(middleware.Cors())

	route := r.Group("/api")
	{
		route.GET("/ping", api.Ping)

		route.GET("/user/info/:user_name", api.UserInfo)

		route.POST("/user/regist", api.Regist)
		route.POST("/user/login", api.Login)
		// route.POST("/user/send_code", api.SendCode)

		route.GET("/article/list", api.ArticleList)
		// route.GET("/article/list/:user_name", api.UserArticleList)
		// route.GET("/article/detail/:article_id", api.ArticleDetail)

		// 普通用户
		route.Use(middleware.AuthJWT(global.CONSUMEER))
		{
			route.GET("/user/info", api.MeInfo)

			route.POST("/article/create", api.CreateArticle)
			// route.POST("/article/update", api.UpdateArticle)
			// route.POST("/article/delete", api.DeleteArticle)
		}

		// 管理员
		route.Use(middleware.AuthJWT(global.ADMIN))
		{
			route.POST("/category/create", api.CreateCategory)
		}
	}

	return r
}
