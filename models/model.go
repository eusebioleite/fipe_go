package models

type Model struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Fipe        string `json:"fipe"`
	Brand_id    int    `json:"brand_id"`
}
