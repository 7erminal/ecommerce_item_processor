package responses

import (
	"item_processor/models"
)

type ItemsResponseDTO struct {
	StatusCode int
	Items      *[]interface{}
	StatusDesc string
}

type ItemResponseDTO struct {
	StatusCode int
	Item       *models.Items
	StatusDesc string
}
