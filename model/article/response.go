/*
 * @Author: fzf404
 * @Date: 2022-02-27 12:48:31
 * @LastEditTime: 2022-02-28 13:30:42
 * @Description: 文章响应
 */
package article

type (
	ListResponse struct {
		UserName string `json:"user_name"`

		Category string   `json:"category"`
		Tags     []string `json:"tags"`
		Cover    string   `json:"cover"`

		Title   string `json:"title"`
		Summary string `json:"summary"`
	}

	DetailResponse struct {
		ListResponse
		Content string `json:"content"`
	}
)
