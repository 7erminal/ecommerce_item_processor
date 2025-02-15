package requests

type AddItemRequest struct {
	ItemName        string
	Description     string
	Weight          string
	Category        int
	AvailableSizes  []string
	AvailableColors []string
	Quantity        int
	QuantityAlert   int
	ItemPrice       float32
	AltItemPrice    float32
	ExtraCharges    float32
	Country         string
	Branch          int64
	CreatedBy       int
}

type GetItemCount struct {
	Category string
	Branch   string
}
