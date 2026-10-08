package responses

import "time"

type Features struct {
	FeatureId    string
	Feature      string
	ImagePath    string
	Visible      bool
	Description  string
	Active       int
	DateCreated  time.Time
	DateModified time.Time
}

type FeaturesResponseDTO struct {
	StatusCode int
	Features   *[]Features
	StatusDesc string
}

type FeatureResponseDTO struct {
	StatusCode int
	Feature    *Features
	StatusDesc string
}
