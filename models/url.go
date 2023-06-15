package models

type UrlModel struct {
	ID       uint   `json:"id"`
	Url      string `json:"url"`
	Redirect string `json:"redirect"`
	UserID   uint   `json:"userId"`
}
