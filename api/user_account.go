/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:19:32
 * @LastEditTime: 2022-02-27 14:58:58
 * @Description: 用户账户
 */

package api

import (
	"gover/model/response"
	"gover/model/user"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 注册
 * @param {*gin.Context} c
 */
func Regist(c *gin.Context) {

	var r user.RegistRequest

	if err := c.ShouldBind(&r); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	userService.Regist(r, c)
}

/**
 * @description: 登录
 * @param {*gin.Context} c
 */
func Login(c *gin.Context) {
	var l user.LoginRequest

	if err := c.ShouldBind(&l); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}
	userService.Login(l, c)
}
