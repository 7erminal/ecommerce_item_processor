package responses

import (
	"time"
)

type Item_reviews struct {
	ItemReviewId string
	Review       string
	Rating       float64
	ReviewBy     string
	Item         *Items
	DateCreated  time.Time
	DateModified time.Time
	Reference    string
	CreatedBy    string
	ModifiedBy   string
	Active       int
}

type ItemReviewsResponseDTO struct {
	StatusCode   int
	ItemsReviews []*Item_reviews
	StatusDesc   string
}

type ItemReviewResponseDTO struct {
	StatusCode int
	ItemReview *Item_reviews
	StatusDesc string
}
