/*
 * @Author: fzf404
 * @Date: 2022-02-28 12:37:41
 * @LastEditTime: 2022-02-28 12:44:26
 * @Description: 文章信息
 */
package api

import (
	"gover/model/article"
	"gover/model/response"

	"github.com/gin-gonic/gin"
)

func ArticleList(c *gin.Context) {

	var al article.ListRequest
	if err := c.ShouldBindQuery(&al); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	articleService.List(al.Page, al.Count, c)
}
