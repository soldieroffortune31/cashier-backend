package model

type Permission struct {
	Id       string `json:"id"`
	Resource string `json:"resource"`
	Action   string `json:"action"`
}
