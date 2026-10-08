package requests

type AddReviewRequest struct {
	Review    string
	ReviewBy  string
	ItemId    string
	Reference string
	Rating    float64
	ImagePath string
}
