package response

import "smiling-blog/entity"

type NavRow struct {
	entity.Nav
	Active string `json:"active"`
}
