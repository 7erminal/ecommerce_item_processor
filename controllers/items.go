package controllers

import (
	"encoding/json"
	"errors"
	"item_processor/models"
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
	c.Mapping("GetItemFeatures", c.GetItemFeatures)
	c.Mapping("GetItemPurposes", c.GetItemPurposes)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Items
// @Param	body		body 	models.ItemsDTO	true		"body for Items content"
// @Success 201 {int} models.ItemsResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *ItemsController) Post() {
	var t models.ItemsDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &t)
	logs.Info("Received ", t)

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
			it := models.Item_prices{ItemPrice: t.ItemPrice, AltItemPrice: 0, ShowAltPrice: false, Currency: cr, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}

			logs.Info("Adding price to item to create at a go")
			if _, err := models.AddItem_prices(&it); err == nil {
				// Add item if getting category and price addition does not result in an error
				v := models.Items{ItemName: t.ItemName, Description: t.Description, Category: p, ItemPrice: &it, AvailableSizes: aSizes, AvailableColors: aColors, Quantity: t.Quantity, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: creator, ModifiedBy: creator}

				if _, err := models.AddItems(&v); err == nil {
					// Add quantity for item
					qu := models.Item_quantity{ItemId: v.ItemId, Quantity: t.Quantity, Active: 1, CreatedBy: creator, DateCreated: time.Now(), ModifiedBy: creator, DateModified: time.Now()}

					if _, err := models.AddItem_quantity(&qu); err == nil {
						c.Ctx.Output.SetStatus(200)

						resp := models.ItemResponseDTO{StatusCode: 200, Item: &v, StatusDesc: "Item successfully added"}
						c.Data["json"] = resp
					} else {
						logs.Error(err.Error())
						resp := models.ItemResponseDTO{StatusCode: 302, Item: &v, StatusDesc: err.Error()}
						c.Data["json"] = resp
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
			// resp := models.ItemsResponseDTO{StatusCode: 302, Item: nil, StatusDesc: err.Error()}
			// c.Data["json"] = resp
			logs.Error(err.Error())
			c.Data["json"] = err.Error()
		}

	} else {
		logs.Error(err.Error())
		resp := models.ItemResponseDTO{StatusCode: 301, Item: nil, StatusDesc: err.Error()}
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
	v, err := models.GetItem_quantityByItemId(id)
	if err != nil {
		resp := models.ItemQuantityResponseDTO{StatusCode: 301, Quantity: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemQuantityResponseDTO{StatusCode: 200, Quantity: v, StatusDesc: "Quantity fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemFeatures ...
// @Title Get Item Features
// @Description get Item_features by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemFeaturesResponseDTO
// @Failure 403 :id is empty
// @router /features/:id [get]
func (c *ItemsController) GetItemFeatures() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_featuresByItemId(id)
	if err != nil {
		resp := models.ItemFeaturesResponseDTO{StatusCode: 301, ItemFeatures: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		logs.Info("Item features fetched are ", v)
		resp := models.ItemFeaturesResponseDTO{StatusCode: 200, ItemFeatures: v, StatusDesc: "Features fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetItemPurposes ...
// @Title Get Item Purposes
// @Description get Item_purposes by Item id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ItemPurposesResponseDTO
// @Failure 403 :id is empty
// @router /purposes/:id [get]
func (c *ItemsController) GetItemPurposes() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_purposesByItemId(id)
	if err != nil {
		resp := models.ItemPurposesResponseDTO{StatusCode: 301, ItemPurposes: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemPurposesResponseDTO{StatusCode: 200, ItemPurposes: v, StatusDesc: "Purposes fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Items
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
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

	l, err := models.GetAllItems(query, fields, sortby, order, offset, limit)
	if err != nil {
		resp := models.ItemsResponseDTO{StatusCode: 301, Items: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemsResponseDTO{StatusCode: 200, Items: &l, StatusDesc: "Items fetched successfully"}
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
				// Add price for item
				it := models.Item_prices{ItemPriceId: iv.ItemPrice.ItemPriceId, ItemPrice: t.ItemPrice, AltItemPrice: 0, ShowAltPrice: false, Currency: cr, Active: 1, ModifiedBy: creator, DateModified: time.Now()}

				logs.Info("Modifying price for item")

				if err := models.UpdateItem_pricesById(&it); err == nil {
					// Add item if getting category and price addition does not result in an error
					v := models.Items{ItemId: id, ItemName: t.ItemName, Description: t.Description, Category: p, ItemPrice: &it, AvailableSizes: aSizes, AvailableColors: aColors, Quantity: t.Quantity, Active: 1, DateModified: time.Now(), ModifiedBy: creator}

					if err := models.UpdateItemsById(&v); err == nil {
						// Add quantity for item

						iq, err := models.GetItem_quantityByItemId(id)
						if err != nil {
						} else {
							qu := models.Item_quantity{ItemQuantityId: iq.ItemQuantityId, ItemId: v.ItemId, Quantity: t.Quantity, Active: 1, ModifiedBy: creator, DateModified: time.Now()}

							if err := models.UpdateItem_quantityById(&qu); err == nil {
								c.Ctx.Output.SetStatus(200)

								resp := models.ItemResponseDTO{StatusCode: 200, Item: &v, StatusDesc: "Item successfully updated"}
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

	q, gerr := models.GetItem_quantityByItemId(id)

	if gerr == nil {
		logs.Info("Quantity ID is ", q.ItemQuantityId)
		if qerr := models.DeleteItem_quantity(q.ItemQuantityId); qerr == nil {
			logs.Info("Quantity deleted ")
		} else {
			panic(qerr)
		}
	}

	i, ierr := models.GetItemsById(id)

	if ierr == nil {
		logs.Info("Item retrieved ")
		if ii, iierr := models.GetItem_imagesByItemId(id); iierr == nil {
			logs.Info("Item images returned are ")
			logs.Info(ii)
			for _, ib := range *ii {
				if imerr := models.DeleteItem_images(ib.ItemImageId); imerr == nil {
					logs.Info("Item images deleted")
					logs.Info("Item is ", i.ItemName)
					if err := models.DeleteItems(id); err == nil {
						if qerr := models.DeleteItem_prices(i.ItemPrice.ItemPriceId); qerr == nil {
							logs.Info("Deleting Item price: ", i.ItemPrice.ItemPriceId)
							logs.Info("Item Deleted ", id)
							c.Data["json"] = "OK"
						} else {
							panic(qerr)
						}
					} else {
						c.Data["json"] = err.Error()
					}
				} else {
					panic(imerr)
				}
			}
		}
	}

	c.ServeJSON()
}
