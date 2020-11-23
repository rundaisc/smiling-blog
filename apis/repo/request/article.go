package request

type ArticleSearchForm struct {
	ArticleName string `json:"article_name" form:"article_name"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	IsShow      int    `json:"is_show" form:"is_show"`
	IsRecycle   int    `json:"is_recycle" form:"is_recycle"`
	IsDraft     int    `json:"is_draft" form:"is_draft"`
	Page        int    `json:"page" form:"page"`
	Tag         string `json:"tag"`
	PageSize    int    `json:"page_size" form:"page_size"`
}

type ArticleForm struct {
	Id           int      `json:"id"`
	ArticleName  string   ` json:"article_name" validate:"required"`
	Description  string   ` json:"description" validate:"required"`
	Sort         int      ` json:"sort" validate:"required"`
	IsShow       int      `json:"is_show" validate:"required"`
	CategoryId   int      `json:"category_id" validate:"required"`
	Cover        string   `json:"cover"`
	Content      string   `json:"content" validate:"required"`
	TagList      []string `json:"tag_list" `
	CategoryList []int    `json:"category_list" `
}
