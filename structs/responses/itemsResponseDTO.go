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

type ItemsResponseDTO struct {
	StatusCode int             `orm:"omitempty"`
	Items      *[]models.Items `orm:"omitempty"`
	StatusDesc string          `orm:"size(255)"`
}

type ItemsResponseDTO2 struct {
	StatusCode int             `orm:"omitempty"`
	Items      *[]models.Items `orm:"omitempty"`
	StatusDesc string          `orm:"size(255)"`
}

type Item_featureResponseDTO struct {
	StatusCode int                   `orm:"omitempty"`
	Result     *models.Item_features `orm:"omitempty"`
	StatusDesc string                `orm:"size(255)"`
}

type Item_featuresResponseDTO struct {
	StatusCode int                     `orm:"omitempty"`
	Result     *[]models.Item_features `orm:"omitempty"`
	StatusDesc string                  `orm:"size(255)"`
}

type Item_purposeResponseDTO struct {
	StatusCode int                   `orm:"omitempty"`
	Result     *models.Item_purposes `orm:"omitempty"`
	StatusDesc string                `orm:"size(255)"`
}

type Item_purposesResponseDTO struct {
	StatusCode int                     `orm:"omitempty"`
	Result     *[]models.Item_purposes `orm:"omitempty"`
	StatusDesc string                  `orm:"size(255)"`
}
