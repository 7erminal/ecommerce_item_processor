package models

type ItemImagesResponseDTO struct {
	StatusCode int            `orm:"omitempty"`
	ItemImages *[]interface{} `orm:"omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type ItemImagesResponseDTO2 struct {
	StatusCode int            `orm:"omitempty"`
	ItemImages *[]Item_images `orm:"omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type ItemImageResponseDTO struct {
	StatusCode int          `orm:"omitempty"`
	ItemImage  *Item_images `orm:"omitempty"`
	StatusDesc string       `orm:"size(255)"`
}
