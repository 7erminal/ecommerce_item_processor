package requests

type ItemTypeRequest struct {
	Name        string
	Description string
	AddedBy     string
}

type ItemQuantityRequest struct {
	ItemId   string
	Quantity int
}

type ItemPriceRequest struct {
	ItemId       string
	Price        float32
	AltPrice     float32
	ExtraCharges float32
}

type ItemsDTO struct {
	ItemName        string
	Description     string
	Category        string
	AvailableSizes  []string
	AvailableColors []string
	Weight          string
	Quantity        int
	QuantityAlert   int
	ItemPrice       float32
	AltPrice        float32
	ExtraCharges    float32
	Country         string
	Branch          string
	CreatedBy       string
}
