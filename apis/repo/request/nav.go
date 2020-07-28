package request

type NavSave struct {
	NavName    string `form:"nav_name" json:"nav_name" validate:"required"`
	CategoryId int    `form:"category_id" json:"category_id" `
	Url        string `form:"url" json:"url" `
	Sort       int    `form:"sort" json:"sort" validate:"required"`
	IsShow     int    `form:"is_show" json:"is_show" validate:"required"`
	Id         int    `form:"id" json:"id" `
}
