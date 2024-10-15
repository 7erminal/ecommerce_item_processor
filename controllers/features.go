package controllers

import (
	"encoding/json"
	"errors"
	helperFunc "item_processor/functions"
	"item_processor/models"
	"item_processor/structs/requests"
	"item_processor/structs/responses"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// FeaturesController operations for Features
type FeaturesController struct {
	beego.Controller
}

// URLMapping ...
func (c *FeaturesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("GetAllFeaturesAndTheirItems", c.GetAllFeaturesAndTheirItems)
}

// Post ...
// @Title Post
// @Description create Features
// @Param	body		body 	models.Features	true		"body for Features content"
// @Success 201 {int} models.Features
// @Failure 403 body is empty
// @router / [post]
func (c *FeaturesController) Post() {

	logs.Info("Data received is ", c.Ctx.Input.Query("FeatureName"))

	file, header, err := c.GetFile("Image")

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
		fileName := header.Filename
		filePath = "/uploads/" + fileName // Define your file path
		err = c.SaveToFile("Image", "."+filePath)
		if err != nil {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			logs.Error("Error saving file", err)
			// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
			errorMessage := "Error: Failed to save the image file"

			resp := models.ErrorResponse{StatusCode: http.StatusInternalServerError, Error: errorMessage, StatusDesc: "Internal Server Error"}

			c.Data["json"] = resp
			c.ServeJSON()
			return
		}
	}

	v := models.Features{FeatureName: c.Ctx.Input.Query("FeatureName"), ImagePath: filePath, Description: c.Ctx.Input.Query("Description"), Active: 1, CreatedBy: 1}

	if _, err := models.AddFeatures(&v); err == nil {
		c.Ctx.Output.SetStatus(201)

		var resp = models.FeatureResponseDTO{StatusCode: 200, Feature: &v, StatusDesc: "Feature has been added successfully"}
		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Features by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Features
// @Failure 403 :id is empty
// @router /:id [get]
func (c *FeaturesController) GetOne() {
	logs.Info("Getting one ")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetFeaturesById(id)
	if err != nil {
		var resp = models.FeatureResponseDTO{StatusCode: 301, Feature: nil, StatusDesc: "Error fetching feature"}
		logs.Info("Error fetching feature ", err.Error())
		c.Data["json"] = resp
	} else {
		var resp = models.FeatureResponseDTO{StatusCode: 200, Feature: v, StatusDesc: "Feature fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllFeaturesAndTheirItems ...
// @Title Get All Features and their items
// @Description get Features
// @Success 200 {object} models.FeaturesResponseFDTO
// @Failure 403 is empty
// @router /items [get]
func (c *FeaturesController) GetAllFeaturesAndTheirItems() {
	logs.Info("Getting all ")
	v, err := models.GetAllFeaturesWithTheirItems()

	if err != nil {
		var resp = responses.FeaturesResponseFDTO{StatusCode: 301, Features: nil, StatusDesc: "Failed to fetch features"}
		logs.Error("Error getting features", err.Error())
		c.Data["json"] = resp
	} else {
		logs.Info("Data is ", v)
		modified := helperFunc.ConvertParamsToFeatures(v)

		var resp = responses.FeaturesResponseFDTO{StatusCode: 200, Features: &modified, StatusDesc: "Features fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Features
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Features
// @Failure 403
// @router / [get]
func (c *FeaturesController) GetAll() {
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

	l, err := models.GetAllFeatures(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var resp = models.FeaturesResponseDTO{StatusCode: 200, Features: &l, StatusDesc: "Features fetched successfully"}

		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Features
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Features	true		"body for Features content"
// @Success 200 {object} models.Features
// @Failure 403 :id is not int
// @router /:id [put]
func (c *FeaturesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Features{FeatureId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateFeaturesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Features
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.VisibilityRequestDTO	true		"body for Features content"
// @Success 200 {object} models.Features
// @Failure 403 :id is not int
// @router /change-visibility/:id [put]
func (c *FeaturesController) ChangeVisibility() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.VisibilityRequestDTO{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if q, err := models.GetFeaturesById(id); err == nil {
		t := models.Features{FeatureId: id, Visible: v.Visibility, FeatureName: q.FeatureName, ImagePath: q.ImagePath, Description: q.Description, Active: q.Active, DateCreated: q.DateCreated, DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1}
		if err := models.UpdateFeaturesById(&t); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Features
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *FeaturesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteFeatures(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
