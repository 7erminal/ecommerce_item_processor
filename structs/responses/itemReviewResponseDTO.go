package responses

import (
	"item_processor/models"
)

type ItemReviewsResponseDTO struct {
	StatusCode   int
	ItemsReviews *[]interface{}
	StatusDesc   string
}

type ItemReviewResponseDTO struct {
	StatusCode int
	ItemReview *models.Item_reviews
	StatusDesc string
}
