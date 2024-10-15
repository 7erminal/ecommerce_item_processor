package models

type ItemsResponseDTO struct {
	StatusCode int            `orm:"omitempty"`
	Items      *[]interface{} `orm:"omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type ItemsResponseDTO2 struct {
	StatusCode int      `orm:"omitempty"`
	Items      *[]Items `orm:"omitempty"`
	StatusDesc string   `orm:"size(255)"`
}

type ItemResponseDTO struct {
	StatusCode int    `orm:"omitempty"`
	Item       *Items `orm:"omitempty"`
	StatusDesc string `orm:"size(255)"`
}
