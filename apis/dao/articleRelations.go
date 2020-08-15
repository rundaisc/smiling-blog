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
}

type articleRelation struct {
	db databases.Db
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

func NewArticleRelationDao() ArticleRelationDao {
	return &articleRelation{
		db: databases.Db{},
	}
}
