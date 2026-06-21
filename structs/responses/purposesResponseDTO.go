package responses

import "item_processor/models"

type PurposesResponseDTO struct {
	StatusCode int                `orm: "omitempty"`
	Purposes   *[]models.Purposes `orm: "omitempty"`
	StatusDesc string             `orm:"size(255)"`
}

type PurposeResponseDTO struct {
	StatusCode int              `orm: "omitempty"`
	Purpose    *models.Purposes `orm: "omitempty"`
	StatusDesc string           `orm:"size(255)"`
}
