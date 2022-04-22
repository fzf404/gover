package service

import (
	"gover/model/article"
	"gover/model/category"
	"gover/model/tag"
	"gover/model/user"
)

type (
	// 文章服务
	ArticleService struct{}
	// 分类服务
	CategoryService struct{}
	// 用户服务
	UserService struct{}
)

var (
	// 用户模型
	userModel user.UserModel
	// 分类模型
	categoryModel category.CategoryModel
	// 标签模型
	tagModel tag.TagModel
	// 文章模型
	articleModel article.ArticleModel
)

func Init() {
	userModel = user.NewModel()
	categoryModel = category.NewModel()
	tagModel = tag.NewModel()
	articleModel = article.NewModel()
}
