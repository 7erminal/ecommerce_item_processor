package requests

type ItemTypeRequest struct {
	Name        string
	Description string
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
