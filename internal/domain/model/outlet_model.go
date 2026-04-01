package model

type Outlet struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	IsHeadOffice bool   `json:"is_head_office"`
}
