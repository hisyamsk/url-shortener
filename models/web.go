package models

type WebResponse struct {
	Success bool  `json:"success"`
	Error   error `json:"error"`
	Data    any   `json:"data"`
}
