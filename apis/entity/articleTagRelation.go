package entity

type ArticleTagRelation struct {
	ID        int    `json:"id"`
	Tag       string `json:"tag"`
	ArticleId int    `json:"article_id"`
}

type FrontTag struct {
	Tag    string `json:"tag"`
	Number int    `json:"number"`
}

func (a *ArticleTagRelation) TableName() string {
	return "article_tag_relations"
}
