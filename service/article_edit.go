/*
 * @Author: fzf404
 * @Date: 2022-02-27 09:10:59
 * @LastEditTime: 2022-03-01 04:17:52
 * @Description: 文章详情
 */
package service

import (
	"gover/model/article"
	"gover/model/response"
	"gover/model/tag"

	"github.com/gin-gonic/gin"
)

func (as *ArticleService) Create(ar article.CreateRequest, userID uint, c *gin.Context) {
	// 标签处理
	var tags []tag.Tag
	for _, value := range ar.Tags {
		// 查找标签是否存在
		t, err := tagModel.FindByName(value)
		if err != nil {
			// 不存在则新建标签
			t = &tag.Tag{Name: value}
			tagModel.Insert(t)
			// 判断是否新建成功
			if t.ID == 0 {
				response.Fail(response.DBErr, c)
				return
			}
		}
		tags = append(tags, *t)
	}

	// 创建文章
	a := &article.Article{
		UserID:     userID,
		CategoryID: ar.CategoryID,
		Tags:       tags,

		Cover: ar.Cover,

		Title:   ar.Title,
		Summary: ar.Summary,
		Content: ar.Content,
	}
	articleModel.Insert(a)

	// 判断是否成功
	if a.ID == 0 {
		response.Fail(response.DBErr, c)
	} else {
		response.Ok(c)
	}
}
