package models

type FeaturesResponseDTO struct {
	StatusCode int    `orm: "omitempty"`
	Features       *[]interface{} `orm: "omitempty"`
	StatusDesc string `orm:"size(255)"`
}

type FeatureResponseDTO struct {
	StatusCode int    `orm: "omitempty"`
	Feature       *Features `orm: "omitempty"` 
	StatusDesc string `orm:"size(255)"`
}