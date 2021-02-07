package response

import "smiling-blog/entity"

type Public struct {
	SiteInfo     entity.Site            `json:"site_info"`
	NavList      []entity.FontNavList   `json:"nav_list"`
	CategoryList []entity.FrontCategory `json:"category_list"`
}
