package responses

import (
	"item_processor/models"
	"time"
)

type PurposesItemsDTO struct {
	PurposeId    int64     `orm:"auto"`
	Purpose      string    `orm:"size(40)"`
	ImagePath    string    `orm:"size(250)"`
	Visible      bool      `orm:"omitempty"`
	Description  string    `orm:"size(250);omitempty"`
	Active       int       `orm:"omitempty"`
	DateCreated  time.Time `orm:"type(datetime);omitempty"`
	DateModified time.Time `orm:"type(datetime);omitempty"`
	CreatedBy    int       `orm:"omitempty"`
	ModifiedBy   int       `orm:"omitempty"`
	Items        []models.Items
}

type PurposesResponseFDTO struct {
	StatusCode int
	Purposes   *[]PurposesItemsDTO
	StatusDesc string
}
