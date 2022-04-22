/*
 * @Author: fzf404
 * @Date: 2022-02-27 10:03:07
 * @LastEditTime: 2022-03-01 04:08:40
 * @Description: 分类请求
 */
package category

type (
	CreateRequest struct {
		Name string `json:"name" binding:"required"`
		Desc string `json:"desc" binding:"required"`
	}
)
