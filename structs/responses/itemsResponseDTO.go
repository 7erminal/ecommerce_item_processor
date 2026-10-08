package responses

import (
	"item_processor/models"
	"time"
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

type Item_prices struct {
	ItemPriceId   string
	ItemPrice     float32
	AltItemPrice  float32
	ShowAltPrice  bool
	Discount      string
	Discount_type string
	ExtraCharges  float32
	Currency      string
	Active        int
	DateCreated   time.Time
	DateModified  time.Time
	CreatedBy     string
	ModifiedBy    string
}

type Categories struct {
	CategoryId   string
	CategoryName string
	ImagePath    string
	Icon         string
	Description  string
	Active       int8
	DateCreated  time.Time
	DateModified time.Time
}

type Status struct {
	StatusId     string
	Status       string
	StatusCode   string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
	Active       int
}

type ItemsDTO struct {
	ItemName        string
	Description     string
	Category        string
	AvailableSizes  []string
	AvailableColors []string
	Weight          string
	Quantity        int
	QuantityAlert   int
	ItemPrice       float32
	AltPrice        float32
	ExtraCharges    float32
	Country         string
	Branch          string
	CreatedBy       string
}

type Items struct {
	ItemId          string
	ItemName        string
	Description     string
	Weight          string
	Category        *Categories
	ItemPrice       *Item_prices
	AvailableSizes  string
	AvailableColors string
	Material        string
	ImagePath       string
	Quantity        int
	Active          int
	DateCreated     time.Time
	DateModified    time.Time
	CreatedBy       string
	ModifiedBy      string
	Country         string
	Branch          string
	Status          *Status
	LastOrderDate   time.Time
	ItemQuantity    *Item_quantity
	ItemFeatures    []*Features
	ItemPurposes    []*Purposes
}

type Item_quantity struct {
	ItemQuantityId string
	Quantity       int
	QuantityAlert  int
	Active         int
	DateCreated    time.Time
	DateModified   time.Time
	CreatedBy      string
	ModifiedBy     string
}

type ItemsStatsResponseDTO struct {
	StatusCode int
	Stats      *StatsDTO
	StatusDesc string
}

type ItemResponseDTO struct {
	StatusCode int
	Item       *Items
	StatusDesc string
}

type ItemBranchCountResponseDTO struct {
	StatusCode int
	Result     *[]ItemBranchCountDTO
	StatusDesc string
}

type ItemBranchCountDTO struct {
	Branch    string
	Category  string
	ItemCount int64
}

type ItemsResponseDTO struct {
	StatusCode int      `orm:"omitempty"`
	Items      *[]Items `orm:"omitempty"`
	StatusDesc string   `orm:"size(255)"`
}

type ItemsResponseDTO2 struct {
	StatusCode int      `orm:"omitempty"`
	Items      *[]Items `orm:"omitempty"`
	StatusDesc string   `orm:"size(255)"`
}

type Item_featureResponseDTO struct {
	StatusCode int            `orm:"omitempty"`
	Result     *Item_features `orm:"omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type Item_featuresResponseDTO struct {
	StatusCode int              `orm:"omitempty"`
	Result     *[]Item_features `orm:"omitempty"`
	StatusDesc string           `orm:"size(255)"`
}

type Item_purposeResponseDTO struct {
	StatusCode int            `orm:"omitempty"`
	Result     *Item_purposes `orm:"omitempty"`
	StatusDesc string         `orm:"size(255)"`
}

type Item_purposesResponseDTO struct {
	StatusCode int              `orm:"omitempty"`
	Result     []*Item_purposes `orm:"omitempty"`
	StatusDesc string           `orm:"size(255)"`
}

type Item_features struct {
	ItemFeatureId string
	Feature       *Features
}

type Item_purposes struct {
	ItemPurposeId string
	Purpose       *Purposes
}

type Item_types struct {
	ItemTypeId   string
	Name         string
	Description  string
	DateCreated  time.Time
	DateModified time.Time
	CreatedBy    string
	ModifiedBy   string
	Active       int
}

type CategoryResponseDTO struct {
	StatusCode int
	Category   *Categories
	StatusDesc string
}

type CategoriesResponseDTO struct {
	StatusCode int
	Categories []*Categories
	StatusDesc string
}
