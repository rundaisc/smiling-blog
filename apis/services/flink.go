package services

import (
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
)

type FlinkServices interface {
	GetList() []entity.Flink
	Save(params *request.FlinkSave) tools.ResponseCode
	Delete(id int) tools.ResponseCode
}

type flink struct {
	FlinkDao dao.FlinkDao
}

func NewFlinkServices() FlinkServices {
	return &flink{
		FlinkDao: dao.NewFlinkDao(),
	}
}

// 所有导航列表
func (slf *flink) GetList() []entity.Flink {
	return slf.FlinkDao.GetList()
}

// 保存导航信息
func (slf *flink) Save(params *request.FlinkSave) tools.ResponseCode {
	link := &entity.Flink{}
	if params.Id > 0 {
		link = slf.FlinkDao.GetById(params.Id)
		if link.ID == 0 {
			return tools.FlinkNotFound
		}
	}
	link.LinkName = params.LinkName
	link.Url = params.Url
	link.Sort = params.Sort
	link.IsShow = params.IsShow

	if err := slf.FlinkDao.SaveLink(link); err != nil {
		return tools.FlinkSaveFailed
	}

	return tools.OK
}

// 删除导航
func (slf *flink) Delete(id int) tools.ResponseCode {
	err := slf.FlinkDao.DeleteById(id)
	if err != nil {
		return tools.UnKnowError
	}
	return tools.OK
}
