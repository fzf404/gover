/*
 * @Author: fzf404
 * @Date: 2022-02-22 20:38:47
 * @LastEditTime: 2022-02-27 10:03:27
 * @Description: 用户响应
 */
package user

type (
	// 登录响应
	LoginResponse struct {
		UserName string `json:"user_name"`
		NickName string `json:"nick_name"`
		Type     uint   `json:"type"`
		Avatar   string `json:"avatar"`
		Token    string `json:"token"`
	}
	// 注册响应
	RegistResponse struct {
		UserName string `json:"user_name"`
	}
	// 用户信息响应
	InfoResponse struct {
		UserName string `json:"user_name"`
		Type     uint   `json:"type"`
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
	}
	// 个人信息响应
	MeInfoResponse struct {
		UserName string `json:"user_name"`
		Type     uint   `json:"type"`
		Email    string `json:"email"`
		NickName string `json:"nick_name"`
		Avatar   string `json:"avatar"`
	}
)
