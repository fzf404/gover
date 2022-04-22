/*
 * @Author: fzf404
 * @Date: 2022-02-27 15:00:08
 * @LastEditTime: 2022-03-01 04:00:58
 * @Description: api 入口
 */
package api

import (
	"gover/service"
)

var (
	userService     service.UserService
	categoryService service.CategoryService
	// tagService      service.TagService
	articleService service.ArticleService
)
