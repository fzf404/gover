/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:19:32
 * @LastEditTime: 2022-02-27 14:57:35
 * @Description: 用户信息
 */

package api

import (
	"gover/model/user"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户信息
 * @param {*gin.Context} c
 */
func UserInfo(c *gin.Context) {

	var ir user.InfoRequest
	c.ShouldBindUri(&ir)

	userService.UserInfo(ir, c)
}

/**
 * @description: 个人信息
 * @param {*gin.Context} c
 */
func MeInfo(c *gin.Context) {
	u := c.MustGet("User").(user.User)
	userService.MeInfo(u, c)
}
