package dao

import (
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type NavDao interface {
	GetList() []entity.Nav
	GetById(id int) *entity.Nav
	SaveNav(nav *entity.Nav) error
	DeleteById(id int) error
}

type nav struct {
	db databases.Db
}

func NewNavDao() NavDao {
	return &nav{
		db: databases.Db{},
	}
}

// 根据id 获取
func (slf *nav) GetById(id int) *entity.Nav {
	nav := entity.Nav{}
	slf.db.DB().First(&nav)
	return &nav
}

// 保存
func (slf *nav) SaveNav(nav *entity.Nav) error {
	return slf.db.DB().Save(nav).Error
}

// 删除
func (slf *nav) DeleteById(id int) error {
	nav := entity.Nav{}
	return slf.db.DB().Where("id=?", id).Delete(&nav).Error
}

// 导航列表
func (slf *nav) GetList() []entity.Nav {
	list := []entity.Nav{}
	nav := entity.Nav{}
	slf.db.DB().Model(&nav).Scan(&list)
	return list
}
