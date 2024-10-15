package models

type ItemPurposeResponseDTO struct {
	StatusCode  int
	ItemPurpose *Item_purposes
	StatusDesc  string
}

type ItemPurposesResponseDTO struct {
	StatusCode   int
	ItemPurposes *[]Item_purposes
	StatusDesc   string
}
