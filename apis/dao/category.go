package dao

import (
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type CategoryDao interface {
	GetList() []entity.Category
	GetById(id int) *entity.Category
	CreateCategory(category *entity.Category) error
	DeleteById(id int) error
	UpdateCategory(category *entity.Category, updateData map[string]interface{}) error
}

type category struct {
	db databases.Db
}

func (slf *category) UpdateCategory(category *entity.Category, updateData map[string]interface{}) error {
	return slf.db.DB().Model(category).Updates(updateData).Error

}

func NewCategoryDao() CategoryDao {
	return &category{
		db: databases.Db{},
	}
}

// 根据id 获取
func (slf *category) GetById(id int) *entity.Category {
	category := entity.Category{ID: id}
	slf.db.DB().First(&category)
	return &category
}

// 保存
func (slf *category) CreateCategory(category *entity.Category) error {
	return slf.db.DB().Save(category).Error
}

// 删除
func (slf *category) DeleteById(id int) error {
	category := entity.Category{}
	return slf.db.DB().Where("id=?", id).Delete(&category).Error
}

// 导航列表
func (slf *category) GetList() []entity.Category {
	list := []entity.Category{}
	category := entity.Category{}
	slf.db.DB().Model(&category).Scan(&list)
	return list
}
