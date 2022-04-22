/*
 * @Author: fzf404
 * @Date: 2022-02-27 03:00:23
 * @LastEditTime: 2022-04-15 19:28:11
 * @Description: 文章模型
 */
package article

import (
	"gover/global"
	"gover/model/category"
	"gover/model/tag"
	"gover/model/user"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type (
	ArticleModel interface {
		Insert(a *Article)
		Find(id uint) (*Article, error)
		FindAll(page int, count int) (*[]Article, error)
		FindAllByUser(userID uint, page int, count int) (*[]Article, error)
		FindAllByTitle(title string, page int, count int) (*Article, error)
		Update(a *Article) error
		Delete(id uint) error
	}

	Model struct {
		DB *gorm.DB
	}

	Article struct {
		gorm.Model
		UserID     uint      `gorm:"comment:用户ID"`
		CategoryID uint      `gorm:"comment:分类ID"`
		Tags       []tag.Tag `gorm:"many2many:article_tag;"`
		// association_jointable_foreignkey:article_id;jointable_foreignkey:tag_id"

		Cover string `gorm:"comment:文章封面"`

		Title   string `gorm:"comment:文章标题"`
		Summary string `gorm:"comment:文章摘要"`
		Content string `gorm:"comment:文章内容"`

		Views   uint `gorm:"default:0;comment:文章浏览量"`
		Likes   uint `gorm:"default:0;comment:文章点赞数"`
		Comment uint `gorm:"default:0;comment:文章评论数"`

		// 自动关联
		Category category.Category // 分类
		User     user.User         // 创建者
	}
)

func NewModel() ArticleModel {
	return &Model{
		DB: global.DB,
	}
}

/**
 * @description: 新增文章
 * @param {*Article} a
 */
func (m *Model) Insert(a *Article) {
	m.DB.Create(a)
}

/**
 * @description: 查找文章
 * @param {uint} id
 */
func (m *Model) Find(id uint) (*Article, error) {
	var a *Article
	err := m.DB.First(a, id).Error
	return a, err
}

/**
 * @description: 查找全部文章
 * @param {int} page
 * @param {int} count
 */
func (m *Model) FindAll(page int, count int) (*[]Article, error) {
	var as *[]Article
	err := m.DB.Limit(count).Offset((page - 1) * count).Preload(clause.Associations).Find(&as).Error
	return as, err
}

/**
 * @description: 通过 用户ID 查找文章
 * @param {uint} userID
 * @param {int} page
 * @param {int} count
 */
func (m *Model) FindAllByUser(userID uint, page int, count int) (*[]Article, error) {
	var as *[]Article
	err := m.DB.Where("user_id = ?", userID).Limit(count).Offset((page - 1) * count).Find(as).Error
	return as, err
}

/**
 * @description: 通过标题查找文章
 * @param {string} title
 * @param {int} page
 * @param {int} count
 */
func (m *Model) FindAllByTitle(title string, page int, count int) (*Article, error) {
	var a *Article
	err := m.DB.Where("title like ?", "%"+title+"%").Limit(count).Offset((page - 1) * count).First(a).Error
	return a, err
}

/**
 * @description: 更新文章
 * @param {*Article} a
 */
func (m *Model) Update(a *Article) error {
	return m.DB.Updates(a).Error
}

/**
 * @description: 删除文章
 * @param {uint} id
 */
func (m *Model) Delete(id uint) error {
	return m.DB.Delete(&Article{}, id).Error
}
