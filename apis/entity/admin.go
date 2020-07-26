package entity

import "time"

type Admin struct {
	ID            int
	Username      string `gorm:"size:30"`
	Password      string `gorm:"size:60"`
	Email         string `gorm:"size:60"`
	Phone         string `gorm:"size:16"`
	IsSuper       int    `gorm:"type:tinyint(2);default:0"`
	LastLoginTime time.Time
	LastLoginIp   string `gorm:"size:30"`
}

func (s *Admin) TableName() string {
	return "admins"
}
