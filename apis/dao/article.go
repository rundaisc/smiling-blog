package dao

import (
	"github.com/jinzhu/gorm"
	"smiling-blog/databases"
	"smiling-blog/entity"
	"smiling-blog/repo/request"
	"smiling-blog/repo/response"
)

type ArticleDao interface {
	GetRecordNumber(params *request.ArticleSearchForm) int
	GetList(params *request.ArticleSearchForm) []response.ArticleDetail
	GetById(id int, fields string) *entity.Article
	UpdateById(id int, data map[string]interface{}, trans interface{}) error
	Create(art *entity.Article, trans interface{}) error
	DeleteById(id int, trans interface{}) error
}

type article struct {
	db databases.Db
}

// 创建文章
func (a *article) Create(art *entity.Article, trans interface{}) error {
	if db, ok := trans.(*gorm.DB); ok {
		return db.Save(art).Error
	}
	return a.db.DB().Save(art).Error

}

// 根据Id修改文章
func (a *article) UpdateById(id int, data map[string]interface{}, trans interface{}) error {
	if db, ok := trans.(*gorm.DB); ok {
		return db.Model(entity.Article{}).Where("id=?", id).Updates(data).Error
	}
	return a.db.DB().Model(entity.Article{}).Where("id=?", id).Updates(data).Error
}

// 根据id 获取详情
func (a *article) GetById(id int, fields string) *entity.Article {
	art := &entity.Article{}
	art.ID = id
	a.db.DB().Select(fields).First(art)
	return art
}

// 获取列表
func (a *article) GetList(params *request.ArticleSearchForm) []response.ArticleDetail {
	list := []response.ArticleDetail{}
	fileds := "articles.id,article_name,articles.is_show,category_id,created_at,updated_at,category_name,deleted_at"
	a.getQueryBuild(params).Model(entity.Article{}).Select(fileds).Joins("left join categories on articles.category_id=categories.id").Limit(params.PageSize).Offset(params.PageSize * (params.Page - 1)).Order("created_at desc").Scan(&list)
	return list
}

// 获取记录总数
func (a *article) GetRecordNumber(params *request.ArticleSearchForm) int {
	total := 0
	query := a.getQueryBuild(params)
	query.Model(entity.Article{}).Count(&total)
	return total
}

// 获取查询链式
func (a *article) getQueryBuild(params *request.ArticleSearchForm) *gorm.DB {
	query := a.db.DB()
	if params.CategoryId > 0 {
		query = query.Where("category_id = ?", params.CategoryId)
	}
	if params.IsShow > 0 {
		query = query.Where("articles.is_show = ?", params.IsShow)
	}
	if params.ArticleName != "" {
		query = query.Where("article_name like ?", params.ArticleName+"%")
	}
	//回收站
	if params.IsRecycle > 0 {
		query = query.Unscoped().Where("articles.deleted_at is not null ")
	}
	return query
}

// 删除文章
func (a *article) DeleteById(id int, trans interface{}) error {
	if db, ok := trans.(*gorm.DB); ok {
		return db.Where("id=?", id).Delete(entity.Article{}).Error
	}
	return a.db.DB().Where("id=?", id).Delete(entity.Article{}).Error
}

func NewArticleDao() ArticleDao {
	return &article{
		db: databases.Db{},
	}
}
