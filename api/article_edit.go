/*
 * @Author: fzf404
 * @Date: 2022-02-12 16:13:38
 * @LastEditTime: 2022-02-27 14:58:17
 * @Description: 文章编辑
 */
package api

import (
	"gover/model/article"
	"gover/model/response"
	"gover/model/user"

	"github.com/gin-gonic/gin"
)

/**
 * @description: 新建文章
 * @param {*gin.Context} c
 */
func CreateArticle(c *gin.Context) {

	var ar article.CreateRequest
	if err := c.ShouldBind(&ar); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	u := c.MustGet("User").(user.User)

	articleService.Create(ar, u.ID, c)
}
