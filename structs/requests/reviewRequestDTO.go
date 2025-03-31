package requests

type AddReviewRequest struct {
	Review    string
	ReviewBy  int64
	ItemId    int64
	Reference int64
	Rating    float64
	ImagePath string
}
