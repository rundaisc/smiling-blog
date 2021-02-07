package services

import (
	"fmt"
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type NavServices interface {
	GetList() []entity.Nav
	GetFrontList(path string) []entity.FontNavList
	Save(params *request.NavSave) tools.ResponseCode
	Delete(id int) tools.ResponseCode
}

type nav struct {
	NavDao dao.NavDao
}

func NewNavServices() NavServices {
	return &nav{
		NavDao: dao.NewNavDao(),
	}
}

// 所有导航列表
func (slf *nav) GetFrontList(path string) []entity.FontNavList {
	list := slf.NavDao.GetFrontList()
	for i, v := range list {
		if v.Url == path {
			v.Active = "active"
		}
		list[i] = v
	}
	return list
}

// 所有导航列表
func (slf *nav) GetList() []entity.Nav {
	return slf.NavDao.GetList()
}

// 保存导航信息
func (slf *nav) Save(params *request.NavSave) tools.ResponseCode {
	nav := &entity.Nav{}

	//如果是分类导航
	if params.CategoryId > 0 {
		params.Url = fmt.Sprintf("/category/%d", params.CategoryId)
	}
	if params.Id > 0 {
		nav = slf.NavDao.GetById(params.Id)
		if nav.ID == 0 {
			return tools.NavNotFound
		}
		updateData := map[string]interface{}{
			"is_show":     params.IsShow,
			"sort":        params.Sort,
			"category_id": params.CategoryId,
			"nav_name":    params.NavName,
			"url":         params.Url,
		}
		if err := slf.NavDao.UpdateNav(nav, updateData); err != nil {
			slog.Error(err)
			return tools.NavSaveFailed
		}
	} else {
		nav.IsShow = params.IsShow
		nav.Sort = params.Sort
		nav.CategoryId = params.CategoryId
		nav.NavName = params.NavName
		nav.Url = params.Url
		if err := slf.NavDao.CreateNav(nav); err != nil {
			slog.Error(err)
			return tools.NavSaveFailed
		}
	}

	return tools.OK
}

// 删除导航
func (slf *nav) Delete(id int) tools.ResponseCode {
	err := slf.NavDao.DeleteById(id)
	if err != nil {
		slog.Error(err)
		return tools.UnKnowError
	}
	return tools.OK
}
