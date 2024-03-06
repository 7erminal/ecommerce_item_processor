package models

type CategoriesResponseDTO struct {
	StatusCode int            `orm: "omitempty"`
	Categories *[]interface{} `orm: "omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type CategoryResponseDTO struct {
	StatusCode int         `orm: "omitempty"`
	Categories *Categories `orm: "omitempty"`
	StatusDesc string      `orm:"size(255)"`
}
