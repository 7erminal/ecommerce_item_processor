package controllers

import (
	"encoding/json"
	"errors"
	"item_processor/models"
	"item_processor/structs/requests"
	"item_processor/structs/responses"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"

	beego "github.com/beego/beego/v2/server/web"
)

// Item_reviewsController operations for Item_reviews
type Item_reviewsController struct {
	beego.Controller
}

// URLMapping ...
func (c *Item_reviewsController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Item_reviews
// @Param	body		body 	requests.AddReviewRequest	true		"body for Item_reviews content"
// @Success 201 {int} responses.ItemReviewResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Item_reviewsController) Post() {
	var v requests.AddReviewRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	if u, err := models.GetUsersById(v.ReviewBy); err == nil {
		item := models.Items{}
		if item_, err := models.GetItemsById(v.ItemId); err == nil {
			// logs.Error("Error returned after attempting to add review is ", err.Error())
			item = *item_
		}
		var r models.Item_reviews = models.Item_reviews{Review: v.Review, Item: &item, Reference: int(v.Reference), Rating: v.Rating, ReviewBy: u, Active: 1, CreatedBy: int(v.ReviewBy), DateCreated: time.Now(), ModifiedBy: int(v.ReviewBy), DateModified: time.Now()}
		if _, err := models.AddItem_reviews(&r); err == nil {
			var resp responses.ItemReviewResponseDTO = responses.ItemReviewResponseDTO{StatusCode: 200, ItemReview: &r, StatusDesc: "Review successfully added"}
			c.Ctx.Output.SetStatus(200)
			c.Data["json"] = resp
		} else {
			logs.Error("Error returned after attempting to add review is ", err.Error())
			var resp responses.ItemReviewResponseDTO = responses.ItemReviewResponseDTO{StatusCode: 301, ItemReview: nil, StatusDesc: "Review addition failed"}
			c.Data["json"] = resp
		}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Item_reviews by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Item_reviews
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Item_reviewsController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_reviewsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Item_reviews
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Item_reviews
// @Failure 403
// @router / [get]
func (c *Item_reviewsController) GetAll() {
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

	l, err := models.GetAllItem_reviews(query, fields, sortby, order, offset, limit)
	if err != nil {
		logs.Error("Error returned after attempting to add review is ", err.Error())
		var resp responses.ItemReviewsResponseDTO = responses.ItemReviewsResponseDTO{StatusCode: 301, ItemsReviews: nil, StatusDesc: "Unable to fetch reviews"}
		c.Data["json"] = resp
	} else {
		var resp responses.ItemReviewsResponseDTO = responses.ItemReviewsResponseDTO{StatusCode: 200, ItemsReviews: &l, StatusDesc: "Reviews fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Item_reviews
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Item_reviews	true		"body for Item_reviews content"
// @Success 200 {object} models.Item_reviews
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Item_reviewsController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Item_reviews{ItemReviewId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateItem_reviewsById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Item_reviews
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Item_reviewsController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteItem_reviews(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
