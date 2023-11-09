package models

type ErrorResponse struct {
	StatusCode int    `orm:"omitempty"`
	Error      string `orm:"omitempty"`
	StatusDesc string `orm:"size(255)" orm:"omitempty"`
}
