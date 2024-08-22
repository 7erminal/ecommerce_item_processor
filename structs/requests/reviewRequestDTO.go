package requests

type AddReviewRequest struct {
	Review   string
	ReviewBy int64
	Rating   float64
}
