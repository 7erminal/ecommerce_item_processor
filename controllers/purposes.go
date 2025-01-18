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

// PurposesController operations for Purposes
type PurposesController struct {
	beego.Controller
}

// URLMapping ...
func (c *PurposesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("ChangeVisibility", c.ChangeVisibility)
}

// Post ...
// @Title Post
// @Description create Purposes
// @Param	body		body 	models.Purposes	true		"body for Purposes content"
// @Success 201 {int} models.Purposes
// @Failure 403 body is empty
// @router / [post]
func (c *PurposesController) Post() {
	logs.Info("Data received is ", c.Ctx.Input.Query("PurposeName"))

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
		filePath = "/uploads/purposes/" + fileName // Define your file path
		err = c.SaveToFile("Image", "../images/"+filePath)

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

	v := models.Purposes{Purpose: c.Ctx.Input.Query("PurposeName"), ImagePath: filePath, Description: c.Ctx.Input.Query("Description"), Active: 1, CreatedBy: 1}

	if _, err := models.AddPurposes(&v); err == nil {
		c.Ctx.Output.SetStatus(201)

		var resp = models.PurposeResponseDTO{StatusCode: 200, Purpose: &v, StatusDesc: "Purpose has been added successfully"}

		c.Data["json"] = resp
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Purposes by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Purposes
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PurposesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPurposesById(id)
	if err != nil {
		var resp = models.PurposeResponseDTO{StatusCode: 301, Purpose: nil, StatusDesc: "Error fetching purpose"}
		c.Data["json"] = resp
	} else {
		var resp = models.PurposeResponseDTO{StatusCode: 200, Purpose: v, StatusDesc: "Purpose has been added successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAllPurposesAndTheirItems ...
// @Title Get All Purposes and their items
// @Description get Features
// @Success 200 {object} models.PurposesResponseFDTO
// @Failure 403 is empty
// @router /items [get]
func (c *PurposesController) GetAllPurposesAndTheirItems() {
	logs.Info("Getting all ")
	v, err := models.GetAllPurposesWithTheirItems()

	if err != nil {
		var resp = responses.PurposesResponseFDTO{StatusCode: 301, Purposes: nil, StatusDesc: "Failed to fetch purposes"}
		logs.Error("Error getting purposes", err.Error())
		c.Data["json"] = resp
	} else {
		logs.Info("Data is ", v)
		modified := helperFunc.ConvertParamsToPurposes(v)

		var resp = responses.PurposesResponseFDTO{StatusCode: 200, Purposes: &modified, StatusDesc: "Purposes fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Purposes
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Purposes
// @Failure 403
// @router / [get]
func (c *PurposesController) GetAll() {
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

	l, err := models.GetAllPurposes(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var resp = models.PurposesResponseDTO{StatusCode: 200, Purposes: &l, StatusDesc: "Purposes fetched successfully"}

		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Purposes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Purposes	true		"body for Purposes content"
// @Success 200 {object} models.Purposes
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PurposesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Purposes{PurposeId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePurposesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Purposes
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	requests.VisibilityRequestDTO	true		"body for Purposes content"
// @Success 200 {object} models.Purposes
// @Failure 403 :id is not int
// @router /change-visibility/:id [put]
func (c *PurposesController) ChangeVisibility() {
	logs.Info("Change visibility request received")
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := requests.VisibilityRequestDTO{Id: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if q, err := models.GetPurposesById(id); err == nil {
		t := models.Purposes{PurposeId: id, Visible: v.Visibility, Purpose: q.Purpose, ImagePath: q.ImagePath, Description: q.Description, Active: q.Active, DateCreated: q.DateCreated, DateModified: time.Now(), CreatedBy: 1, ModifiedBy: 1}
		if err := models.UpdatePurposesById(&t); err == nil {
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
// @Description delete the Purposes
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PurposesController) Delete() {
	logs.Info("Request delete purpose")
	idStr := c.Ctx.Input.Param(":id")
	logs.Info("Delete purpose request received ", idStr)
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePurposes(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
