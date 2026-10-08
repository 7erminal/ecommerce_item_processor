package requests

type AddItemRequest struct {
	ItemName        string
	Description     string
	Weight          string
	Category        string
	AvailableSizes  []string
	AvailableColors []string
	Quantity        int
	QuantityAlert   int
	ItemPrice       float32
	AltItemPrice    float32
	ExtraCharges    float32
	Country         string
	Branch          string
	CreatedBy       string
}

type GetItemCount struct {
	Category string
	Branch   string
}
