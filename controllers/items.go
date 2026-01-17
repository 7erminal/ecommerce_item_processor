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
	c.Mapping("GetItemFeatures", c.GetItemFeatures)
	c.Mapping("GetItemPurposes", c.GetItemPurposes)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetItemsByCategory", c.GetItemsByCategory)
	c.Mapping("UpdateItemImage", c.UpdateItemImage)
	c.Mapping("GetItemStats", c.GetItemStats)
	c.Mapping("GetAllByBranch", c.GetAllByBranch)
	c.Mapping("GetItemCountByType", c.GetItemCountByType)
	c.Mapping("GetItemCount", c.GetItemCount)
	c.Mapping("CheckItemQuantity", c.CheckItemQuantity)
	c.Mapping("GetItemCountWithTypeAndBranch", c.GetItemCountWithTypeAndBranch)
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

	p, err := models.GetCategoriesById(int64(t.Category))

	if err == nil {
		// Get Currency
		cr, cerr := models.GetCurrenciesById(int64(1))

		if cerr == nil {
			// Add price for item
			logs.Info("Adding price for item with price ", t.ItemPrice, " and alt price ", t.AltItemPrice, " and extra charges ", t.ExtraCharges)
			it := models.Item_prices{ItemPrice: t.ItemPrice, AltItemPrice: t.AltItemPrice, ShowAltPrice: false, ExtraCharges: t.ExtraCharges, Currency: cr, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}

			logs.Info("Adding price to item to create at a go")
			if _, err := models.AddItem_prices(&it); err == nil {
				country, _ := models.GetCountriesByCode(t.Country)
				branch, _ := models.GetBranchesById(t.Branch)

				status := models.Status{}

				if status_, err := models.GetStatusByName("GOOD"); err == nil {
					status = *status_
				}

				// Add item if getting category and price addition does not result in an error
				v := models.Items{ItemName: t.ItemName, Description: t.Description, Weight: t.Weight, Category: p, ItemPrice: &it, AvailableSizes: aSizes, AvailableColors: aColors, Quantity: t.Quantity, Country: country, Branch: branch, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: creator, ModifiedBy: creator, Status: &status}

				if _, err := models.AddItems(&v); err == nil {
					// Add quantity for item
					qu := models.Item_quantity{Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}

					if _, err := models.AddItem_quantity(&qu); err == nil {
						c.Ctx.Output.SetStatus(200)
						logs.Info("Item added successfully with ID ", v.ItemId)

						item, err := models.GetItemsById(v.ItemId)
						if err != nil {

							logs.Error(err.Error())
							resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
							c.Data["json"] = resp
						} else {
							logs.Info("Item fetched successfully")
							logs.Info("Item quantity is ", item.ItemQuantity)
							logs.Info("Item quantity added successfully")
						}

						errorCode = 200
						message = "Item added successfully with ID " + strconv.FormatInt(v.ItemId, 10)

						resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: item, StatusDesc: message}
						c.Data["json"] = resp
					} else {
						logs.Error(err.Error())
						message = "Error adding item quantity: " + err.Error()
						resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: &v, StatusDesc: message}
						c.Data["json"] = resp
					}

				} else {
					logs.Error(err.Error())
					errorCode = 301
					message = "Error adding item: " + err.Error()
					resp := responses.ItemResponseDTO{StatusCode: errorCode, Item: &v, StatusDesc: message}
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
		resp := models.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemResponseDTO{StatusCode: 200, Item: v, StatusDesc: "Item fetched successfully"}
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
		resp := models.ItemQuantityResponseDTO{StatusCode: 301, Quantity: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemQuantityResponseDTO{StatusCode: 200, Quantity: v, StatusDesc: "Quantity fetched successfully"}
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

// // GetItemFeaturesByItem ...
// // @Title Get Item Features by item
// // @Description get Item_features by Item id
// // @Param	id		path 	string	true		"The key for staticblock"
// // @Success 200 {object} models.ItemsResponseDTO2
// // @Failure 403 :id is empty
// // @router /features/item/:id [get]
// func (c *ItemsController) GetItemFeaturesByItem() {
// 	idStr := c.Ctx.Input.Param(":id")
// 	id, _ := strconv.ParseInt(idStr, 0, 64)
// 	v, err := models.GetItemsWithItem(id)
// 	if err != nil {
// 		resp := models.ItemsResponseDTO2{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
// 		c.Data["json"] = resp
// 	} else {
// 		logs.Info("Item features fetched are ", v)
// 		resp := models.ItemsResponseDTO2{StatusCode: 200, Items: v, StatusDesc: "Features fetched successfully"}
// 		c.Data["json"] = resp
// 	}
// 	c.ServeJSON()
// }

// GetItemFeatures ...
// @Title Get Item Features
// @Description get Item_features by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemsResponseDTO2
// @Failure 403 :id is empty
// @router /features/feature/:id [get]
func (c *ItemsController) GetItemFeatures() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItemsByFeatureId(id)
	if err != nil {
		resp := models.ItemsResponseDTO2{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Item features fetched are ", v)
		resp := models.ItemsResponseDTO2{StatusCode: 200, Items: v, StatusDesc: "Features fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemsByCategory ...
// @Title Get Item Features
// @Description get Item_features by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemsResponseDTO2
// @Failure 403 :id is empty
// @router /categories/:id [get]
func (c *ItemsController) GetItemsByCategory() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItemsByCategoryId(id)
	if err != nil {
		resp := models.ItemsResponseDTO2{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Items fetched are ", v)
		resp := models.ItemsResponseDTO2{StatusCode: 200, Items: v, StatusDesc: "Items fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemPurposes ...
// @Title Get Item Purposes
// @Description get Item_purposes by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemsResponseDTO2
// @Failure 403 :id is empty
// @router /purposes/purpose/:id [get]
func (c *ItemsController) GetItemPurposes() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItemsByPurposeId(id)
	if err != nil {
		resp := models.ItemsResponseDTO2{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemsResponseDTO2{StatusCode: 200, Items: v, StatusDesc: "Purposes fetched successfully"}
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

	activeQueryt := "Test:1"
	if v := activeQueryt; v != "" {
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
		resp := models.ItemsResponseDTO{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		if l == nil {
			l = []interface{}{}
		}
		logs.Info("Items returned are ", l)
		resp := models.ItemsResponseDTO{StatusCode: 200, Items: &l, StatusDesc: "Items fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllByBranch ...
// @Title Get All Items by branch
// @Description get Items
// @Param	branch_id		path 	string	true		"The id you want to update"
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Items
// @Failure 403
// @router /branch/:branch_id [get]
func (c *ItemsController) GetAllByBranch() {
	branchidStr := c.Ctx.Input.Param(":branch_id")
	branchid, _ := strconv.ParseInt(branchidStr, 0, 64)
	var fields []string
	var sortby []string
	var order []string
	var query = make(map[string]string)
	var limit int64 = 10
	var offset int64

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

	if branch, err := models.GetBranchesById(branchid); err == nil {
		l, err := models.GetAllItemsByBranch(branch, query, fields, sortby, order, offset, limit)
		if err != nil {
			resp := models.ItemsResponseDTO{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		} else {
			resp := models.ItemsResponseDTO{StatusCode: 200, Items: &l, StatusDesc: "Items fetched successfully"}
			c.Data["json"] = resp
		}
	} else {
		logs.Error("Branch does not exist")
		resp := models.ItemsResponseDTO{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
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
	var t models.ItemsDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)

	creator := t.CreatedBy

	// Structure Available Sizes
	aSizes := strings.Join(t.AvailableSizes, ",")

	// Structure Available Colors
	aColors := strings.Join(t.AvailableColors, ",")

	p, err := models.GetCategoriesById(int64(t.Category))

	if err == nil {
		iv, err := models.GetItemsById(id)
		if err != nil {
			resp := models.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		} else {
			// Get Currency
			cr, cerr := models.GetCurrenciesById(int64(1))

			if cerr == nil {
				if ip, err := models.GetItem_pricesById(iv.ItemPrice.ItemPriceId); err == nil {
					// Add price for item
					it := models.Item_prices{ItemPriceId: iv.ItemPrice.ItemPriceId, ItemPrice: t.ItemPrice, AltItemPrice: t.AltPrice, ShowAltPrice: false, ExtraCharges: t.ExtraCharges, Currency: cr, Active: 1, ModifiedBy: creator, DateCreated: ip.DateCreated, CreatedBy: ip.CreatedBy, DateModified: time.Now()}

					logs.Info("Modifying price for item")

					if err := models.UpdateItem_pricesById(&it); err == nil {
						// Add item if getting category and price addition does not result in an error
						country, _ := models.GetCountriesByCode(t.Country)
						branch, _ := models.GetBranchesById(t.Branch)
						v := models.Items{ItemId: id, ItemName: t.ItemName, Country: country, Branch: branch, Description: t.Description, Category: p, ImagePath: iv.ImagePath, ItemPrice: &it, AvailableSizes: aSizes, AvailableColors: aColors, Quantity: t.Quantity, Active: 1, DateModified: time.Now(), ModifiedBy: creator, CreatedBy: iv.CreatedBy, DateCreated: iv.DateCreated, Weight: t.Weight}

						if err := models.UpdateItemsById(&v); err == nil {
							// Add quantity for item

							iq, err := models.GetItem_quantityByItemId(id)
							if err != nil {
								resp := models.ItemResponseDTO{StatusCode: 304, Item: &v, StatusDesc: "Item quantity not set"}
								c.Data["json"] = resp
								qu := models.Item_quantity{Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}
								if _, err := models.AddItem_quantity(&qu); err == nil {
									item, err := models.GetItemsById(v.ItemId)
									if err != nil {

										logs.Error(err.Error())
										resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
										c.Data["json"] = resp
									} else {
										logs.Info("Item fetched successfully")
										logs.Info("Item quantity is ", item.ItemQuantity)
										logs.Info("Item quantity added successfully")
									}
									c.Ctx.Output.SetStatus(200)

									resp := responses.ItemResponseDTO{StatusCode: 200, Item: item, StatusDesc: "Item successfully added"}
									c.Data["json"] = resp
								} else {
									logs.Error(err.Error())
									resp := responses.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
									c.Data["json"] = resp
								}
							} else {
								qu := models.Item_quantity{ItemQuantityId: iq.ItemQuantityId, Item: &v, Quantity: t.Quantity, QuantityAlert: t.QuantityAlert, Active: 1, CreatedBy: iq.CreatedBy, DateCreated: iq.DateCreated, ModifiedBy: creator, DateModified: time.Now()}

								if err := models.UpdateItem_quantityById(&qu); err == nil {
									item, err := models.GetItemsById(v.ItemId)
									if err != nil {

										logs.Error(err.Error())
										resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
										c.Data["json"] = resp
									} else {
										logs.Info("Item fetched successfully")
										logs.Info("Item quantity is ", item.ItemQuantity)
										logs.Info("Item quantity added successfully")
									}

									c.Ctx.Output.SetStatus(200)

									resp := models.ItemResponseDTO{StatusCode: 200, Item: item, StatusDesc: "Item successfully updated"}
									c.Data["json"] = resp
								} else {
									logs.Error(err.Error())
									resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
									c.Data["json"] = resp
								}
							}

						} else {
							logs.Error(err.Error())
							resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
							c.Data["json"] = resp
						}
					} else {
						logs.Error(err.Error())
						resp := models.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
						c.Data["json"] = resp
					}
				} else {
					logs.Error(err.Error())
					resp := models.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
					c.Data["json"] = resp
				}
			} else {
				resp := models.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
				c.Data["json"] = resp
			}
		}

	} else {
		logs.Error(err.Error())
		resp := models.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
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
		resp := models.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		iv.ImagePath = t.ImagePath

		if err := models.UpdateItemsById(iv); err == nil {
			// Add quantity for item

			c.Ctx.Output.SetStatus(200)

			resp := models.ItemResponseDTO{StatusCode: 200, Item: iv, StatusDesc: "Item successfully updated"}
			c.Data["json"] = resp

		} else {
			resp := models.ItemResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
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

		if err := models.DeleteItems(id); err == nil {
			logs.Error("Item Deleted ", id)
			// if qerr := models.DeleteItem_prices(i.ItemPrice.ItemPriceId); qerr == nil {
			// 	logs.Info("Deleting Item price: ", i.ItemPrice.ItemPriceId)
			// 	logs.Info("Item Deleted ", id)
			// 	c.Data["json"] = "OK"
			// } else {
			// 	panic(qerr)
			// }
			i.Active = 0
			if qerr := models.UpdateItemsById(i); qerr == nil {
				logs.Info("Deactivating Item: ", i.ItemPrice.ItemPriceId)
				logs.Info("Item Deleted ", id)
				c.Data["json"] = "OK"
			} else {
				panic(qerr)
			}
		} else {
			c.Data["json"] = err.Error()
		}

	} else {
		logs.Error("An error occurred")
		logs.Error(ierr)
	}

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
