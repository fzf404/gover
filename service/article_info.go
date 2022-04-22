/*
 * @Author: fzf404
 * @Date: 2022-02-27 09:11:42
 * @LastEditTime: 2022-03-01 04:18:30
 * @Description: 文章编辑
 */

package service

import (
	"gover/model/article"
	"gover/model/response"
	"log"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 文章列表
 * @param {int} page
 * @param {int} count
 * @param {*gin.Context} c
 */
func (as *ArticleService) List(page int, count int, c *gin.Context) {

	// 全部文章列表
	al, err := articleModel.FindAll(page, count) // order
	if err != nil {
		log.Fatal(err)
		response.Fail(response.DBErr, c)
		return
	}

	// 是否存在
	if len(*al) == 0 {
		response.FailWithMessage(response.NotFind, "什么都没有", c)
		return
	}

	// 解析文章
	var alr []article.ListResponse
	for _, a := range *al {
		// 解析标签
		var tags []string
		for _, t := range a.Tags {
			tags = append(tags, t.Name)
		}

		alr = append(alr, article.ListResponse{
			UserName: a.User.UserName,

			Category: a.Category.Name,
			Tags:     tags,
			Cover:    a.Cover,

			Title:   a.Title,
			Summary: a.Summary,
		})
	}
	response.OkWithData(alr, c)

}
