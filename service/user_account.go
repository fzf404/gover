/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:26:10
 * @LastEditTime: 2022-02-27 10:06:40
 * @Description: 用户账户
 */

package service

import (
	"gover/model/response"
	"gover/model/user"
	"gover/utils"
	"log"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @description: 注册服务
 * @param {user.RegistRequest} r
 * @param {*gin.Context} c
 */
func (us *UserService) Regist(r user.RegistRequest, c *gin.Context) {

	// 验证码验证
	if ok := utils.VerifyEmail(r.Email, r.VerifyCode); !ok {
		response.FailWithMessage(response.ParamErr, "验证码错误", c)
		return
	}

	// 用户名 & 邮箱验证
	if _, err := userModel.FindByUserName(r.UserName); err == nil {
		response.FailWithMessage(response.ParamErr, "用户名已存在", c)
		return
	}
	if _, err := userModel.FindByEmail(r.Email); err == nil {
		response.FailWithMessage(response.ParamErr, "邮箱已存在", c)
		return
	}

	// 密码加密
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(r.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Print("Password hash error:", err)
		response.FailWithMessage(response.ServerErr, "密码加密错误", c)
	}

	// 创建用户
	u := user.User{
		UUID:     uuid.NewV4().String(),
		UserName: r.UserName,
		NickName: r.UserName,
		Email:    r.Email,
		Password: string(hashPassword),
	}
	userModel.Insert(&u)

	// 判断是否成功
	if u.ID == 0 {
		response.Fail(response.DBErr, c)
	} else {
		response.Ok(c)
	}
}

/**
 * @description: 登录服务
 * @param {user.LoginRequest} l
 * @param {*gin.Context} c
 */
func (us *UserService) Login(l user.LoginRequest, c *gin.Context) {

	// 通过用户名查询
	u, err := userModel.FindByUserName(l.UserName)
	if err != nil {
		// 通过邮箱查询
		u, err = userModel.FindByEmail(l.UserName)
		if err != nil {
			response.FailWithMessage(response.ParamErr, "用户名或密码错误", c)
			return
		}
	}

	// 密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		response.FailWithMessage(response.ParamErr, "用户名或密码错误", c)
		return
	}

	// 生成 token
	token, err := utils.NewJWT().GenerateToken(u)
	if err != nil {
		log.Print("Token generate error: ", err)
		response.FailWithMessage(response.ServerErr, "Token分发错误", c)
	} else {
		response.OkWithData(user.LoginResponse{
			UserName: u.UserName,
			NickName: u.NickName,
			Avatar:   u.Avatar,
			Token:    token,
		}, c)
	}
}
