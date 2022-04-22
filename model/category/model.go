/*
 * @Author: fzf404
 * @Date: 2022-02-27 09:58:58
 * @LastEditTime: 2022-02-27 10:43:21
 * @Description: 分类模型
 */
package category

import (
	"gover/global"

	"gorm.io/gorm"
)

type (
	CategoryModel interface {
		Insert(c *Category)       // 新增分类
		Update(c *Category) error // 更新分类
	}

	Model struct {
		DB *gorm.DB
	}

	Category struct {
		gorm.Model
		Name string `gorm:"comment:分类名称"`
		Desc string `gorm:"comment:分类描述"`
	}
)

func NewModel() CategoryModel {
	return &Model{
		DB: global.DB,
	}
}

/**
 * @description: 新增分类
 * @param {*Category} c
 */
func (m *Model) Insert(c *Category) {
	m.DB.Create(c)
}

/**
 * @description: 更新分类
 * @param {*Category} c
 */
func (m *Model) Update(c *Category) error {
	return m.DB.Updates(c).Error
}
