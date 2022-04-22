/*
 * @Author: fzf404
 * @Date: 2022-02-27 12:54:21
 * @LastEditTime: 2022-02-28 11:42:01
 * @Description: 标签模型
 */
package tag

import (
	"gover/global"

	"gorm.io/gorm"
)

type (
	TagModel interface {
		Insert(c *Tag)
		FindByName(name string) (*Tag, error)
		Update(c *Tag) error
	}

	Model struct {
		DB *gorm.DB
	}

	Tag struct {
		gorm.Model
		Name string `gorm:"comment:分类名称"`
	}
)

func NewModel() TagModel {
	return &Model{
		DB: global.DB,
	}
}

/**
 * @description: 新增标签
 * @param {*Tag} a
 */
func (m *Model) Insert(a *Tag) {
	m.DB.Create(a)
}

/**
 * @description: 通过名称查找标签
 * @param {string} name
 */
func (m *Model) FindByName(name string) (*Tag, error) {
	var tag Tag
	err := m.DB.Where("name = ?", name).First(&tag).Error
	return &tag, err
}

/**
 * @description: 更新标签
 * @param {*Tag} c
 */
func (m *Model) Update(c *Tag) error {
	return m.DB.Updates(c).Error
}
