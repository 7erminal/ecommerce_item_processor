package responses

type ItemImagesResponseDTO struct {
	StatusCode int
	ItemImages *[]interface{}
	StatusDesc string
}

type ItemImagesResponseDTO2 struct {
	StatusCode int
	ItemImages *[]ItemImagesDTO
	StatusDesc string
}

type ItemImageResponseDTO struct {
	StatusCode int
	ItemImage  *ItemImagesDTO
	StatusDesc string
}
