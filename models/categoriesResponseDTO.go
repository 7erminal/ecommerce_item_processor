package models

type CategoriesResponseDTO struct {
	StatusCode int
	Categories *[]Categories
	StatusDesc string
}

type CategoryResponseDTO struct {
	StatusCode int
	Category   *Categories
	StatusDesc string
}
