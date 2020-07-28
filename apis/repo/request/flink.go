package request

type FlinkSave struct {
	LinkName string `form:"link_name" json:"link_name" validate:"required"`
	Url      string `form:"url" json:"url" `
	Sort     int    `form:"sort" json:"sort" validate:"required"`
	IsShow   int    `form:"is_show" json:"is_show" validate:"required"`
	Id       int    `form:"id" json:"id" `
}
