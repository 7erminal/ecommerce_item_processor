package models

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
