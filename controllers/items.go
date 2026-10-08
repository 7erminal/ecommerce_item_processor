package controllers

import (
	"encoding/json"
	"errors"
	"item_processor/controllers/functions"
	"item_processor/models"
	"item_processor/structs/requests"
	"item_processor/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// ItemsController operations for Items
type ItemsController struct {
	beego.Controller
}

// URLMapping ...
func (c *ItemsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetItemQuantity", c.GetItemQuantity)
	// c.Mapping("GetItemFeaturesByItem", c.GetItemFeaturesByItem)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UpdateItemImage", c.UpdateItemImage)
	c.Mapping("GetItemStats", c.GetItemStats)
	c.Mapping("GetItemCountByType", c.GetItemCountByType)
	c.Mapping("GetItemCount", c.GetItemCount)
	c.Mapping("CheckItemQuantity", c.CheckItemQuantity)
	c.Mapping("GetItemCountWithTypeAndBranch", c.GetItemCountWithTypeAndBranch)
	c.Mapping("UpdateItemQuantity", c.UpdateItemQuantity)
}

// Post ...
// @Title Post
// @Description create Items
// @Param	body		body 	requests.AddItemRequest	true		"body for Items content"
// @Success 201 {int} responses.ItemsResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *ItemsController) Post() {
	var t requests.AddItemRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)
	logs.Info("Received ", t)
	logs.Info("Item country received is ", t.Country)

	errorCode := 302
	message := "Error adding item"

	creator := t.CreatedBy

	// Structure Available Sizes
	aSizes := strings.Join(t.AvailableSizes, ",")

	// Structure Available Colors
	aColors := strings.Join(t.AvailableColors, ",")

	categoryId, err := strconv.ParseInt(t.Category, 10, 64)
	p, err := models.GetCategoriesById(categoryId)

	if err == nil {
		// Get Currency
		cr, cerr := functions.GetCurrencyWithName(&c.Controller, "GH₵")

		if cr.StatusCode == 200 {
			// Add price for item
			logs.Info("Adding price for item with price ", t.ItemPrice, " and alt price ", t.AltItemPrice, " and extra charges ", t.ExtraCharges)
			it := models.Item_prices{
				ItemPrice:    t.ItemPrice,
				AltItemPrice: t.AltItemPrice,
				ShowAltPrice: false,
				ExtraCharges: t.ExtraCharges,
				Currency:     strconv.FormatInt(cr.Result.CurrencyId, 10),
				Active:       1,
				CreatedBy:    creator,
				DateCreated:  time.Now(),
				ModifiedBy:   creator,
				DateModified: time.Now()}

			logs.Info("Adding price to item to create at a go")
			if _, err := models.AddItem_prices(&it); err == nil {
				country := functions.GetCountryWithCode(&c.Controller, t.Country)
				branchId, err := strconv.ParseInt(t.Branch, 10, 64)
				if err != nil {
					logs.Error(err.Error())
					message = "Error parsing branch ID: " + err.Error()
					resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
					c.Data["json"] = resp
					c.ServeJSON()
					return
				}
				branch := functions.GetBranch(&c.Controller, branchId)

				status := models.Status{}

				if status_, err := models.GetStatusByCode("GOOD"); err == nil {
					status = *status_
				}

				// Add item if getting category and price addition does not result in an error
				v := models.Items{
					ItemName:        t.ItemName,
					Description:     t.Description,
					Weight:          t.Weight,
					Category:        p,
					ItemPrice:       &it,
					AvailableSizes:  aSizes,
					AvailableColors: aColors,
					Quantity:        t.Quantity,
					Country:         country.Result.CountryId,
					Branch:          branch.Result.BranchId,
					Active:          1,
					DateCreated:     time.Now(),
					DateModified:    time.Now(),
					CreatedBy:       creator,
					ModifiedBy:      creator,
					Status:          &status,
				}

				if _, err := models.AddItems(&v); err == nil {
					// Add quantity for item
					qu := models.Item_quantity{Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}

					if _, err := models.AddItem_quantity(&qu); err == nil {
						c.Ctx.Output.SetStatus(200)
						logs.Info("Item added successfully with ID ", v.ItemId)

						errorCode = 200
						message = "Item added successfully with ID " + strconv.FormatInt(v.ItemId, 10)

						categoryData := responses.Categories{
							CategoryId:   strconv.FormatInt(v.Category.CategoryId, 10),
							CategoryName: v.Category.CategoryName,
							Description:  v.Category.Description,
							ImagePath:    v.Category.ImagePath,
							Active:       v.Category.Active,
							Icon:         v.Category.Icon,
							DateCreated:  v.Category.DateCreated,
							DateModified: v.Category.DateModified,
						}
						price := responses.Item_prices{
							ItemPriceId:   strconv.FormatInt(v.ItemPrice.ItemPriceId, 10),
							ItemPrice:     v.ItemPrice.ItemPrice,
							AltItemPrice:  v.ItemPrice.AltItemPrice,
							ShowAltPrice:  v.ItemPrice.ShowAltPrice,
							Discount:      v.ItemPrice.Discount,
							Discount_type: v.ItemPrice.Discount_type,
							ExtraCharges:  v.ItemPrice.ExtraCharges,
							Currency:      v.ItemPrice.Currency,
							Active:        v.ItemPrice.Active,
							DateCreated:   v.ItemPrice.DateCreated,
							DateModified:  v.ItemPrice.DateModified,
							CreatedBy:     v.ItemPrice.CreatedBy,
							ModifiedBy:    v.ItemPrice.ModifiedBy,
						}
						status := responses.Status{
							StatusId:     strconv.FormatInt(v.Status.StatusId, 10),
							Status:       v.Status.Status,
							StatusCode:   v.Status.StatusCode,
							DateCreated:  v.Status.DateCreated,
							DateModified: v.Status.DateModified,
							Active:       v.Status.Active,
						}
						quantity := responses.Item_quantity{
							ItemQuantityId: strconv.FormatInt(v.ItemQuantity.ItemQuantityId, 10),
							Quantity:       v.ItemQuantity.Quantity,
							QuantityAlert:  v.ItemQuantity.QuantityAlert,
							Active:         v.ItemQuantity.Active,
							DateCreated:    v.ItemQuantity.DateCreated,
							DateModified:   v.ItemQuantity.DateModified,
							CreatedBy:      v.ItemQuantity.CreatedBy,
							ModifiedBy:     v.ItemQuantity.ModifiedBy,
						}
						features := []*responses.Features{}
						for _, f := range v.ItemFeatures {
							features = append(features, &responses.Features{
								FeatureId:    strconv.FormatInt(f.FeatureId, 10),
								Feature:      f.FeatureName,
								Description:  f.Description,
								DateCreated:  f.DateCreated,
								DateModified: f.DateModified,
								ImagePath:    f.ImagePath,
								Active:       f.Active,
							})
						}

						purposes := []*responses.Purposes{}
						for _, p := range v.ItemPurposes {
							purposes = append(purposes, &responses.Purposes{
								PurposeId:    strconv.FormatInt(p.PurposeId, 10),
								Purpose:      p.Purpose,
								Description:  p.Description,
								DateCreated:  p.DateCreated,
								DateModified: p.DateModified,
								ImagePath:    p.ImagePath,
								Active:       p.Active,
							})
						}
						respData := responses.Items{
							ItemId:          strconv.FormatInt(v.ItemId, 10),
							ItemName:        v.ItemName,
							Description:     v.Description,
							Weight:          v.Weight,
							Category:        &categoryData,
							ItemPrice:       &price,
							AvailableSizes:  v.AvailableSizes,
							AvailableColors: v.AvailableColors,
							Material:        v.Material,
							ImagePath:       v.ImagePath,
							Quantity:        v.Quantity,
							Active:          v.Active,
							DateCreated:     v.DateCreated,
							DateModified:    v.DateModified,
							CreatedBy:       v.CreatedBy,
							ModifiedBy:      v.ModifiedBy,
							Country:         strconv.FormatInt(v.Country, 10),
							Branch:          strconv.FormatInt(v.Branch, 10),
							Status:          &status,
							LastOrderDate:   v.LastOrderDate,
							ItemQuantity:    &quantity,
							ItemFeatures:    features,
							ItemPurposes:    purposes,
						}
						resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: &respData, StatusDesc: message}
						c.Data["json"] = resp
					} else {
						logs.Error(err.Error())
						message = "Error adding item quantity: " + err.Error()
						resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
						c.Data["json"] = resp
					}
				} else {
					logs.Error(err.Error())
					errorCode = 301
					message = "Error adding item: " + err.Error()
					resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
					c.Data["json"] = resp
				}
			} else {
				logs.Error(err.Error())
				errorCode = 301
				message = "Error adding item price: " + err.Error()
				resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
				c.Data["json"] = resp
			}
		} else {
			// resp := models.ItemsResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			// c.Data["json"] = resp
			errorCode = 301
			message = "Error fetching currency: " + cerr.Error()
			resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
			c.Data["json"] = resp
			logs.Error(cerr.Error())
		}
	} else {
		logs.Error(err.Error())
		errorCode = 301
		message = "Error fetching category: " + err.Error()
		resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: nil, StatusDesc: message}
		c.Data["json"] = resp
		// c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Items by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Items
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ItemsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItemsById(id)
	if err != nil {
		logs.Error("Error fetching item ", err)
		resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		categoryData := responses.Categories{
			CategoryId:   strconv.FormatInt(v.Category.CategoryId, 10),
			CategoryName: v.Category.CategoryName,
			Description:  v.Category.Description,
			ImagePath:    v.Category.ImagePath,
			Active:       v.Category.Active,
			Icon:         v.Category.Icon,
			DateCreated:  v.Category.DateCreated,
			DateModified: v.Category.DateModified,
		}
		price := responses.Item_prices{
			ItemPriceId:   strconv.FormatInt(v.ItemPrice.ItemPriceId, 10),
			ItemPrice:     v.ItemPrice.ItemPrice,
			AltItemPrice:  v.ItemPrice.AltItemPrice,
			ShowAltPrice:  v.ItemPrice.ShowAltPrice,
			Discount:      v.ItemPrice.Discount,
			Discount_type: v.ItemPrice.Discount_type,
			ExtraCharges:  v.ItemPrice.ExtraCharges,
			Currency:      v.ItemPrice.Currency,
			Active:        v.ItemPrice.Active,
			DateCreated:   v.ItemPrice.DateCreated,
			DateModified:  v.ItemPrice.DateModified,
			CreatedBy:     v.ItemPrice.CreatedBy,
			ModifiedBy:    v.ItemPrice.ModifiedBy,
		}
		status := responses.Status{
			StatusId:     strconv.FormatInt(v.Status.StatusId, 10),
			Status:       v.Status.Status,
			StatusCode:   v.Status.StatusCode,
			DateCreated:  v.Status.DateCreated,
			DateModified: v.Status.DateModified,
			Active:       v.Status.Active,
		}
		quantity := responses.Item_quantity{
			ItemQuantityId: strconv.FormatInt(v.ItemQuantity.ItemQuantityId, 10),
			Quantity:       v.ItemQuantity.Quantity,
			QuantityAlert:  v.ItemQuantity.QuantityAlert,
			Active:         v.ItemQuantity.Active,
			DateCreated:    v.ItemQuantity.DateCreated,
			DateModified:   v.ItemQuantity.DateModified,
			CreatedBy:      v.ItemQuantity.CreatedBy,
			ModifiedBy:     v.ItemQuantity.ModifiedBy,
		}
		features := []*responses.Features{}
		for _, f := range v.ItemFeatures {
			features = append(features, &responses.Features{
				FeatureId:    strconv.FormatInt(f.FeatureId, 10),
				Feature:      f.FeatureName,
				Description:  f.Description,
				DateCreated:  f.DateCreated,
				DateModified: f.DateModified,
				ImagePath:    f.ImagePath,
				Active:       f.Active,
			})
		}

		purposes := []*responses.Purposes{}
		for _, p := range v.ItemPurposes {
			purposes = append(purposes, &responses.Purposes{
				PurposeId:    strconv.FormatInt(p.PurposeId, 10),
				Purpose:      p.Purpose,
				Description:  p.Description,
				DateCreated:  p.DateCreated,
				DateModified: p.DateModified,
				ImagePath:    p.ImagePath,
				Active:       p.Active,
			})
		}
		respData := responses.Items{
			ItemId:          strconv.FormatInt(v.ItemId, 10),
			ItemName:        v.ItemName,
			Description:     v.Description,
			Weight:          v.Weight,
			Category:        &categoryData,
			ItemPrice:       &price,
			AvailableSizes:  v.AvailableSizes,
			AvailableColors: v.AvailableColors,
			Material:        v.Material,
			ImagePath:       v.ImagePath,
			Quantity:        v.Quantity,
			Active:          v.Active,
			DateCreated:     v.DateCreated,
			DateModified:    v.DateModified,
			CreatedBy:       v.CreatedBy,
			ModifiedBy:      v.ModifiedBy,
			Country:         strconv.FormatInt(v.Country, 10),
			Branch:          strconv.FormatInt(v.Branch, 10),
			Status:          &status,
			LastOrderDate:   v.LastOrderDate,
			ItemQuantity:    &quantity,
			ItemFeatures:    features,
			ItemPurposes:    purposes,
		}
		resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemQuantity ...
// @Title Get Item Quantity
// @Description get Item_quantity by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemQuantityResponseDTO
// @Failure 403 :id is empty
// @router /quantity/:id [get]
func (c *ItemsController) GetItemQuantity() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	logs.Debug("Item ID to get quantity is ", id)
	// q, err := models.GetItemsById(id)
	v, err := models.GetItem_quantityByItemId(id)
	if err != nil {
		logs.Error("Error fetching quantity of item ... ", err.Error())
		resp := responses.ItemQuantityResponseDTO{StatusCode: 301, Quantity: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		respData := responses.Item_quantity{
			ItemQuantityId: strconv.FormatInt(v.ItemQuantityId, 10),
			Quantity:       v.Quantity,
			QuantityAlert:  v.QuantityAlert,
			Active:         v.Active,
			DateCreated:    v.DateCreated,
			DateModified:   v.DateModified,
			CreatedBy:      v.CreatedBy,
			ModifiedBy:     v.ModifiedBy,
		}
		resp := responses.ItemQuantityResponseDTO{StatusCode: 200, Quantity: &respData, StatusDesc: "Quantity fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemCount ...
// @Title Get Item Quantity
// @Description get Item_quantity by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/ [get]
func (c *ItemsController) GetItemCount() {
	// q, err := models.GetItemsById(id)
	var query = make(map[string]string)
	var search = make(map[string]string)
	logs.Info("Getting count")

	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	v, err := models.GetItemCount(query, search)
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of items ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemCountByType ...
// @Title Get Item Quantity
// @Description get Item_quantity by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/type/:name [get]
func (c *ItemsController) GetItemCountByType() {
	// q, err := models.GetItemsById(id)
	name := c.Ctx.Input.Param(":name")

	v, err := models.GetItemCountByType(name)
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of items ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemCountWithTypeAndBranch ...
// @Title Get Item Count with type and branch
// @Description get item count with type and branch
// @Param	body		body 	requests.GetItemCount	true		"body for Items content"
// @Success 200 {object} responses.StringResponseDTO
// @Failure 403 :id is empty
// @router /count/type [post]
func (c *ItemsController) GetItemCountWithTypeAndBranch() {
	// q, err := models.GetItemsById(id)
	var t requests.GetItemCount
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)

	logs.Info("Category is ", t.Category, " and branch is ", t.Branch)

	v, err := models.GetItemCountWithTypeAndBranch(t.Category, t.Branch)
	count := strconv.FormatInt(v, 10)
	if err != nil {
		logs.Error("Error fetching count of items ... ", err.Error())
		resp := responses.StringResponseDTO{StatusCode: 301, Value: "", StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := responses.StringResponseDTO{StatusCode: 200, Value: count, StatusDesc: "Count fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Items
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	search	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Items
// @Failure 403
// @router / [get]
func (c *ItemsController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var search = make(map[string]string)
	var limit int64 = 100
	var offset int64

	logs.Info("Getting all items")

	// fields: col1,col2,entity.col3
	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	// limit: 10 (default is 10)
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	// offset: 0 (default is 0)
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	// sortby: col1,col2
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	// order: desc,asc
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	// query: k:v,k:v
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// set active
	activeQuery := "Active:1"
	if v := activeQuery; v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// Get only items with active categories
	activeCategoryQuery := "Category__active:1"
	if v := activeCategoryQuery; v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	// search: k:v,k:v
	if v := c.GetString("search"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = errors.New("Error: invalid search key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			search[k] = v
		}
	}

	logs.Info("Limit being sent is ", limit, " query is ", query, " and search is ", search)

	l, err := models.GetAllItems(query, fields, sortby, order, offset, limit, search)
	if err != nil {
		resp := responses.ItemsResponseDTO{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		itemsResp := []responses.Items{}
		for _, urs := range l {
			m := urs.(models.Items)

			categoryData := responses.Categories{
				CategoryId:   strconv.FormatInt(m.Category.CategoryId, 10),
				CategoryName: m.Category.CategoryName,
				Description:  m.Category.Description,
				ImagePath:    m.Category.ImagePath,
				Active:       m.Category.Active,
				Icon:         m.Category.Icon,
				DateCreated:  m.Category.DateCreated,
				DateModified: m.Category.DateModified,
			}
			price := responses.Item_prices{
				ItemPriceId:   strconv.FormatInt(m.ItemPrice.ItemPriceId, 10),
				ItemPrice:     m.ItemPrice.ItemPrice,
				AltItemPrice:  m.ItemPrice.AltItemPrice,
				ShowAltPrice:  m.ItemPrice.ShowAltPrice,
				Discount:      m.ItemPrice.Discount,
				Discount_type: m.ItemPrice.Discount_type,
				ExtraCharges:  m.ItemPrice.ExtraCharges,
				Currency:      m.ItemPrice.Currency,
				Active:        m.ItemPrice.Active,
				DateCreated:   m.ItemPrice.DateCreated,
				DateModified:  m.ItemPrice.DateModified,
				CreatedBy:     m.ItemPrice.CreatedBy,
				ModifiedBy:    m.ItemPrice.ModifiedBy,
			}
			status := responses.Status{
				StatusId:     strconv.FormatInt(m.Status.StatusId, 10),
				Status:       m.Status.Status,
				StatusCode:   m.Status.StatusCode,
				DateCreated:  m.Status.DateCreated,
				DateModified: m.Status.DateModified,
				Active:       m.Status.Active,
			}
			quantity := responses.Item_quantity{
				ItemQuantityId: strconv.FormatInt(m.ItemQuantity.ItemQuantityId, 10),
				Quantity:       m.ItemQuantity.Quantity,
				QuantityAlert:  m.ItemQuantity.QuantityAlert,
				Active:         m.ItemQuantity.Active,
				DateCreated:    m.ItemQuantity.DateCreated,
				DateModified:   m.ItemQuantity.DateModified,
				CreatedBy:      m.ItemQuantity.CreatedBy,
				ModifiedBy:     m.ItemQuantity.ModifiedBy,
			}
			features := []*responses.Features{}
			for _, f := range m.ItemFeatures {
				features = append(features, &responses.Features{
					FeatureId:    strconv.FormatInt(f.FeatureId, 10),
					Feature:      f.FeatureName,
					Description:  f.Description,
					DateCreated:  f.DateCreated,
					DateModified: f.DateModified,
					ImagePath:    f.ImagePath,
					Active:       f.Active,
				})
			}

			purposes := []*responses.Purposes{}
			for _, p := range m.ItemPurposes {
				purposes = append(purposes, &responses.Purposes{
					PurposeId:    strconv.FormatInt(p.PurposeId, 10),
					Purpose:      p.Purpose,
					Description:  p.Description,
					DateCreated:  p.DateCreated,
					DateModified: p.DateModified,
					ImagePath:    p.ImagePath,
					Active:       p.Active,
				})
			}
			respData := responses.Items{
				ItemId:          strconv.FormatInt(m.ItemId, 10),
				ItemName:        m.ItemName,
				Description:     m.Description,
				Weight:          m.Weight,
				Category:        &categoryData,
				ItemPrice:       &price,
				AvailableSizes:  m.AvailableSizes,
				AvailableColors: m.AvailableColors,
				Material:        m.Material,
				ImagePath:       m.ImagePath,
				Quantity:        m.Quantity,
				Active:          m.Active,
				DateCreated:     m.DateCreated,
				DateModified:    m.DateModified,
				CreatedBy:       m.CreatedBy,
				ModifiedBy:      m.ModifiedBy,
				Country:         strconv.FormatInt(m.Country, 10),
				Branch:          strconv.FormatInt(m.Branch, 10),
				Status:          &status,
				LastOrderDate:   m.LastOrderDate,
				ItemQuantity:    &quantity,
				ItemFeatures:    features,
				ItemPurposes:    purposes,
			}
			itemsResp = append(itemsResp, respData)
		}
		logs.Info("Items returned are ", l)
		resp := responses.ItemsResponseDTO{StatusCode: 200, Items: &itemsResp, StatusDesc: "Items fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Items
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ItemsDTO	true		"body for Items content"
// @Success 200 {object} models.Items
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ItemsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var t requests.ItemsDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)

	creator := t.CreatedBy

	// Structure Available Sizes
	aSizes := strings.Join(t.AvailableSizes, ",")

	// Structure Available Colors
	aColors := strings.Join(t.AvailableColors, ",")

	categoryInt, err := strconv.ParseInt(t.Category, 10, 64)
	p, err := models.GetCategoriesById(categoryInt)

	if err == nil {
		iv, err := models.GetItemsById(id)
		if err != nil {
			resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		} else {
			// Get Currency
			cr, cerr := functions.GetCurrencyWithName(&c.Controller, "GHC")

			if cerr == nil {
				if ip, err := models.GetItem_pricesById(iv.ItemPrice.ItemPriceId); err == nil {
					// Add price for item
					it := models.Item_prices{ItemPriceId: iv.ItemPrice.ItemPriceId, ItemPrice: t.ItemPrice, AltItemPrice: t.AltPrice, ShowAltPrice: false, ExtraCharges: t.ExtraCharges, Currency: strconv.FormatInt(cr.Result.CurrencyId, 10), Active: 1, ModifiedBy: creator, DateCreated: ip.DateCreated, CreatedBy: ip.CreatedBy, DateModified: time.Now()}

					logs.Info("Modifying price for item")

					if err := models.UpdateItem_pricesById(&it); err == nil {
						// Add item if getting category and price addition does not result in an error
						country := functions.GetCountryWithCode(&c.Controller, t.Country)
						branchid, err := strconv.ParseInt(t.Branch, 10, 64)
						branch := functions.GetBranch(&c.Controller, branchid)
						if err != nil {
							logs.Error("Failed to parse branch ID")
							resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: "Invalid branch ID"}
							c.Data["json"] = resp
						} else if country.StatusCode != 200 {
							logs.Error("Country does not exist")
							resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: "Country does not exist"}
							c.Data["json"] = resp
						} else if branch.StatusCode != 200 {
							logs.Error("Branch does not exist")
							resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: "Branch does not exist"}
							c.Data["json"] = resp
						} else {
							logs.Info("Modifying item with ID ", id)
							v := models.Items{
								ItemId:          id,
								ItemName:        t.ItemName,
								Country:         country.Result.CountryId,
								Branch:          branch.Result.BranchId,
								Description:     t.Description,
								Category:        p,
								ImagePath:       iv.ImagePath,
								ItemPrice:       &it,
								AvailableSizes:  aSizes,
								AvailableColors: aColors,
								Quantity:        t.Quantity,
								Active:          1,
								DateModified:    time.Now(),
								ModifiedBy:      creator,
								CreatedBy:       iv.CreatedBy,
								DateCreated:     iv.DateCreated,
								Weight:          t.Weight}

							if err := models.UpdateItemsById(&v); err == nil {
								// Add quantity for item

								iq, err := models.GetItem_quantityByItemId(id)
								if err != nil {
									resp := responses.ItemResponseDTO{StatusCode: 304, Item: nil, StatusDesc: "Item quantity not set"}
									c.Data["json"] = resp
									qu := models.Item_quantity{Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}
									if _, err := models.AddItem_quantity(&qu); err == nil {
										item, err := models.GetItemsById(v.ItemId)
										if err != nil {

											logs.Error(err.Error())
											resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
											c.Data["json"] = resp
										} else {
											logs.Info("Item fetched successfully")
											logs.Info("Item quantity is ", item.ItemQuantity)
											logs.Info("Item quantity added successfully")
										}
										c.Ctx.Output.SetStatus(200)

										categoryData := responses.Categories{
											CategoryId:   strconv.FormatInt(v.Category.CategoryId, 10),
											CategoryName: v.Category.CategoryName,
											Description:  v.Category.Description,
											ImagePath:    v.Category.ImagePath,
											Active:       v.Category.Active,
											Icon:         v.Category.Icon,
											DateCreated:  v.Category.DateCreated,
											DateModified: v.Category.DateModified,
										}
										price := responses.Item_prices{
											ItemPriceId:   strconv.FormatInt(v.ItemPrice.ItemPriceId, 10),
											ItemPrice:     v.ItemPrice.ItemPrice,
											AltItemPrice:  v.ItemPrice.AltItemPrice,
											ShowAltPrice:  v.ItemPrice.ShowAltPrice,
											Discount:      v.ItemPrice.Discount,
											Discount_type: v.ItemPrice.Discount_type,
											ExtraCharges:  v.ItemPrice.ExtraCharges,
											Currency:      v.ItemPrice.Currency,
											Active:        v.ItemPrice.Active,
											DateCreated:   v.ItemPrice.DateCreated,
											DateModified:  v.ItemPrice.DateModified,
											CreatedBy:     v.ItemPrice.CreatedBy,
											ModifiedBy:    v.ItemPrice.ModifiedBy,
										}
										status := responses.Status{
											StatusId:     strconv.FormatInt(v.Status.StatusId, 10),
											Status:       v.Status.Status,
											StatusCode:   v.Status.StatusCode,
											DateCreated:  v.Status.DateCreated,
											DateModified: v.Status.DateModified,
											Active:       v.Status.Active,
										}
										quantity := responses.Item_quantity{
											ItemQuantityId: strconv.FormatInt(v.ItemQuantity.ItemQuantityId, 10),
											Quantity:       v.ItemQuantity.Quantity,
											QuantityAlert:  v.ItemQuantity.QuantityAlert,
											Active:         v.ItemQuantity.Active,
											DateCreated:    v.ItemQuantity.DateCreated,
											DateModified:   v.ItemQuantity.DateModified,
											CreatedBy:      v.ItemQuantity.CreatedBy,
											ModifiedBy:     v.ItemQuantity.ModifiedBy,
										}
										features := []*responses.Features{}
										for _, f := range v.ItemFeatures {
											features = append(features, &responses.Features{
												FeatureId:    strconv.FormatInt(f.FeatureId, 10),
												Feature:      f.FeatureName,
												Description:  f.Description,
												DateCreated:  f.DateCreated,
												DateModified: f.DateModified,
												ImagePath:    f.ImagePath,
												Active:       f.Active,
											})
										}

										purposes := []*responses.Purposes{}
										for _, p := range v.ItemPurposes {
											purposes = append(purposes, &responses.Purposes{
												PurposeId:    strconv.FormatInt(p.PurposeId, 10),
												Purpose:      p.Purpose,
												Description:  p.Description,
												DateCreated:  p.DateCreated,
												DateModified: p.DateModified,
												ImagePath:    p.ImagePath,
												Active:       p.Active,
											})
										}
										respData := responses.Items{
											ItemId:          strconv.FormatInt(v.ItemId, 10),
											ItemName:        v.ItemName,
											Description:     v.Description,
											Weight:          v.Weight,
											Category:        &categoryData,
											ItemPrice:       &price,
											AvailableSizes:  v.AvailableSizes,
											AvailableColors: v.AvailableColors,
											Material:        v.Material,
											ImagePath:       v.ImagePath,
											Quantity:        v.Quantity,
											Active:          v.Active,
											DateCreated:     v.DateCreated,
											DateModified:    v.DateModified,
											CreatedBy:       v.CreatedBy,
											ModifiedBy:      v.ModifiedBy,
											Country:         strconv.FormatInt(v.Country, 10),
											Branch:          strconv.FormatInt(v.Branch, 10),
											Status:          &status,
											LastOrderDate:   v.LastOrderDate,
											ItemQuantity:    &quantity,
											ItemFeatures:    features,
											ItemPurposes:    purposes,
										}

										resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item successfully added"}
										c.Data["json"] = resp
									} else {
										logs.Error(err.Error())
										resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
										c.Data["json"] = resp
									}
								} else {
									qu := models.Item_quantity{ItemQuantityId: iq.ItemQuantityId, Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: iq.CreatedBy, DateCreated: iq.DateCreated, ModifiedBy: creator, DateModified: time.Now()}

									if err := models.UpdateItem_quantityById(&qu); err == nil {
										item, err := models.GetItemsById(v.ItemId)
										if err != nil {

											logs.Error(err.Error())
											resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
											c.Data["json"] = resp
										} else {
											logs.Info("Item fetched successfully")
											logs.Info("Item quantity is ", item.ItemQuantity)
											logs.Info("Item quantity added successfully")
										}

										c.Ctx.Output.SetStatus(200)

										categoryData := responses.Categories{
											CategoryId:   strconv.FormatInt(v.Category.CategoryId, 10),
											CategoryName: v.Category.CategoryName,
											Description:  v.Category.Description,
											ImagePath:    v.Category.ImagePath,
											Active:       v.Category.Active,
											Icon:         v.Category.Icon,
											DateCreated:  v.Category.DateCreated,
											DateModified: v.Category.DateModified,
										}
										price := responses.Item_prices{
											ItemPriceId:   strconv.FormatInt(v.ItemPrice.ItemPriceId, 10),
											ItemPrice:     v.ItemPrice.ItemPrice,
											AltItemPrice:  v.ItemPrice.AltItemPrice,
											ShowAltPrice:  v.ItemPrice.ShowAltPrice,
											Discount:      v.ItemPrice.Discount,
											Discount_type: v.ItemPrice.Discount_type,
											ExtraCharges:  v.ItemPrice.ExtraCharges,
											Currency:      v.ItemPrice.Currency,
											Active:        v.ItemPrice.Active,
											DateCreated:   v.ItemPrice.DateCreated,
											DateModified:  v.ItemPrice.DateModified,
											CreatedBy:     v.ItemPrice.CreatedBy,
											ModifiedBy:    v.ItemPrice.ModifiedBy,
										}
										status := responses.Status{
											StatusId:     strconv.FormatInt(v.Status.StatusId, 10),
											Status:       v.Status.Status,
											StatusCode:   v.Status.StatusCode,
											DateCreated:  v.Status.DateCreated,
											DateModified: v.Status.DateModified,
											Active:       v.Status.Active,
										}
										quantity := responses.Item_quantity{
											ItemQuantityId: strconv.FormatInt(v.ItemQuantity.ItemQuantityId, 10),
											Quantity:       v.ItemQuantity.Quantity,
											QuantityAlert:  v.ItemQuantity.QuantityAlert,
											Active:         v.ItemQuantity.Active,
											DateCreated:    v.ItemQuantity.DateCreated,
											DateModified:   v.ItemQuantity.DateModified,
											CreatedBy:      v.ItemQuantity.CreatedBy,
											ModifiedBy:     v.ItemQuantity.ModifiedBy,
										}
										features := []*responses.Features{}
										for _, f := range v.ItemFeatures {
											features = append(features, &responses.Features{
												FeatureId:    strconv.FormatInt(f.FeatureId, 10),
												Feature:      f.FeatureName,
												Description:  f.Description,
												DateCreated:  f.DateCreated,
												DateModified: f.DateModified,
												ImagePath:    f.ImagePath,
												Active:       f.Active,
											})
										}

										purposes := []*responses.Purposes{}
										for _, p := range v.ItemPurposes {
											purposes = append(purposes, &responses.Purposes{
												PurposeId:    strconv.FormatInt(p.PurposeId, 10),
												Purpose:      p.Purpose,
												Description:  p.Description,
												DateCreated:  p.DateCreated,
												DateModified: p.DateModified,
												ImagePath:    p.ImagePath,
												Active:       p.Active,
											})
										}
										respData := responses.Items{
											ItemId:          strconv.FormatInt(v.ItemId, 10),
											ItemName:        v.ItemName,
											Description:     v.Description,
											Weight:          v.Weight,
											Category:        &categoryData,
											ItemPrice:       &price,
											AvailableSizes:  v.AvailableSizes,
											AvailableColors: v.AvailableColors,
											Material:        v.Material,
											ImagePath:       v.ImagePath,
											Quantity:        v.Quantity,
											Active:          v.Active,
											DateCreated:     v.DateCreated,
											DateModified:    v.DateModified,
											CreatedBy:       v.CreatedBy,
											ModifiedBy:      v.ModifiedBy,
											Country:         strconv.FormatInt(v.Country, 10),
											Branch:          strconv.FormatInt(v.Branch, 10),
											Status:          &status,
											LastOrderDate:   v.LastOrderDate,
											ItemQuantity:    &quantity,
											ItemFeatures:    features,
											ItemPurposes:    purposes,
										}

										resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item successfully updated"}
										c.Data["json"] = resp
									} else {
										logs.Error(err.Error())
										resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
										c.Data["json"] = resp
									}
								}

							} else {
								logs.Error(err.Error())
								resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
								c.Data["json"] = resp
							}
						}
					} else {
						logs.Error(err.Error())
						resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
						c.Data["json"] = resp
					}
				} else {
					logs.Error(err.Error())
					resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
					c.Data["json"] = resp
				}
			} else {
				resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
				c.Data["json"] = resp
			}
		}

	} else {
		logs.Error(err.Error())
		resp := responses.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
		// c.Data["json"] = err.Error()
	}

	c.ServeJSON()
}

// UpdateItemImage ...
// @Title Update Item Image
// @Description update the Item's image
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ImageUpdateRequest	true		"body for Items content"
// @Success 200 {object} models.Items
// @Failure 403 :id is not int
// @router /update-item-image/:id [put]
func (c *ItemsController) UpdateItemImage() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var t requests.ImageUpdateRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)
	logs.Info("Request received. Image path is ", t.ImagePath, " Item id is ", idStr)

	iv, err := models.GetItemsById(id)
	if err != nil {
		resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		iv.ImagePath = t.ImagePath

		if err := models.UpdateItemsById(iv); err == nil {
			// Add quantity for item

			c.Ctx.Output.SetStatus(200)

			categoryData := responses.Categories{
				CategoryId:   strconv.FormatInt(iv.Category.CategoryId, 10),
				CategoryName: iv.Category.CategoryName,
				Description:  iv.Category.Description,
				ImagePath:    iv.Category.ImagePath,
				Active:       iv.Category.Active,
				Icon:         iv.Category.Icon,
				DateCreated:  iv.Category.DateCreated,
				DateModified: iv.Category.DateModified,
			}
			price := responses.Item_prices{
				ItemPriceId:   strconv.FormatInt(iv.ItemPrice.ItemPriceId, 10),
				ItemPrice:     iv.ItemPrice.ItemPrice,
				AltItemPrice:  iv.ItemPrice.AltItemPrice,
				ShowAltPrice:  iv.ItemPrice.ShowAltPrice,
				Discount:      iv.ItemPrice.Discount,
				Discount_type: iv.ItemPrice.Discount_type,
				ExtraCharges:  iv.ItemPrice.ExtraCharges,
				Currency:      iv.ItemPrice.Currency,
				Active:        iv.ItemPrice.Active,
				DateCreated:   iv.ItemPrice.DateCreated,
				DateModified:  iv.ItemPrice.DateModified,
				CreatedBy:     iv.ItemPrice.CreatedBy,
				ModifiedBy:    iv.ItemPrice.ModifiedBy,
			}
			status := responses.Status{
				StatusId:     strconv.FormatInt(iv.Status.StatusId, 10),
				Status:       iv.Status.Status,
				StatusCode:   iv.Status.StatusCode,
				DateCreated:  iv.Status.DateCreated,
				DateModified: iv.Status.DateModified,
				Active:       iv.Status.Active,
			}
			quantity := responses.Item_quantity{
				ItemQuantityId: strconv.FormatInt(iv.ItemQuantity.ItemQuantityId, 10),
				Quantity:       iv.ItemQuantity.Quantity,
				QuantityAlert:  iv.ItemQuantity.QuantityAlert,
				Active:         iv.ItemQuantity.Active,
				DateCreated:    iv.ItemQuantity.DateCreated,
				DateModified:   iv.ItemQuantity.DateModified,
				CreatedBy:      iv.ItemQuantity.CreatedBy,
				ModifiedBy:     iv.ItemQuantity.ModifiedBy,
			}
			features := []*responses.Features{}
			for _, f := range iv.ItemFeatures {
				features = append(features, &responses.Features{
					FeatureId:    strconv.FormatInt(f.FeatureId, 10),
					Feature:      f.FeatureName,
					Description:  f.Description,
					DateCreated:  f.DateCreated,
					DateModified: f.DateModified,
					ImagePath:    f.ImagePath,
					Active:       f.Active,
				})
			}

			purposes := []*responses.Purposes{}
			for _, p := range iv.ItemPurposes {
				purposes = append(purposes, &responses.Purposes{
					PurposeId:    strconv.FormatInt(p.PurposeId, 10),
					Purpose:      p.Purpose,
					Description:  p.Description,
					DateCreated:  p.DateCreated,
					DateModified: p.DateModified,
					ImagePath:    p.ImagePath,
					Active:       p.Active,
				})
			}
			respData := responses.Items{
				ItemId:          strconv.FormatInt(iv.ItemId, 10),
				ItemName:        iv.ItemName,
				Description:     iv.Description,
				Weight:          iv.Weight,
				Category:        &categoryData,
				ItemPrice:       &price,
				AvailableSizes:  iv.AvailableSizes,
				AvailableColors: iv.AvailableColors,
				Material:        iv.Material,
				ImagePath:       iv.ImagePath,
				Quantity:        iv.Quantity,
				Active:          iv.Active,
				DateCreated:     iv.DateCreated,
				DateModified:    iv.DateModified,
				CreatedBy:       iv.CreatedBy,
				ModifiedBy:      iv.ModifiedBy,
				Country:         strconv.FormatInt(iv.Country, 10),
				Branch:          strconv.FormatInt(iv.Branch, 10),
				Status:          &status,
				LastOrderDate:   iv.LastOrderDate,
				ItemQuantity:    &quantity,
				ItemFeatures:    features,
				ItemPurposes:    purposes,
			}

			resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item successfully updated"}
			c.Data["json"] = resp

		} else {
			resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// UpdateItemQuantity ...
// @Title Update Item Quantity
// @Description update the Item's quantity
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ItemQuantityRequest	true		"body for Items content"
// @Success 200 {object} models.Items
// @Failure 403 :id is not int
// @router /quantity/:id [put]
func (c *ItemsController) UpdateItemQuantity() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var t requests.ItemQuantityRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)
	logs.Info("Request received. Quantity is ", t.Quantity, " Item id is ", idStr)

	iv, err := models.GetItemsById(id)
	if err != nil {
		resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		iv.Quantity = int(t.Quantity)

		if err := models.UpdateItemsById(iv); err == nil {
			// Add quantity for item

			if iq, err := models.GetItem_quantityByItemId(id); err == nil {
				iq.Quantity = t.Quantity
				if err := models.UpdateItem_quantityById(iq); err == nil {
					logs.Info("Item quantity updated successfully")
				} else {
					logs.Error("Error updating item quantity: ", err.Error())
				}
			} else {
				logs.Error("Error fetching item quantity: ", err.Error())
			}
			c.Ctx.Output.SetStatus(200)

			categoryData := responses.Categories{
				CategoryId:   strconv.FormatInt(iv.Category.CategoryId, 10),
				CategoryName: iv.Category.CategoryName,
				Description:  iv.Category.Description,
				ImagePath:    iv.Category.ImagePath,
				Active:       iv.Category.Active,
				Icon:         iv.Category.Icon,
				DateCreated:  iv.Category.DateCreated,
				DateModified: iv.Category.DateModified,
			}
			price := responses.Item_prices{
				ItemPriceId:   strconv.FormatInt(iv.ItemPrice.ItemPriceId, 10),
				ItemPrice:     iv.ItemPrice.ItemPrice,
				AltItemPrice:  iv.ItemPrice.AltItemPrice,
				ShowAltPrice:  iv.ItemPrice.ShowAltPrice,
				Discount:      iv.ItemPrice.Discount,
				Discount_type: iv.ItemPrice.Discount_type,
				ExtraCharges:  iv.ItemPrice.ExtraCharges,
				Currency:      iv.ItemPrice.Currency,
				Active:        iv.ItemPrice.Active,
				DateCreated:   iv.ItemPrice.DateCreated,
				DateModified:  iv.ItemPrice.DateModified,
				CreatedBy:     iv.ItemPrice.CreatedBy,
				ModifiedBy:    iv.ItemPrice.ModifiedBy,
			}
			status := responses.Status{
				StatusId:     strconv.FormatInt(iv.Status.StatusId, 10),
				Status:       iv.Status.Status,
				StatusCode:   iv.Status.StatusCode,
				DateCreated:  iv.Status.DateCreated,
				DateModified: iv.Status.DateModified,
				Active:       iv.Status.Active,
			}
			quantity := responses.Item_quantity{
				ItemQuantityId: strconv.FormatInt(iv.ItemQuantity.ItemQuantityId, 10),
				Quantity:       iv.ItemQuantity.Quantity,
				QuantityAlert:  iv.ItemQuantity.QuantityAlert,
				Active:         iv.ItemQuantity.Active,
				DateCreated:    iv.ItemQuantity.DateCreated,
				DateModified:   iv.ItemQuantity.DateModified,
				CreatedBy:      iv.ItemQuantity.CreatedBy,
				ModifiedBy:     iv.ItemQuantity.ModifiedBy,
			}
			features := []*responses.Features{}
			for _, f := range iv.ItemFeatures {
				features = append(features, &responses.Features{
					FeatureId:    strconv.FormatInt(f.FeatureId, 10),
					Feature:      f.FeatureName,
					Description:  f.Description,
					DateCreated:  f.DateCreated,
					DateModified: f.DateModified,
					ImagePath:    f.ImagePath,
					Active:       f.Active,
				})
			}

			purposes := []*responses.Purposes{}
			for _, p := range iv.ItemPurposes {
				purposes = append(purposes, &responses.Purposes{
					PurposeId:    strconv.FormatInt(p.PurposeId, 10),
					Purpose:      p.Purpose,
					Description:  p.Description,
					DateCreated:  p.DateCreated,
					DateModified: p.DateModified,
					ImagePath:    p.ImagePath,
					Active:       p.Active,
				})
			}
			respData := responses.Items{
				ItemId:          strconv.FormatInt(iv.ItemId, 10),
				ItemName:        iv.ItemName,
				Description:     iv.Description,
				Weight:          iv.Weight,
				Category:        &categoryData,
				ItemPrice:       &price,
				AvailableSizes:  iv.AvailableSizes,
				AvailableColors: iv.AvailableColors,
				Material:        iv.Material,
				ImagePath:       iv.ImagePath,
				Quantity:        iv.Quantity,
				Active:          iv.Active,
				DateCreated:     iv.DateCreated,
				DateModified:    iv.DateModified,
				CreatedBy:       iv.CreatedBy,
				ModifiedBy:      iv.ModifiedBy,
				Country:         strconv.FormatInt(iv.Country, 10),
				Branch:          strconv.FormatInt(iv.Branch, 10),
				Status:          &status,
				LastOrderDate:   iv.LastOrderDate,
				ItemQuantity:    &quantity,
				ItemFeatures:    features,
				ItemPurposes:    purposes,
			}
			resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item successfully updated"}
			c.Data["json"] = resp

		} else {
			resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// UpdateItemPrice ...
// @Title Update Item Price
// @Description update the Item's price
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.ItemPriceRequest	true		"body for Items content"
// @Success 200 {object} models.Items
// @Failure 403 :id is not int
// @router /price/:id [put]
func (c *ItemsController) UpdateItemPrice() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	var t requests.ItemPriceRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)
	logs.Info("Request received. Price is ", t.Price, " Item id is ", idStr)

	iv, err := models.GetItemsById(id)
	var respData responses.Items
	if err != nil {
		resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		iv.ItemPrice.ItemPrice = t.Price
		iv.ItemPrice.AltItemPrice = t.AltPrice
		iv.ItemPrice.ExtraCharges = t.ExtraCharges
		respData = responses.Items{
			ItemId:          strconv.FormatInt(iv.ItemId, 10),
			ItemName:        iv.ItemName,
			Description:     iv.Description,
			Weight:          iv.Weight,
			Category:        nil,
			ItemPrice:       nil,
			AvailableSizes:  iv.AvailableSizes,
			AvailableColors: iv.AvailableColors,
			Material:        iv.Material,
			ImagePath:       iv.ImagePath,
			Quantity:        iv.Quantity,
			Active:          iv.Active,
			DateCreated:     iv.DateCreated,
			DateModified:    iv.DateModified,
			CreatedBy:       iv.CreatedBy,
			ModifiedBy:      iv.ModifiedBy,
			Country:         strconv.FormatInt(iv.Country, 10),
			Branch:          strconv.FormatInt(iv.Branch, 10),
			Status:          nil,
			LastOrderDate:   iv.LastOrderDate,
			ItemQuantity:    nil,
			ItemFeatures:    []*responses.Features{},
			ItemPurposes:    []*responses.Purposes{},
		}

		if err := models.UpdateItemsById(iv); err == nil {
			// Update price for item

			c.Ctx.Output.SetStatus(200)

			resp := responses.ItemResponseDTO{StatusCode: 200, Item: &respData, StatusDesc: "Item successfully updated"}
			c.Data["json"] = resp

		} else {
			resp := responses.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	}

	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Items
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ItemsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)

	i, ierr := models.GetItemsById(id)

	statusCode := 200
	message := "Item successfully deleted"

	if ierr == nil {
		logs.Info("Item retrieved ")
		logs.Info(i)

		if ii, iierr := models.GetItem_imagesByItemId(id); iierr == nil {
			logs.Info("Item images returned are ")
			logs.Info(ii)
			for _, ib := range *ii {
				// if imerr := models.DeleteItem_images(ib.ItemImageId); imerr == nil {
				// 	logs.Info("Item images deleted")
				// 	logs.Info("Item is ", i.ItemName)
				// } else {
				// 	panic(imerr)
				// }
				ib.Active = 0
				if imerr := models.UpdateItem_imagesById(&ib); imerr == nil {
					logs.Info("Item images deactivated")
					logs.Info("Item is ", i.ItemName)
				} else {
					panic(imerr)
				}
			}
		} else {
			logs.Error("An error occurred")
			logs.Error(iierr)
		}

		q, gerr := models.GetItem_quantityByItemId(id)

		if gerr == nil {
			logs.Info("Quantity ID is ", q.ItemQuantityId)
			// if qerr := models.DeleteItem_quantity(q.ItemQuantityId); qerr == nil {
			// 	logs.Info("Quantity deleted ")
			// } else {
			// 	panic(qerr)
			// }
			q.Active = 0
			if qerr := models.UpdateItem_quantityById(q); qerr == nil {
				logs.Info("Quantity deactivated ")
			} else {
				panic(qerr)
			}
		} else {
			logs.Error("An error occurred")
			logs.Error(gerr)
		}

		logs.Info("Deleting item features and purposes ", i.ItemId)

		// if qerr := models.DeleteItem_featuresByItem(i.ItemId); qerr == nil {
		// 	logs.Info("Item feature deleted ")
		// } else {
		// 	logs.Error("No item features to delete ")
		// 	logs.Error(qerr)
		// 	// panic(qerr)
		// }
		itf, ferr := models.GetItem_featuresByItemId(i.ItemId)
		if ferr == nil {
			for _, ib := range *itf {
				// if imerr := models.DeleteItem_images(ib.ItemImageId); imerr == nil {
				// 	logs.Info("Item images deleted")
				// 	logs.Info("Item is ", i.ItemName)
				// } else {
				// 	panic(imerr)
				// }
				ib.Active = 0
				if imerr := models.UpdateItem_featuresById(&ib); imerr == nil {
					logs.Info("Item feature deactivated")
					logs.Info("Item is ", i.ItemName)
				} else {
					panic(imerr)
				}
			}
		} else {
			logs.Error("An error occurred")
			logs.Error(ferr)
		}

		// if qerr := models.DeleteItem_purposesByItem(i.ItemId); qerr == nil {
		// 	logs.Info("Item purpose deleted ")
		// } else {
		// 	logs.Error("No item purposes to delete ")
		// 	logs.Error(qerr)
		// 	// panic(qerr)
		// }
		itp, perr := models.GetItem_purposesByItemId(i.ItemId)
		if perr == nil {
			for _, ib := range *itp {
				// if imerr := models.DeleteItem_images(ib.ItemImageId); imerr == nil {
				// 	logs.Info("Item images deleted")
				// 	logs.Info("Item is ", i.ItemName)
				// } else {
				// 	panic(imerr)
				// }
				ib.Active = 0
				if imerr := models.UpdateItem_purposesById(&ib); imerr == nil {
					logs.Info("Item purpose deactivated")
					logs.Info("Item is ", i.ItemName)
				} else {
					panic(imerr)
				}
			}
		} else {
			logs.Error("An error occurred")
			logs.Error(perr)
		}

		// Finally delete item
		i.Active = 0
		if qerr := models.UpdateItemsById(i); qerr == nil {
			logs.Info("Deactivating Item: ", i.ItemPrice.ItemPriceId)
			logs.Info("Item Deleted ", id)
			statusCode = 200
			message = "Item successfully deactivated"
		} else {
			panic(qerr)
			statusCode = 301
			message = qerr.Error()
		}

		// if err := models.DeleteItems(id); err == nil {
		// 	logs.Error("Item Deleted ", id)
		// if qerr := models.DeleteItem_prices(i.ItemPrice.ItemPriceId); qerr == nil {
		// 	logs.Info("Deleting Item price: ", i.ItemPrice.ItemPriceId)
		// 	logs.Info("Item Deleted ", id)
		// 	c.Data["json"] = "OK"
		// } else {
		// 	panic(qerr)
		// }
		// } else {
		// 	c.Data["json"] = err.Error()
		// }

	} else {
		logs.Error("An error occurred")
		logs.Error(ierr)
		statusCode = 301
		message = ierr.Error()
	}

	resp := responses.StringResponseDTO{StatusCode: statusCode, Value: "OK", StatusDesc: message}
	c.Data["json"] = resp

	c.ServeJSON()
}

// GetItemStats ...
// @Title Get Item Stats
// @Description get item stats
// @Param	branch_id		path 	string	true		"The id you want to update"
// @Success 200 {object} responses.ItemsStatsResponseDTO
// @Failure 403 wrong request
// @router /get-item-stats/:branch_id [get]
func (c *ItemsController) GetItemStats() {
	logs.Info("Getting item stats")
	branchidStr := c.Ctx.Input.Param(":branch_id")
	branchid, _ := strconv.ParseInt(branchidStr, 0, 64)

	stats := responses.StatsDTO{}
	itemCategoryStats := []models.ItemsCategoryCountDTO{}
	itemBranchStats := []models.ItemsCategoryCountDTO{}

	if r, err := models.GetItemsStatsByCategory(); err == nil {
		itemCategoryStats = *r
	}

	stats.CategoryStats = &itemCategoryStats

	if r, err := models.GetItemsCategoryStatsByBranch(branchid); err == nil {
		// for _, a := range *r {
		// 	itemCategoryStats = append(itemCategoryStats, a)
		// }
		// itemCategoryStats = append(itemCategoryStats, *r...)

		itemBranchStats = *r
	}

	stats.BranchStats = &itemBranchStats

	resp := responses.ItemsStatsResponseDTO{StatusCode: 200, Stats: &stats, StatusDesc: "Successfully fetched stats"}
	c.Data["json"] = resp

	c.ServeJSON()
}

// CheckItemQuantity ...
// @Title Check Item Quantity
// @Description check item quantity
// @Param	item_id		path 	string	true		"The id you want to update"
// @Success 200 {object} responses.StringResponseFDTO
// @Failure 403 wrong request
// @router /check-item-quantity/:item_id [get]
func (c *ItemsController) CheckItemQuantity() {
	logs.Info("Checking item quantity")
	itemidStr := c.Ctx.Input.Param(":item_id")
	itemid, _ := strconv.ParseInt(itemidStr, 0, 64)

	if item, err := models.GetItemsById(itemid); err == nil {
		iResp := functions.CheckItemCount(item.ItemId, item.ItemName)

		if iResp {
			resp := responses.StringResponseDTO{StatusCode: 200, Value: "LOW", StatusDesc: "Item quantity low"}
			c.Data["json"] = resp
		} else {
			resp := responses.StringResponseDTO{StatusCode: 200, Value: "OK", StatusDesc: "Item quantity ok"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Error getting item ", err.Error())

		resp := responses.StringResponseDTO{StatusCode: 608, Value: "ERROR", StatusDesc: "Notification send failed"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}
