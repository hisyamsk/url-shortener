package models

type UrlModel struct {
	ID       uint   `json:"id"`
	Url      string `json:"url"`
	Redirect string `json:"redirect"`
	UserID   uint   `json:"userId"`
}

type UrlCreateRequest struct {
	Url      string `json:"url" validate:"omitempty,alphanum,min=6,max=20"`
	Redirect string `json:"redirect" validate:"required,url"`
	UserID   uint   `json:"userId" validate:"required"`
}

type UrlUpdateRequest struct {
	Url      string `json:"url" validate:"omitempty,alphanum,min=6,max=20"`
	Redirect string `json:"redirect" validate:"omitempty,url"`
	UserID   uint   `json:"userId" validate:"required"`
}
