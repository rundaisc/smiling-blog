package dao

import (
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type SiteDao interface {
	Save(info *entity.Site) error
	GetInfo() entity.Site
}

type Site struct {
	db databases.Db
}

func NewSiteDao() SiteDao {
	return &Site{
		db: databases.Db{},
	}
}

// 保存site信息
func (sd *Site) Save(info *entity.Site) error {
	return sd.db.DB().Save(info).Error

}

// 获取site信息
func (sd *Site) GetInfo() entity.Site {
	siteInfo := entity.Site{}
	sd.db.DB().First(&siteInfo)
	return siteInfo
}
