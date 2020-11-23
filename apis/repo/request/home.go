package request

type HomeForm struct {
	ArticleName string `json:"article_name" form:"article_name"`
	CategoryId  int    `json:"category_id" form:"category_id"`
	Tag         string `json:"tag" form:"tag"`
	Page        int    `json:"page" form:"page"`
	PageSize    int    `json:"page_size" form:"page_size"`
}
