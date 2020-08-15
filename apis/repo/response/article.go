package response

import "smiling-blog/entity"

type ArticleDetail struct {
	entity.Article
	CreatedAtFormat string   `json:"created_at"`
	UpdatedAtFormat string   `json:"updated_at"`
	DeleteAtFormat  string   `json:"deleted_at,omitempty"`
	CategoryName    string   `json:"category_name"`
	TagList         []string `json:"tag_list"`
}
