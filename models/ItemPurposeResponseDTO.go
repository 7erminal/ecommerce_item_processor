package models

type ItemPurposeResponseDTO struct {
	StatusCode  int            `orm: "omitempty"`
	ItemPurpose *Item_purposes `orm: "omitempty"`
	StatusDesc  string         `orm:"size(255)"`
}

type ItemPurposesResponseDTO struct {
	StatusCode   int              `orm: "omitempty"`
	ItemPurposes *[]Item_purposes `orm: "omitempty"`
	StatusDesc   string           `orm:"size(255)"`
}
