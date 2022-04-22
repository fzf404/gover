/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:17:18
 * @LastEditTime: 2022-02-27 03:18:20
 * @Description: 通用
 */
package api

import "github.com/gin-gonic/gin"

/**
 * @description: Ping
 * @param {*gin.Context} c
 */
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": gin.H{},
		"msg":  "pong",
	})
}
