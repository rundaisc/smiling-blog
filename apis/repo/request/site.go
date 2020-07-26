package request

type SiteForm struct {
	ID          int    `json:"id"`
	Title       string `gorm:"size:100" json:"title"`
	Keyword     string `gorm:"size:100" json:"keyword"`
	Description string `gorm:"size:100" json:"description"`
	Cover       string `gorm:"size:255" json:"cover"`
	Code        string `gorm:"text" json:"code"`
}
