package requests

type AddItemRequest struct {
	ItemName        string
	Description     string
	Weight          string
	Category        int
	AvailableSizes  []string
	AvailableColors []string
	Quantity        int
	ItemPrice       float32
	CreatedBy       int
}
