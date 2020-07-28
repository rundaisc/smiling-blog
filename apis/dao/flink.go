package dao

import (
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type FlinkDao interface {
	GetList() []entity.Flink
	GetById(id int) *entity.Flink
	SaveLink(flink *entity.Flink) error
	DeleteById(id int) error
}

type flink struct {
	db databases.Db
}

func NewFlinkDao() FlinkDao {
	return &flink{
		db: databases.Db{},
	}
}

// 根据id 获取Link
func (slf *flink) GetById(id int) *entity.Flink {
	link := entity.Flink{}
	slf.db.DB().First(&link)
	return &link
}

// 保存链接
func (slf *flink) SaveLink(flink *entity.Flink) error {
	return slf.db.DB().Save(flink).Error
}

// 删除链接
func (slf *flink) DeleteById(id int) error {
	link := entity.Flink{}
	return slf.db.DB().Where("id=?", id).Delete(&link).Error
}

// 链接列表
func (slf *flink) GetList() []entity.Flink {
	list := []entity.Flink{}
	flink := entity.Flink{}
	slf.db.DB().Model(&flink).Scan(&list)
	return list
}
