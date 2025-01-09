package models

type ItemsDTO struct {
	ItemName        string `orm:"size(80)"`
	Description     string `orm:"size(250)"`
	Category        int
	AvailableSizes  []string `orm:"size(250)"`
	AvailableColors []string `orm:"size(250)"`
	Quantity        int
	ItemPrice       float32
	CreatedBy       int `orm:"omitempty"`
}
