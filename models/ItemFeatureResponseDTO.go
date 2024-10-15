package models

import (
	"time"
)

type Item_features2 struct {
	ItemFeatureId int64
	ItemId        *Items
	FeatureId     *Features
	Active        int
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     int
	ModifiedBy    int
}

type ItemFeatureResponseDTO struct {
	StatusCode  int
	ItemFeature *Item_features
	StatusDesc  string
}

type ItemFeaturesResponseDTO struct {
	StatusCode   int
	ItemFeatures *[]Item_features
	StatusDesc   string
}

type ItemFeaturesResponse2DTO struct {
	StatusCode   int
	ItemFeatures *[]Item_features2
	StatusDesc   string
}
