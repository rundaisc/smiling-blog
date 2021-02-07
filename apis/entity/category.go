package entity

type Category struct {
	ID           int    `json:"id"`
	CategoryName string `gorm:"size:30" json:"category_name"`
	Sort         int    `gorm:"type:tinyint(4);default:99" json:"sort"`
	IsShow       int    `gorm:"type:tinyint(2);default:1" json:"is_show"`
}

type FrontCategory struct {
	ID           int    `json:"id"`
	CategoryName string `json:"category_name"`
	Number       int    `json:"number"`
}

func (n *Category) TableName() string {
	return "categories"
}
