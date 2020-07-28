package entity

type Flink struct {
	ID       int    `json:"id"`
	LinkName string `gorm:"size:30" json:"link_name"`
	Url      string `gorm:"size:100" json:"url"`
	Sort     int    `gorm:"type:tinyint(4);default:99" json:"sort"`
	IsShow   int    `gorm:"type:tinyint(2);default:1" json:"is_show"`
}

func (l *Flink) TableName() string {
	return "flinks"
}
