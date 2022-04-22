/*
 * @Author: fzf404
 * @Date: 2022-02-27 14:57:09
 * @LastEditTime: 2022-03-01 04:16:52
 * @Description: 分类请求
 */
package api

import (
	"gover/model/category"
	"gover/model/response"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	var cr category.CreateRequest
	if err := c.ShouldBind(&cr); err != nil {
		response.FailWithDetailed(response.ParamErr, err.Error(), "提交信息非法", c)
		return
	}

	categoryService.Create(cr.Name, cr.Desc, c)
}
