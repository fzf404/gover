/*
 * @Author: fzf404
 * @Date: 2022-02-22 20:02:44
 * @LastEditTime: 2022-02-27 12:30:18
 * @Description: 用户模型
 */
package user

import (
	"gover/global"

	"gorm.io/gorm"
)

type (
	UserModel interface {
		Insert(u *User)                                // 新增用户, 通过 u.ID 判断是否成功
		Find(id uint) (*User, error)                   // 通过 ID 查询用户
		FindByEmail(email string) (*User, error)       // 通过邮箱查询用户
		FindByUserName(userName string) (*User, error) // 通过用户名查询用户
		Update(u *User) error                          // 更新用户
		Delete(id string) error                        // 删除用户
	}

	Model struct {
		DB *gorm.DB
	}

	User struct {
		gorm.Model
		UUID     string `gorm:"comment:用户唯一标识"`
		UserName string `gorm:"comment:用户登录名"`
		NickName string `gorm:"comment:用户昵称"`
		Password string `gorm:"comment:用户密码"`
		Type     uint   `gorm:"default:1;comment:用户角色"` // 0 管理员 1 普通用户
		Email    string `gorm:"comment:用户邮箱"`
		Avatar   string `gorm:"default:'/static/uploads/avatar.jpg';comment:用户头像"`
	}
)

func NewModel() UserModel {
	return &Model{
		DB: global.DB,
	}
}

/**
 * @description: 新增用户
 * @param {*User} u
 */
func (m *Model) Insert(u *User) {
	m.DB.Create(u)
}

/**
 * @description: 通过 ID 查找用户
 * @param {string} id
 */
func (m *Model) Find(id uint) (*User, error) {
	u := &User{}
	err := m.DB.First(u, id).Error
	return u, err
}

/**
 * @description: 通过邮箱查找用户
 * @param {string} email
 */
func (m *Model) FindByEmail(email string) (*User, error) {
	u := &User{}
	err := m.DB.Where("email = ?", email).First(u).Error
	return u, err
}

/**
 * @description: 通过用户名查找用户
 * @param {string} name
 */
func (m *Model) FindByUserName(userName string) (*User, error) {
	u := &User{}
	err := m.DB.Where("user_name = ?", userName).First(u).Error
	return u, err
}

/**
 * @description: 更新用户信息
 * @param {*User} u
 */
func (m *Model) Update(u *User) error {
	return m.DB.Updates(u).Error
}

/**
 * @description: 删除用户
 * @param {string} id
 */
func (m *Model) Delete(id string) error {
	return m.DB.Where("id = ?", id).Delete(User{}).Error
}
