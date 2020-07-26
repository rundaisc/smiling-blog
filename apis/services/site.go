package services

type SiteServices interface {
	Save()
}

type Site struct {
}

func NewSiteServices() SiteServices {
	return &Site{}
}

func (site *Site) Save() {

}
