package entity

import "time"

type Article struct {
	ID          int        `json:"id"`
	ArticleName string     `gorm:"size:30" json:"article_name"`
	Description string     `gorm:"size:255" json:"description"`
	Sort        int        `gorm:"type:int(5);default:99" json:"sort"`
	IsShow      int        `gorm:"type:tinyint(2);default:1" json:"is_show"`
	Cover       string     `gorm:"size:200" json:"cover,omitempty"`
	Content     string     `gorm:"type:text" json:"content,omitempty"`
	CategoryId  int        `json:"category_id"`
	View        int        `json:"view"`
	Comment     int        `json:"comment"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at,omitempty"`
}

func (a *Article) TableName() string {
	return "articles"
}
