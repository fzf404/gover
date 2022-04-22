/*
 * @Author: fzf404
 * @Date: 2022-02-22 20:38:43
 * @LastEditTime: 2022-02-27 10:03:20
 * @Description: 用户请求
 */
package user

type (
	// 登录请求
	LoginRequest struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password" binding:"min=8,max=32"`
	}
	// 注册请求
	RegistRequest struct {
		UserName   string `json:"user_name" binding:"min=3,max=20"`
		Password   string `json:"password"  binding:"min=8,max=32"`
		Email      string `json:"email" binding:"email"`
		VerifyCode string `json:"verify_code"  binding:"len=6"`
	}
	// 个人信息请求
	InfoRequest struct {
		UserName string `uri:"user_name" binding:"min=3,max=20"`
	}
)
