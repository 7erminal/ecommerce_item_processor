package responses

import (
	"time"
)

type FeaturesItemsDTO struct {
	FeatureId    string
	FeatureName  string
	ImagePath    string
	Visible      bool
	Description  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
	Items        []Items
}

type FeaturesResponseFDTO struct {
	StatusCode int
	Features   *[]FeaturesItemsDTO
	StatusDesc string
}
