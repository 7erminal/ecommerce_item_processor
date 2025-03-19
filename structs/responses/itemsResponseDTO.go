package responses

import (
	"item_processor/models"
)

type ItemsBranchesStatsDTO struct {
	BranchStats *[]models.ItemsBranchCountDTO
}

type ItemsCategoryStatsDTO struct {
	CategoryStats *[]models.ItemsCategoryCountDTO
}

type StatsDTO struct {
	BranchStats   *[]models.ItemsCategoryCountDTO
	CategoryStats *[]models.ItemsCategoryCountDTO
}

type ItemsStatsResponseDTO struct {
	StatusCode int
	Stats      *StatsDTO
	StatusDesc string
}

type ItemsResponseDTO struct {
	StatusCode int
	Items      *[]interface{}
	StatusDesc string
}

type ItemResponseDTO struct {
	StatusCode int
	Item       *models.Items
	StatusDesc string
}

type ItemBranchCountResponseDTO struct {
	StatusCode int
	Result     *[]models.ItemBranchCountDTO
	StatusDesc string
}

type ItemBranchCountDTO struct {
	Branch    string
	Category  string
	ItemCount int64
}
