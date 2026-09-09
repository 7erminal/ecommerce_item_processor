package models

type ItemImagesResponseDTO struct {
	StatusCode int
	ItemImages *[]interface{}
	StatusDesc string
}

type ItemImagesResponseDTO2 struct {
	StatusCode int
	ItemImages *[]Item_images
	StatusDesc string
}

type ItemImageResponseDTO struct {
	StatusCode int
	ItemImage  *Item_images
	StatusDesc string
}
