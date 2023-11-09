package models

type ItemsDTO struct {
	ItemName        string `orm:"size(80)"`
	Description     string `orm:"size(250)" orm:"omitempty"`
	Category        int
	AvailableSizes  []string `orm:"size(250)" orm:"omitempty"`
	AvailableColors []string `orm:"size(250)" orm:"omitempty"`
	Quantity        int
	ItemPrice       float32
	CreatedBy       int `orm:"omitempty"`
}
