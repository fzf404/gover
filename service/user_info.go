/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:26:10
 * @LastEditTime: 2022-02-27 15:05:19
 * @Description: 用户账户
 */

package service

import (
	"gover/model/response"
	"gover/model/user"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 用户信息
 * @param {user.InfoRequest} infoRequest
 * @param {*gin.Context} c
 */
func (us *UserService) UserInfo(infoRequest user.InfoRequest, c *gin.Context) {

	// 查询用户信息
	u, err := userModel.FindByUserName(infoRequest.UserName)
	if err != nil {
		response.FailWithMessage(response.NotFind, "用户不存在", c)
	} else {
		response.OkWithData(user.InfoResponse{
			UserName: u.UserName,
			Type:     u.Type,
			NickName: u.NickName,
			Avatar:   u.Avatar,
		}, c)
	}
}

/**
 * @description: 个人信息
 * @param {user.User} u
 * @param {*gin.Context} c
 */
func (us *UserService) MeInfo(u user.User, c *gin.Context) {

	response.OkWithData(user.MeInfoResponse{
		UserName: u.UserName,
		Type:     u.Type,
		Email:    u.Email,
		NickName: u.NickName,
		Avatar:   u.Avatar,
	}, c)

}
