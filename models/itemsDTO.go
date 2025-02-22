package models

type ItemsDTO struct {
	ItemName        string `orm:"size(80)"`
	Description     string `orm:"size(250)"`
	Category        int
	AvailableSizes  []string `orm:"size(250)"`
	AvailableColors []string `orm:"size(250)"`
	Quantity        int
	QuantityAlert   int
	ItemPrice       float32
	AltPrice        float32
	Country         string
	Branch          int64
	CreatedBy       int `orm:"omitempty"`
}
