package controllers

import (
	"encoding/json"
	"errors"
	"item_processor/models"
	"item_processor/structs/responses"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Item_imagesController operations for Item_images
type Item_imagesController struct {
	beego.Controller
}

// URLMapping ...
func (c *Item_imagesController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
	c.Mapping("UploadPictures", c.UploadPictures)
}

// Post ...
// @Title Post
// @Description create Item_images
// @Param	body		body 	models.Item_images	true		"body for Item_images content"
// @Success 201 {int} models.Item_images
// @Failure 403 body is empty
// @router / [post]
func (c *Item_imagesController) Post() {
	// var v models.Item_images

	logs.Info("Data received is ", c.Ctx.Input.Query("ItemID"))

	file, header, err := c.GetFile("Image")
	logs.Info("Data received is ", file)

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Error("Failed to get the file ", err)
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Save the uploaded file
	fileName := header.Filename
	logs.Info("File Name Extracted is ", fileName)
	filePath := "/uploads/items/" + fileName // Define your file path
	logs.Info("File Path Extracted is ", filePath)
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

	itemId, errr := strconv.ParseInt(c.Ctx.Input.Query("ItemID"), 0, 64)

	if errr != nil {
		panic(errr)
	}

	logs.Info("Saving ... ", filePath)
	// json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	v := models.Item_images{ItemId: itemId, ImagePath: filePath, IsDefault: 0, CreatedBy: 1, Active: 1, DateCreated: time.Now(), DateModified: time.Now()}

	if _, err := models.AddItem_images(&v); err == nil {
		c.Ctx.Output.SetStatus(200)

		if k, err := models.GetItemsById(itemId); err == nil {
			// i := models.Items{ItemId: itemId, ImagePath: filePath, ItemName: k.ItemName, Category: k.Category, Description: k.Description, ItemPrice: k.ItemPrice, AvailableSizes: k.AvailableSizes, AvailableColors: k.AvailableColors, Quantity: k.Quantity, DateCreated: k.DateCreated, DateModified: k.DateModified, CreatedBy: k.CreatedBy, ModifiedBy: k.ModifiedBy, Active: k.Active}
			k.ImagePath = filePath

			if err := models.UpdateItemsById(k); err != nil {
				logs.Error(err.Error())
				resp := models.ItemResponseDTO{StatusCode: 302, Item: k, StatusDesc: err.Error()}
				c.Data["json"] = resp
			}

			resp := models.ItemImageResponseDTO{StatusCode: 200, ItemImage: &v, StatusDesc: "Images uploaded successfully"}
			c.Data["json"] = resp
		} else {
			resp := models.ItemImageResponseDTO{StatusCode: 301, ItemImage: nil, StatusDesc: err.Error()}
			c.Data["json"] = resp
		}
	} else {
		resp := models.ItemImageResponseDTO{StatusCode: 301, ItemImage: nil, StatusDesc: err.Error()}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// UploadPicture ...
// @Title Upload Picture
// @Description Upload a picture to items
// @Param	Image		formData 	file	true		"Item Image"
// @Success 200 {int} responses.StringResponseFDTO
// @Failure 403 body is empty
// @router /upload-pictures [post]
func (c *Item_imagesController) UploadPictures() {
	// var v models.Item_images
	file, header, err := c.GetFile("Image")
	logs.Info("Data received is ", file)

	contentType := c.Ctx.Input.Header("Content-Type")
	logs.Info("Content-Type of incoming request:", contentType)

	if err != nil {
		// c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to get image file."}
		logs.Error("Failed to get the file ", err)
		c.ServeJSON()
		return
	}
	defer file.Close()

	// Check the file size
	fileInfo := header.Size
	logs.Info("Received file size:", fileInfo)

	// Save the uploaded file
	fileName := filepath.Base(header.Filename)
	logs.Info("File Name Extracted is ", fileName, "Time now is ", time.Now().Format("20060102150405"))
	filePath := "/uploads/items/" + time.Now().Format("20060102150405") + fileName // Define your file path
	logs.Info("File Path Extracted is ", filePath)
	err = c.SaveToFile("Image", "../images"+filePath)
	if err != nil || header.Size < 1 {
		filePath = ""
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		logs.Error("Error saving file", err)
		// c.Data["json"] = map[string]string{"error": "Failed to save the image file."}
		errorMessage := "Error: Failed to save the image file"

		resp := responses.StringResponseFDTO{StatusCode: http.StatusInternalServerError, Value: &errorMessage, StatusDesc: "Internal Server Error"}

		c.Data["json"] = resp
		c.ServeJSON()
		return
	} else {
		resp := responses.StringResponseFDTO{StatusCode: 200, Value: &filePath, StatusDesc: "Images uploaded successfully"}
		c.Data["json"] = resp
	}

	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get Item_images by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Item_images
// @Failure 403 :id is empty
// @router /:id [get]
func (c *Item_imagesController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetItem_imagesByItemId(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var resp = models.ItemImagesResponseDTO2{StatusCode: 200, ItemImages: v, StatusDesc: "Images fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get Item_images
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Item_images
// @Failure 403
// @router / [get]
func (c *Item_imagesController) GetAll() {
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

	l, err := models.GetAllItem_images(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		var resp = models.ItemImagesResponseDTO{StatusCode: 200, ItemImages: &l, StatusDesc: "Images fetched successfully"}
		c.Data["json"] = resp
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the Item_images
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Item_images	true		"body for Item_images content"
// @Success 200 {object} models.Item_images
// @Failure 403 :id is not int
// @router /:id [put]
func (c *Item_imagesController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Item_images{ItemImageId: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateItem_imagesById(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the Item_images
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *Item_imagesController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteItem_images(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
