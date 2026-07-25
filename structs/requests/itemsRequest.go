package requests

type ItemTypeRequest struct {
	Name        string
	Description string
}

type ItemQuantityRequest struct {
	ItemId   string
	Quantity int
}
