package controllers

import (
	"encoding/json"
	"genesis/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// MaterialController for Material
type MaterialController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *MaterialController) URLMapping() {
	c.Mapping("PostNews", c.PostNews)
	c.Mapping("GetOneNews", c.GetOneNews)
	c.Mapping("GetAllMaterialNews", c.GetAllMaterialNews)
	c.Mapping("DeleteOneNews", c.DeleteOneNews)

	c.Mapping("PostMedia", c.PostMedia)
	c.Mapping("GetOneMedia", c.GetOneMedia)
	c.Mapping("GetAllMaterialMedia", c.GetAllMaterialMedia)
	c.Mapping("DeleteOneMedia", c.DeleteOneMedia)
}

// Prepare 拦截请求
func (c *MaterialController) Prepare() {
	token := c.Ctx.Request.Header.Get("Token")
	err := models.CheckSessionByToken(token)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}
}

// PostNews PostNews
// @Title Post
// @Description create Material
// @Param	body		body 	models.MaterialArticles	true		"body for Material content"
// @Success 201 {int} models.MaterialArticles
// @Failure 403 body is empty
// @router /news [post]
func (c *MaterialController) PostNews() {
	var v models.MaterialNews
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if r, err := models.AddMaterialNews(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.GetReturnData(0, "OK", r)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// GetAllMaterialNews GetAllMaterialNews
// @Title Get All
// @Description get Material
// @Success 200 {object} models.ReturnData
// @Failure 403
// @router /news [get]
func (c *MaterialController) GetAllMaterialNews() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

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
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = models.GetReturnData(-1, "Error: invalid query key/value pair", nil)
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMaterialNews(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", l)
	}
	c.ServeJSON()
}

// GetOneNews GetOneNews
// @Title Get
// @Description get Material by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MaterialArticles
// @Failure 403 :id is empty
// @router /news/:id [get]
func (c *MaterialController) GetOneNews() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMaterialNewsByID(id)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", v)
	}
	c.ServeJSON()
}

// DeleteOneNews DeleteOneNews
// @Title Delete
// @Description delete the Material
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /news/:id [delete]
func (c *MaterialController) DeleteOneNews() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteMaterialNewsByID(id); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// PostMedia PostMedia
// @Title Post
// @Description create Material
// @Param	body		body 	models.MaterialArticles	true		"body for Material content"
// @Success 201 {int} models.MaterialArticles
// @Failure 403 body is empty
// @router /media [post]
func (c *MaterialController) PostMedia() {
	var v models.MaterialMedia
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if r, err := models.AddMaterialMedia(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.GetReturnData(0, "OK", r)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// GetOneMedia GetOneMedia
// @Title Get
// @Description get Material by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MaterialArticles
// @Failure 403 :id is empty
// @router /media/:id [get]
func (c *MaterialController) GetOneMedia() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetMaterialMediaByID(id)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", v)
	}
	c.ServeJSON()
}

// GetAllMaterialMedia GetAllMaterialMedia
// @Title Get All
// @Description get Material
// @Success 200 {object} models.ReturnData
// @Failure 403
// @router /media [get]
func (c *MaterialController) GetAllMaterialMedia() {
	var fields []string
	var sortby []string
	var order []string
	var query map[string]string = make(map[string]string)
	var limit int64 = 10
	var offset int64 = 0

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
			kv := strings.Split(cond, ":")
			if len(kv) != 2 {
				c.Data["json"] = models.GetReturnData(-1, "Error: invalid query key/value pair", nil)
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllMaterialMedia(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", l)
	}
	c.ServeJSON()
}

// DeleteOneMedia DeleteOneMedia
// @Title Delete
// @Description delete the Material
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /media/:id [delete]
func (c *MaterialController) DeleteOneMedia() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteMaterialMediaByID(id); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
