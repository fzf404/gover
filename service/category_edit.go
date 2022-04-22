/*
 * @Author: fzf404
 * @Date: 2022-02-27 12:44:00
 * @LastEditTime: 2022-03-01 03:58:37
 * @Description: description
 */
package service

import (
	"gover/model/category"
	"gover/model/response"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 创建分类
 * @param {string} name
 * @param {string} desc
 * @param {*gin.Context} c
 */
func (cs *CategoryService) Create(name string, desc string, c *gin.Context) {
	// 创建分类
	category := &category.Category{
		Name: name,
		Desc: desc,
	}
	categoryModel.Insert(category)
	if category.ID == 0 {
		response.Fail(response.DBErr, c)
	} else {
		response.Ok(c)
	}
}
