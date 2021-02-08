package response

import "smiling-blog/entity"

type Public struct {
	SiteInfo     entity.Site            `json:"site_info"`
	NavList      []entity.FontNavList   `json:"nav_list"`
	CategoryList []entity.FrontCategory `json:"category_list"`
	TagList      []entity.FrontTag      `json:"tag_list"`
	FlinkList    []entity.Flink         `json:"flink_list"`
}
