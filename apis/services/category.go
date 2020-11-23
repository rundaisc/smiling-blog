package services

import (
	"smiling-blog/dao"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/tools"
	"smiling-blog/tools/slog"
)

type CategoryServices interface {
	GetList() []entity.Category
	FrontList() []entity.Category
	Save(params *request.CategorySave) tools.ResponseCode
	Delete(id int) tools.ResponseCode
}

type category struct {
	categoryDao dao.CategoryDao
}

func NewCategoryServices() CategoryServices {
	return &category{
		categoryDao: dao.NewCategoryDao(),
	}
}

// 所有列表
func (slf *category) FrontList() []entity.Category {
	return slf.categoryDao.GetList()
}

// 所有列表
func (slf *category) GetList() []entity.Category {
	return slf.categoryDao.GetList()
}

// 保存信息
func (slf *category) Save(params *request.CategorySave) tools.ResponseCode {
	category := &entity.Category{}
	if params.Id > 0 {
		category = slf.categoryDao.GetById(params.Id)
		if category.ID == 0 {
			return tools.CategoryNotFound
		}
		updateData := map[string]interface{}{
			"is_show":       params.IsShow,
			"sort":          params.Sort,
			"category_name": params.CategoryName,
		}
		if err := slf.categoryDao.UpdateCategory(category, updateData); err != nil {
			slog.Log.Error(err)
			return tools.CategorySaveFailed
		}
	} else {
		category.IsShow = params.IsShow
		category.Sort = params.Sort
		category.CategoryName = params.CategoryName
		if err := slf.categoryDao.CreateCategory(category); err != nil {
			slog.Log.Error(err)
			return tools.CategorySaveFailed
		}
	}

	return tools.OK
}

// 删除
func (slf *category) Delete(id int) tools.ResponseCode {
	err := slf.categoryDao.DeleteById(id)
	if err != nil {
		slog.Log.Error(err)
		return tools.UnKnowError
	}
	return tools.OK
}
