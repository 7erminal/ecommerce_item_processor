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

// Item_featuresController operations for Item_features
type Item_featuresController struct {
	beego.Controller
}

// URLMapping ...
func (c *Item_featuresController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Item_features
// @Param	body		body 	models.ItemFeatureRequestDTO	true		"body for Item_features content"
// @Success 200 {int} models.ItemFeatureResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Item_featuresController) Post() {
	var v models.ItemFeatureRequestDTO
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	itemid, _ := strconv.ParseInt(v.ItemId, 0, 64)
	featureid, _ := strconv.ParseInt(v.FeatureId, 0, 64)

	if it, err := models.GetItemsById(itemid); err == nil {
		if ft, err := models.GetFeaturesById(featureid); err == nil {
			iff := models.Item_features{Item: it, Feature: ft, Active: 1, DateCreated: time.Now(), DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1}

			if _, err := models.AddItem_features(&iff); err == nil {
				c.Ctx.Output.SetStatus(200)

				resp := models.ItemFeatureResponseDTO{StatusCode: 200, ItemFeature: &iff, StatusDesc: "Item feature added successfully"}
				c.Data["json"] = resp
			} else {
				logs.Error(err.Error())
				resp := models.ItemFeatureResponseDTO{StatusCode: 301, ItemFeature: nil, StatusDesc: err.Error()}
				c.Data["json"] = resp
			}
		} else {
			logs.Error(err.Error())
			resp := models.ItemFeatureResponseDTO{StatusCode: 301, ItemFeature: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	} else {
		logs.Error(err.Error())
		resp := models.ItemFeatureResponseDTO{StatusCode: 301, ItemFeature: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Item_features by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Item_features
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Item_featuresController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_featuresById(id)
	if err != nil {
		resp := models.ItemFeatureResponseDTO{StatusCode: 200, ItemFeature: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	} else {
		resp := models.ItemFeatureResponseDTO{StatusCode: 200, ItemFeature: v, StatusDesc: "Item feature fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Item_features
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Item_features
// @Failure 403
// @router / [get]
func (c *Item_featuresController) GetAll() {
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

	l, err := models.GetAllItem_features(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Item_features
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Item_features	true		"body for Item_features content"
// @Success 200 {object} models.Item_features
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Item_featuresController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Item_features{ItemFeatureId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateItem_featuresById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Item_features
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Item_featuresController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteItem_features(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
