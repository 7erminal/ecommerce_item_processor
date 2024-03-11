package models

import (
	"time"
)

type Item_features2 struct {
	ItemFeatureId int64 `orm:"auto"`
	ItemId        *Items
	FeatureId     *Features
	Active        int
	DateCreated   time.Time `orm:"type(datetime)"`
	DateModified  time.Time `orm:"type(datetime)"`
	CreatedBy     int
	ModifiedBy    int
}

type ItemFeatureResponseDTO struct {
	StatusCode  int            `orm: "omitempty"`
	ItemFeature *Item_features `orm: "omitempty"`
	StatusDesc  string         `orm:"size(255)"`
}

type ItemFeaturesResponseDTO struct {
	StatusCode   int              `orm: "omitempty"`
	ItemFeatures *[]Item_features `orm: "omitempty"`
	StatusDesc   string           `orm:"size(255)"`
}

type ItemFeaturesResponse2DTO struct {
	StatusCode   int               `orm: "omitempty"`
	ItemFeatures *[]Item_features2 `orm: "omitempty"`
	StatusDesc   string            `orm:"size(255)"`
}
