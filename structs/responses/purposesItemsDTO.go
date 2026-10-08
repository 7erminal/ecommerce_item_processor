package responses

import (
	"time"
)

type PurposesItemsDTO struct {
	PurposeId    string
	Purpose      string
	ImagePath    string
	Visible      bool
	Description  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	Items        []Items
}

type PurposesResponseFDTO struct {
	StatusCode int
	Purposes   *[]PurposesItemsDTO
	StatusDesc string
}
