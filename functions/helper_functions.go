package functions

import (
	"fmt"
	"item_processor/models"
	"item_processor/structs/responses"
	"strconv"
	"time"

	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/core/logs"
)

func ConvertParamsToFeatures(params []orm.Params) []responses.FeaturesItemsDTO {
	var features []responses.FeaturesItemsDTO
	// var items []models.Items

	logs.Info("converting orm params to Features model ")

	q := 0

	for _, param := range params {
		var feature_ responses.FeaturesItemsDTO
		exists := false
		for _, it := range features {
			if it.FeatureId == toInt64(param["feature_id"]) {
				exists = true
				feature_ = it
			}
		}

		if !exists {
			feature := responses.FeaturesItemsDTO{
				FeatureId:    toInt64(param["feature_id"]),
				FeatureName:  toString(param["feature_name"]),
				ImagePath:    toString(param["image_path"]),
				Visible:      toBool(param["visible"]),
				Description:  toString(param["description"]),
				Active:       toInt(param["active"]),
				DateCreated:  toTime(param["date_created"]),
				DateModified: toTime(param["date_modified"]),
				CreatedBy:    toInt(param["created_by"]),
				ModifiedBy:   toInt(param["modified_by"]),
				// Id:            toInt(param["id"]),
				// Name:          toString(param["name"]),
				// ItemPurposeId: toInt(param["item_purpose_id"]),
				// Purpose:       toString(param["purpose"]),
			}

			cat, _ := models.GetCategoriesById(toInt64(param["item_id"]))
			it_price, _ := models.GetItem_pricesById(toInt64(param["item_price_id"]))
			countr_, _ := models.GetCountriesById(toInt64(param["country"]))

			item_ := models.Items{
				ItemId:          toInt64(param["item_id"]),
				ItemName:        toString(param["item_name"]),
				Description:     toString(param["description"]),
				Weight:          toString(param["weight"]),
				Category:        cat,
				ItemPrice:       it_price,
				AvailableSizes:  toString(param["available_sizes"]),
				AvailableColors: toString(param["available_colors"]),
				Material:        toString(param["material"]),
				ImagePath:       toString(param["image_path"]),
				Quantity:        toInt(param["quantity"]),
				Active:          toInt(param["active"]),
				DateCreated:     toTime(param["date_created"]),
				DateModified:    toTime(param["date_modified"]),
				CreatedBy:       toInt(param["created_by"]),
				ModifiedBy:      toInt(param["modified_by"]),
				Country:         countr_,
			}

			feature.Items = append(feature.Items, item_)

			features = append(features, feature)
		} else {
			cat, _ := models.GetCategoriesById(toInt64(param["category_id"]))
			it_price, _ := models.GetItem_pricesById(toInt64(param["item_price_id"]))
			countr_, _ := models.GetCountriesById(toInt64(param["country"]))

			item_ := models.Items{
				ItemId:          toInt64(param["item_id"]),
				ItemName:        toString(param["item_name"]),
				Description:     toString(param["description"]),
				Weight:          toString(param["weight"]),
				Category:        cat,
				ItemPrice:       it_price,
				AvailableSizes:  toString(param["available_sizes"]),
				AvailableColors: toString(param["available_colors"]),
				Material:        toString(param["material"]),
				ImagePath:       toString(param["image_path"]),
				Quantity:        toInt(param["quantity"]),
				Active:          toInt(param["active"]),
				DateCreated:     toTime(param["date_created"]),
				DateModified:    toTime(param["date_modified"]),
				CreatedBy:       toInt(param["created_by"]),
				ModifiedBy:      toInt(param["modified_by"]),
				Country:         countr_,
			}

			feature_.Items = append(feature_.Items, item_)

			logs.Info("On feature ", toString(param["feature_name"]), "Feature Items are ", len(feature_.Items), "on Item ", toString(param["item_name"]))

			// Find the feature and update
			qs := 0
			for _, it := range features {
				if it.FeatureId == toInt64(param["feature_id"]) {
					logs.Info("It matches so replacing and updating")

					features[qs] = feature_
				}
				qs++
			}
		}
		q++
	}

	return features
}

func ConvertParamsToPurposes(params []orm.Params) []responses.PurposesItemsDTO {
	var purposes []responses.PurposesItemsDTO
	// var items []models.Items

	logs.Info("converting orm params to Features model ")

	q := 0

	for _, param := range params {
		var purpose_ responses.PurposesItemsDTO
		exists := false
		for _, it := range purposes {
			if it.PurposeId == toInt64(param["purpose_id"]) {
				exists = true
				purpose_ = it
			}
		}

		if !exists {
			purpose := responses.PurposesItemsDTO{
				PurposeId:    toInt64(param["purpose_id"]),
				Purpose:      toString(param["purpose"]),
				ImagePath:    toString(param["image_path"]),
				Visible:      toBool(param["visible"]),
				Description:  toString(param["description"]),
				Active:       toInt(param["active"]),
				DateCreated:  toTime(param["date_created"]),
				DateModified: toTime(param["date_modified"]),
				CreatedBy:    toInt(param["created_by"]),
				ModifiedBy:   toInt(param["modified_by"]),
				// Id:            toInt(param["id"]),
				// Name:          toString(param["name"]),
				// ItemPurposeId: toInt(param["item_purpose_id"]),
				// Purpose:       toString(param["purpose"]),
			}

			cat, _ := models.GetCategoriesById(toInt64(param["category_id"]))
			it_price, _ := models.GetItem_pricesById(toInt64(param["item_price_id"]))
			countr_, _ := models.GetCountriesById(toInt64(param["country"]))

			item_ := models.Items{
				ItemId:          toInt64(param["item_id"]),
				ItemName:        toString(param["item_name"]),
				Description:     toString(param["description"]),
				Weight:          toString(param["weight"]),
				Category:        cat,
				ItemPrice:       it_price,
				AvailableSizes:  toString(param["available_sizes"]),
				AvailableColors: toString(param["available_colors"]),
				Material:        toString(param["material"]),
				ImagePath:       toString(param["image_path"]),
				Quantity:        toInt(param["quantity"]),
				Active:          toInt(param["active"]),
				DateCreated:     toTime(param["date_created"]),
				DateModified:    toTime(param["date_modified"]),
				CreatedBy:       toInt(param["created_by"]),
				ModifiedBy:      toInt(param["modified_by"]),
				Country:         countr_,
			}

			purpose.Items = append(purpose.Items, item_)

			purposes = append(purposes, purpose)
		} else {
			cat, _ := models.GetCategoriesById(toInt64(param["item_id"]))
			it_price, _ := models.GetItem_pricesById(toInt64(param["item_price_id"]))
			countr_, _ := models.GetCountriesById(toInt64(param["country"]))

			item_ := models.Items{
				ItemId:          toInt64(param["item_id"]),
				ItemName:        toString(param["item_name"]),
				Description:     toString(param["description"]),
				Weight:          toString(param["weight"]),
				Category:        cat,
				ItemPrice:       it_price,
				AvailableSizes:  toString(param["available_sizes"]),
				AvailableColors: toString(param["available_colors"]),
				Material:        toString(param["material"]),
				ImagePath:       toString(param["image_path"]),
				Quantity:        toInt(param["quantity"]),
				Active:          toInt(param["active"]),
				DateCreated:     toTime(param["date_created"]),
				DateModified:    toTime(param["date_modified"]),
				CreatedBy:       toInt(param["created_by"]),
				ModifiedBy:      toInt(param["modified_by"]),
				Country:         countr_,
			}

			purpose_.Items = append(purpose_.Items, item_)

			logs.Info("On feature ", toString(param["feature_name"]), "Feature Items are ", len(purpose_.Items), "on Item ", toString(param["item_name"]))

			// Find the feature and update
			qs := 0
			for _, it := range purposes {
				if it.PurposeId == toInt64(param["feature_id"]) {
					logs.Info("It matches so replacing and updating")

					purposes[qs] = purpose_
				}
				qs++
			}
		}
		q++
	}

	return purposes
}

func toInt(val interface{}) int {
	if val == nil {
		return 0
	}
	switch v := val.(type) {
	case int:
		return v
	case int64:
		return int(v)
	case string:
		intValue, _ := strconv.Atoi(v)
		return intValue
	default:
		return 0
	}
}

func toInt64(val interface{}) int64 {
	if val == nil {
		return 0
	}

	val_ := toString(val)
	id, _ := strconv.ParseInt(val_, 0, 64)

	return id
}

func toString(val interface{}) string {
	if val == nil {
		return ""
	}
	return fmt.Sprintf("%v", val)
}

func toBool(val interface{}) bool {
	if v, ok := val.(bool); ok {
		return v
	}
	return false // or return a default value
}

func toTime(val interface{}) time.Time {
	switch v := val.(type) {
	case time.Time:
		return v
	case string:
		// Parse string to time.Time
		layout := "2006-01-02 15:04:05" // Define your expected layout
		t, err := time.Parse(layout, v)
		if err != nil {
			return time.Time{}
		}
		return t
	case int64:
		// Convert Unix timestamp to time.Time
		return time.Unix(v, 0)
	case int:
		return time.Unix(int64(v), 0)
	case float64:
		// Handle timestamps stored as float64
		return time.Unix(int64(v), 0)
	default:
		return time.Time{}
		// fmt.Errorf("unsupported type: %T", val)
	}
}
