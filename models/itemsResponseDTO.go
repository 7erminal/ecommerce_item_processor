package models

type ItemsResponseDTO struct {
	StatusCode int    `orm:"omitempty"`
	Item       *Items `orm:"omitempty"`
	StatusDesc string `orm:"size(255)"`
}
