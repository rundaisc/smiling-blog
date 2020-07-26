package dao

type SiteDao interface {
	Save()
	GetInfo()
}

type Site struct {
}

func NewSiteDao() SiteDao {
	return &Site{}
}

func (site *Site) Save() {

}

func (site *Site) GetInfo() {

}
