package services

import (
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
)

type SiteServices interface {
	Save(params *request.SiteForm) error
	GetInfo() *entity.Site
}

type Site struct {
	SiteDao dao.SiteDao
}

func NewSiteServices() SiteServices {
	return &Site{
		SiteDao: dao.NewSiteDao(),
	}
}

// 保存site信息
func (site *Site) Save(params *request.SiteForm) error {
	siteInfo := site.SiteDao.GetInfo()
	siteInfo.Keyword = params.Keyword
	siteInfo.Title = params.Title
	siteInfo.Code = params.Code
	siteInfo.Cover = params.Cover
	siteInfo.Description = params.Description
	return site.SiteDao.Save(siteInfo)
}

// 获取site信息
func (site *Site) GetInfo() *entity.Site {
	return site.SiteDao.GetInfo()
}
