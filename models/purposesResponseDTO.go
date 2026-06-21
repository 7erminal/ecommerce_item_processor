package models

type PurposesResponseDTO struct {
	StatusCode int         `orm: "omitempty"`
	Purposes   *[]Purposes `orm: "omitempty"`
	StatusDesc string      `orm:"size(255)"`
}

type PurposeResponseDTO struct {
	StatusCode int       `orm: "omitempty"`
	Purpose    *Purposes `orm: "omitempty"`
	StatusDesc string    `orm:"size(255)"`
}
