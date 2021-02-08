package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"smiling-blog/databases"
	"smiling-blog/entity"
)

type ArticleRelationDao interface {
	DeleteByArticleId(aid int, trans interface{}) error
	Create(aid int, tagList []string, trans interface{}) error
	GetTaglistById(id int) []entity.ArticleTagRelation
	FrontShow() []entity.FrontTag
}

type articleRelation struct {
	db databases.Db
}

// 获取文章tag列表
func (a articleRelation) GetTaglistById(id int) []entity.ArticleTagRelation {
	list := []entity.ArticleTagRelation{}
	a.db.DB().Table("article_tag_relations").Select("tag").Where("article_id=?", id).Scan(&list)
	return list
}

// 保存文章标签
func (a articleRelation) Create(aid int, tagList []string, trans interface{}) error {
	tagLen := len(tagList)
	if tagLen > 0 {
		sql := "insert into article_tag_relations (tag,article_id)  values "
		for i, tag := range tagList {
			sql += fmt.Sprintf("('%s',%d)", tag, aid)
			if i < tagLen-1 {
				sql += ","
			}
		}
		if db, ok := trans.(*gorm.DB); ok {
			return db.Exec(sql).Error
		} else {
			return a.db.DB().Exec(sql).Error
		}
	}
	return nil

}

// 删除文章tag
func (a articleRelation) DeleteByArticleId(aid int, trans interface{}) error {
	if db, ok := trans.(*gorm.DB); ok {
		return db.Where("article_id=?", aid).Delete(entity.ArticleTagRelation{}).Error
	}
	return a.db.DB().Where("article_id=?", aid).Delete(entity.ArticleTagRelation{}).Error
}

//前台展示标签
func (a articleRelation) FrontShow() []entity.FrontTag {
	list := []entity.FrontTag{}
	a.db.DB().Table("article_tag_relations").Joins("left join articles as a on a.id=article_tag_relations.article_id").
		Where("a.is_show=1 and a.deleted_at is null").Select("tag,count(*) as number").Group("tag").Scan(&list)
	return list
}

func NewArticleRelationDao() ArticleRelationDao {
	return &articleRelation{
		db: databases.Db{},
	}
}
