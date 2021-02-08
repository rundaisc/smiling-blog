package services

import (
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type FlinkServices interface {
	GetList() []entity.Flink
	GetFrontList() []entity.Flink
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

// 所有链接列表
func (slf *flink) GetList() []entity.Flink {
	return slf.FlinkDao.GetList()
}

// 前台链接列表
func (slf *flink) GetFrontList() []entity.Flink {
	return slf.FlinkDao.GetList()
}

// 保存链接信息
func (slf *flink) Save(params *request.FlinkSave) tools.ResponseCode {
	link := &entity.Flink{}
	if params.Id > 0 {
		link = slf.FlinkDao.GetById(params.Id)
		if link.ID == 0 {
			return tools.FlinkNotFound
		}
		update := map[string]interface{}{
			"link_name": params.LinkName,
			"url":       params.Url,
			"sort":      params.Sort,
			"is_show":   params.IsShow,
		}
		if err := slf.FlinkDao.UpdateLink(link, update); err != nil {
			slog.Error(err)
			return tools.FlinkSaveFailed
		}

	} else {
		link.LinkName = params.LinkName
		link.Url = params.Url
		link.Sort = params.Sort
		link.IsShow = params.IsShow
		if err := slf.FlinkDao.SaveLink(link); err != nil {
			slog.Error(err)
			return tools.FlinkSaveFailed
		}
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
