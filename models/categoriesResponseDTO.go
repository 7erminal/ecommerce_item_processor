package models

type CategoriesResponseDTO struct {
	StatusCode int
	Categories *[]interface{}
	StatusDesc string
}

type CategoryResponseDTO struct {
	StatusCode int
	Categories *Categories
	StatusDesc string
}
