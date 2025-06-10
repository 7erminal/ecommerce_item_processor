package controllers

import (
	"encoding/json"
	"errors"
	"item_processor/models"
	"item_processor/structs/requests"
	"item_processor/structs/responses"
	"net/http"
	"path/filepath"
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
	c.Mapping("GetOneWithReference", c.GetOneWithReference)
}

// Post ...
// @Title Post
// @Description create Item_reviews
// @Param	body		body 	requests.AddReviewRequest	true		"body for Item_reviews content"
// @Param	ImagePath	query	string	false	"Image path"
// @Param	ReviewBy	query	string	false	"ID of the user that is doing the review"
// @Param	ItemId	query	string	false	"ID of the item"
// @Param	Rating	query	string	false	"Rating"
// @Param	Reference	query	string	false	"Reference. Could be an ID"
// @Param	Review	query	string	false	"Review text"
// @Success 201 {int} responses.ItemReviewResponseDTO
// @Failure 403 body is empty
// @router / [post]
func (c *Item_reviewsController) Post() {
	var v requests.AddReviewRequest
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)

	file, header, err := c.GetFile("ImagePath")
	var filePath string = ""

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Info("Failed to get the file ", err)
		// c.ServeJSON()
		// return
	} else {
		defer file.Close()

		// Save the uploaded file
		fileName := filepath.Base(header.Filename)
		filePath = "/uploads/reviews/" + time.Now().Format("20060102150405") + fileName // Define your file path
		err = c.SaveToFile("ImagePath", "../images/"+filePath)
		if err != nil || header.Size < 1 {
			filePath = ""
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := responses.ItemReviewResponseDTO{StatusCode: 601, ItemReview: nil, StatusDesc: "Error updating review. " + errorMessage}

			c.Data["json"] = resp
		}
	}

	reviewByStr := c.Ctx.Input.Query("ReviewBy")
	reviewBy, err := strconv.ParseInt(reviewByStr, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid ReviewBy parameter"}
		c.ServeJSON()
		return
	}

	itemStr := c.Ctx.Input.Query("ItemId")
	item_, err := strconv.ParseInt(itemStr, 10, 64)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid ReviewBy parameter"}
		c.ServeJSON()
		return
	}

	if u, err := models.GetUsersById(reviewBy); err == nil {
		item := models.Items{}
		if item_, err := models.GetItemsById(item_); err == nil {
			// logs.Error("Error returned after attempting to add review is ", err.Error())
			item = *item_
		}

		ratingStr := c.Ctx.Input.Query("Rating")
		rating, err := strconv.ParseFloat(ratingStr, 64)
		if err != nil {
			logs.Error("An error occurred converting ", ratingStr, " to float ", err.Error())
		}

		referenceStr := c.Ctx.Input.Query("Reference")
		reference, err := strconv.ParseInt(referenceStr, 10, 64)
		if err != nil {
			logs.Error("An error occurred converting ", referenceStr, " to int ", err.Error())
		}
		var r models.Item_reviews = models.Item_reviews{Review: c.Ctx.Input.Query("Review"), Item: &item, Reference: reference, Rating: rating, ReviewBy: u, Active: 1, CreatedBy: int(reviewBy), DateCreated: time.Now(), ModifiedBy: int(reviewBy), DateModified: time.Now()}
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

// GetOneByReference ...
// @Title Get One By Reference
// @Description get Item_reviews by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Item_reviews
// @Failure 403 :id is empty
// @router /get-review-with-reference/:id [get]
func (c *Item_reviewsController) GetOneWithReference() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_reviewsById(id)
	if err != nil {
		c.Data["json"] = err.Error()
		var resp responses.ItemReviewResponseDTO = responses.ItemReviewResponseDTO{StatusCode: 301, ItemReview: nil, StatusDesc: "Fetching review failed"}
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resp
	} else {
		c.Data["json"] = v
		var resp responses.ItemReviewResponseDTO = responses.ItemReviewResponseDTO{StatusCode: 200, ItemReview: v, StatusDesc: "Review successfully added"}
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = resp
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
		// logs.Info("Reviews fetched successfully")
		logs.Info("Reviews fetched successfully ", l)
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
