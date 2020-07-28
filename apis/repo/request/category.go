package request

type CategorySave struct {
	CategoryName string `form:"category_name" json:"category_name" validate:"required"`
	Sort         int    `form:"sort" json:"sort" validate:"required"`
	IsShow       int    `form:"is_show" json:"is_show" validate:"required"`
	Id           int    `form:"id" json:"id"`
}
