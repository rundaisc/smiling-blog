package services

import (
	"fmt"
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
)

type NavServices interface {
	GetList() []entity.Nav
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
func (slf *nav) GetList() []entity.Nav {
	return slf.NavDao.GetList()
}

// 保存导航信息
func (slf *nav) Save(params *request.NavSave) tools.ResponseCode {
	nav := &entity.Nav{}
	if params.Id > 0 {
		nav = slf.NavDao.GetById(params.Id)
		if nav.ID == 0 {
			return tools.NavNotFound
		}
	}
	nav.IsShow = params.IsShow
	nav.Sort = params.Sort
	nav.CategoryId = params.CategoryId
	nav.NavName = params.NavName

	//如果是分类导航
	if params.CategoryId > 0 {
		nav.Url = fmt.Sprintf("/category/%d", params.CategoryId)
	} else {
		nav.Url = params.Url
	}

	if err := slf.NavDao.SaveNav(nav); err != nil {
		return tools.NavSaveFailed
	}

	return tools.OK
}

// 删除导航
func (slf *nav) Delete(id int) tools.ResponseCode {
	err := slf.NavDao.DeleteById(id)
	if err != nil {
		return tools.UnKnowError
	}
	return tools.OK
}
