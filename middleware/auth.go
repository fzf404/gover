/*
 * @Author: fzf404
 * @Date: 2022-02-25 09:45:58
 * @LastEditTime: 2022-02-28 11:05:57
 * @Description: 鉴权
 */
package middleware

import (
	"gover/model/response"
	"gover/model/user"
	"gover/utils"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 鉴权
 * @param {uint} userType 用户类型
 */
func AuthJWT(userType uint) gin.HandlerFunc {
	return func(c *gin.Context) {

		// 读取 token
		tokenString := c.Request.Header.Get("x-token")
		if tokenString == "" {
			response.FailWithMessage(response.NoLogin, "未登录", c)
			c.Abort()
			return
		}

		// 验证 token
		j := utils.NewJWT()
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			response.FailWithMessage(response.TokenExpired, "登录已过期", c)
			c.Abort()
			return
		}

		// 用户验证
		userID := claims.UserID
		userModel := user.NewModel()
		user, err := userModel.Find(userID)
		if err != nil {
			response.FailWithMessage(response.TokenExpired, "登录已过期", c)
			c.Abort()
			return
		}

		if user.Type > userType {
			response.FailWithMessage(response.TokenExpired, "权限不足", c)
			c.Abort()
			return
		}

		// 将用户信息写入上下文
		c.Set("User", *user)
	}
}
