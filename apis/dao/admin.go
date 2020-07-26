package dao

import (
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type AdminDao interface {
	GetInfoByName(username string) *entity.Admin
	UpdateById(id int, data interface{}) error
}

type Admin struct {
	db databases.Db
}

func NewAdminDao() AdminDao {
	return &Admin{
		db: databases.Db{},
	}
}

// 通过名称获取用户
func (slf *Admin) GetInfoByName(username string) *entity.Admin {
	adminInfo := entity.Admin{}
	slf.db.DB().Where("username=?", username).First(&adminInfo)
	return &adminInfo
}

// 更新用户信息
func (slf *Admin) UpdateById(id int, data interface{}) error {
	adminInfo := entity.Admin{}
	return slf.db.DB().Model(adminInfo).Where("id=?", id).Updates(data).Error
}
