/*
 * @Author: fzf404
 * @Date: 2022-02-27 06:16:42
 * @LastEditTime: 2022-02-27 06:48:29
 * @Description: description
 */
package utils

import (
	"fmt"
	"gover/global"
	"gover/model/response"
	"log"
	"net/smtp"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jordan-wright/email"
)

/**
 * @description: 发送邮箱验证码
 * @param {string} emailAddr
 * @param {*gin.Context} c
 */
func SendCode(emailAddr string, c *gin.Context) {

	key := emailAddr + ".sended"

	// 是否超过频率
	if err := global.REDIS.Get(key).Err(); err == nil {
		response.FailWithMessage(response.ParamErr, "验证码发送过于频繁", c)
		return
	}

	// 验证码生成
	verfiyCode := RandomNumber(6)
	// 验证码过期时间
	expires := global.CONFIG.Email.Expires
	// 构造邮件的内容
	mail := &email.Email{
		From:    fmt.Sprintf("%v <%v>", global.CONFIG.Email.Name, global.CONFIG.Email.Addr),
		To:      []string{emailAddr},
		Subject: "邮箱验证",
		Text:    []byte("【Golang】验证码：" + verfiyCode + "，有效期 " + strconv.Itoa(expires) + " 分钟"),
	}

	// 发送邮件
	if err := mail.Send(global.CONFIG.Email.Smtp, smtp.PlainAuth("", global.CONFIG.Email.Addr, global.CONFIG.Email.Password, global.CONFIG.Email.Host)); err != nil {
		log.Print("Email send error: ", err)
		response.FailWithMessage(response.ParamErr, "发送失败, 请确认邮箱是否可用", c)
	} else {
		// 重新发送限制
		global.REDIS.Set(key, 1, time.Second*50)
		// 存储验证码
		global.REDIS.Set(emailAddr, verfiyCode, 15*time.Minute)

		response.Ok(c)
	}
}

/**
 * @description: 验证邮箱验证码
 * @param {string} email
 * @param {string} verfiyCode
 */
func VerifyEmail(email string, verfiyCode string) bool {

	// 测试时允许 114514
	if gin.Mode() == gin.DebugMode {
		if verfiyCode == "114514" {
			return true
		}
	}

	// 获取 Redis 中的验证码
	redisCode, err := global.REDIS.Get(email).Result()
	if err != nil {
		return false
	}

	// 比较验证码
	return verfiyCode == redisCode
}
