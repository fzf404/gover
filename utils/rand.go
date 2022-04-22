/*
 * @Author: fzf404
 * @Date: 2022-02-11 00:45:45
 * @LastEditTime: 2022-02-27 14:20:41
 * @Description: 随机生成
 */
package utils

import (
	"math/rand"
	"time"
)

/**
 * @description: 随机数字字符串
 * @param {int} length
 */
func RandomNumber(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

/**
 * @description: 随机字符串
 * @param {int} length
 */
func RandomString(length int) string {
	rand.Seed(time.Now().UnixNano())
	letters := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
