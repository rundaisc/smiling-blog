package entity

type Nav struct {
	ID         int    `json:"id"`
	NavName    string `gorm:"size:30" json:"nav_name"`
	Url        string `gorm:"size:100" json:"url"`
	CategoryId int    `gorm:"size:30" json:"category_id"`
	Sort       int    `gorm:"type:tinyint(4);default:99" json:"sort"`
	IsShow     int    `gorm:"type:tinyint(2);default:1" json:"is_show"`
}

type FontNavList struct {
	Nav
	Active string `json:"active"`
}

func (n *Nav) TableName() string {
	return "navs"
}
