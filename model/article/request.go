/*
 * @Author: fzf404
 * @Date: 2022-02-27 12:48:31
 * @LastEditTime: 2022-02-28 13:28:15
 * @Description: 文章请求
 */
package article

type (
	CreateRequest struct {
		CategoryID uint     `json:"category_id" binding:"required"`
		Tags       []string `json:"tags" binding:"max=5"`
		Cover      string   `json:"cover" binding:"required"`

		Title   string `json:"title" binding:"min=3,max=50"`
		Summary string `json:"summary" binding:"min=10,max=50"`
		Content string `json:"content" binding:"min=20,max=10000"`
	}
	
	UpdateRequest CreateRequest

	ListRequest struct {
		Page  int `form:"page" binding:"required"`
		Count int `form:"count" binding:"required"`
	}
)
