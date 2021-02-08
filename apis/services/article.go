package services

import (
	"errors"
	"smiling-blog/dao"
	"smiling-blog/databases"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/repo/response"
	"smiling-blog/tools/slog"
)

type ArticleServices interface {
	GetAdminList(params *request.ArticleSearchForm) (int, []response.ArticleDetail)
	Save(params *request.ArticleForm) error
	Delete(id int) error
	GetAdminDetail(id int) (response.ArticleDetail, error)
	GetFrontList(params *request.HomeForm) (int, []response.ArticleDetail)
	GetArticleNumberByCategory(categoryIds []int) map[int]int
}

//获取分类的文章数据
func (a *article) GetArticleNumberByCategory(categoryIds []int) map[int]int {
	return a.ArticleDao.GetArticleNumberByCategory(categoryIds)
}

//前台文章列表
func (a *article) GetFrontList(params *request.HomeForm) (int, []response.ArticleDetail) {
	searchParams := request.ArticleSearchForm{
		ArticleName: params.ArticleName,
		CategoryId:  params.CategoryId,
		IsShow:      1,
		IsRecycle:   0,
		IsDraft:     0,
		Page:        params.Page,
		Tag:         params.Tag,
		PageSize:    params.PageSize,
	}
	total := a.ArticleDao.GetRecordNumber(&searchParams)
	list := []response.ArticleDetail{}
	if total > 0 {
		fileds := "articles.id,article_name,category_id,view,comment,created_at,updated_at,category_name,description"
		list = a.ArticleDao.GetList(&searchParams, fileds)
		for i, v := range list {
			v.CreatedAtFormat = v.CreatedAt.Format("2006-01-02 15:04:05")
			v.UpdatedAtFormat = v.UpdatedAt.Format("2006-01-02 15:04:05")
			list[i] = v
		}
	}
	return total, list
}

type article struct {
	ArticleDao         dao.ArticleDao
	ArticleRelationDao dao.ArticleRelationDao
}

// 后台文章详情
func (a *article) GetAdminDetail(id int) (response.ArticleDetail, error) {
	detail := response.ArticleDetail{}
	article := a.ArticleDao.GetById(id, "*")
	if article.ID == 0 {
		return detail, errors.New("文章不存在")
	}
	detail.Article = *article
	taglist := a.ArticleRelationDao.GetTaglistById(id)
	outTgs := []string{}
	for _, tag := range taglist {
		outTgs = append(outTgs, tag.Tag)
	}
	detail.TagList = outTgs
	return detail, nil
}

// 保存文章
func (a *article) Save(params *request.ArticleForm) error {
	db := databases.Db{}
	trans := db.DB().Begin()
	art := &entity.Article{}
	if params.Id > 0 {
		art = a.ArticleDao.GetById(params.Id, "id")
		if art.ID == 0 {
			trans.Rollback()
			return errors.New("参数错误")
		}
		updateData := map[string]interface{}{
			"ArticleName": params.ArticleName,
			"Description": params.Description,
			"Cover":       params.Cover,
			"IsShow":      params.IsShow,
			"Content":     params.Content,
			"Sort":        params.Sort,
			"CategoryId":  params.CategoryId,
		}
		if err := a.ArticleDao.UpdateById(params.Id, updateData, trans); err != nil {
			trans.Rollback()
			slog.Error(err)
			return errors.New("保存文章失败")
		}
	} else {
		art.ArticleName = params.ArticleName
		art.Description = params.Description
		art.Cover = params.Cover
		art.IsShow = params.IsShow
		art.Content = params.Content
		art.Sort = params.Sort
		art.CategoryId = params.CategoryId
		if err := a.ArticleDao.Create(art, trans); err != nil {
			trans.Rollback()
			slog.Error(err)
			return errors.New("保存文章失败")
		}
	}

	//删除旧tag
	if err := a.ArticleRelationDao.DeleteByArticleId(art.ID, trans); err != nil {
		trans.Rollback()
		slog.Error(err)
		return errors.New("保存文章失败")
	}

	//添加 tag
	if err := a.ArticleRelationDao.Create(art.ID, params.TagList, trans); err != nil {
		trans.Rollback()
		slog.Error(err)
		return errors.New("保存文章失败")
	}

	trans.Commit()
	return nil

}

// 获取后台列表
func (a *article) GetAdminList(params *request.ArticleSearchForm) (int, []response.ArticleDetail) {
	total := a.ArticleDao.GetRecordNumber(params)
	list := []response.ArticleDetail{}
	if total > 0 {
		fileds := "articles.id,article_name,articles.is_show,category_id,created_at,updated_at,category_name,deleted_at"
		list = a.ArticleDao.GetList(params, fileds)
		for i, v := range list {
			v.CreatedAtFormat = v.CreatedAt.Format("2006-01-02 15:04:05")
			v.UpdatedAtFormat = v.UpdatedAt.Format("2006-01-02 15:04:05")
			if v.DeletedAt != nil {
				v.DeleteAtFormat = v.DeletedAt.Format("2006-01-02 15:04:05")
			}
			list[i] = v
		}
	}
	return total, list
}

// 删除文章
func (a *article) Delete(id int) error {
	//删除文章
	db := databases.Db{}
	trans := db.DB().Begin()
	err := a.ArticleDao.DeleteById(id, trans)
	if err != nil {
		trans.Rollback()
		slog.Error(err)
		return errors.New("删除文章失败")
	}
	//删除tag
	if err := a.ArticleRelationDao.DeleteByArticleId(id, trans); err != nil {
		trans.Rollback()
		slog.Error(err)
		return errors.New("删除文章失败")
	}
	trans.Commit()
	return nil
}

func NewArticleServices() ArticleServices {
	return &article{
		ArticleDao:         dao.NewArticleDao(),
		ArticleRelationDao: dao.NewArticleRelationDao(),
	}
}
