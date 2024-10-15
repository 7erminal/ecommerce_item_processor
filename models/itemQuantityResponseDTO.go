package models

type ItemQuantityResponseDTO struct {
	StatusCode int            `orm: "omitempty"`
	Quantity   *Item_quantity `orm: "omitempty"`
	StatusDesc string         `orm:"size(255)"`
}
