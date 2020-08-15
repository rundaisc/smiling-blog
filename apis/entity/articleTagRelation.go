package entity

type ArticleTagRelation struct {
	ID        int    `json:"id"`
	Tag       string `json:"tag"`
	ArticleId int    `json:"article_id"`
}

func (a *ArticleTagRelation) TableName() string {
	return "article_tag_relations"
}
