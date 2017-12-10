package controllers

import (
	"encoding/json"
	"errors"
	"genesis/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// PermissionController for Permission
type PermissionController struct {
	beego.Controller
}

// Prepare 拦截请求
func (c *PermissionController) Prepare() {
	token := c.Ctx.Request.Header.Get("Token")
	session, errGetSessionByToken := models.GetSessionByToken(token)
	if errGetSessionByToken != nil {
		c.Data["json"] = models.GetReturnData(-1, errGetSessionByToken.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}
	administrator, errGetAdministratorByID := models.GetAdministratorByID(session.ID)
	if errGetAdministratorByID != nil {
		c.Data["json"] = models.GetReturnData(-1, errGetAdministratorByID.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}
	role, action, resource := administrator.Role, c.Ctx.Request.Method, c.Ctx.Request.RequestURI
	errCheckPermission := models.CheckPermission(&models.Permission{Role: role, Action: action, Resource: resource})
	if errCheckPermission != nil {
		c.Data["json"] = models.GetReturnData(-1, errCheckPermission.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}
}

// URLMapping URLMapping
func (c *PermissionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post Post
// @Title Post
// @Description create Permission
// @Param	body		body 	models.Permission	true		"body for Permission content"
// @Success 201 {int} models.Permission
// @Failure 403 body is empty
// @router / [post]
func (c *PermissionController) Post() {
	var v models.Permission
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddPermission(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetOne GetOne
// @Title Get
// @Description get Permission by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Permission
// @Failure 403 :id is empty
// @router /:id [get]
func (c *PermissionController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetPermissionByID(id)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// GetAll GetAll
// @Title Get All
// @Description get Permission
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Permission
// @Failure 403
// @router / [get]
func (c *PermissionController) GetAll() {
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
				c.Data["json"] = errors.New("Error: invalid query key/value pair")
				c.ServeJSON()
				return
			}
			k, v := kv[0], kv[1]
			query[k] = v
		}
	}

	l, err := models.GetAllPermission(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// Put Put
// @Title Update
// @Description update the Permission
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Permission	true		"body for Permission content"
// @Success 200 {object} models.Permission
// @Failure 403 :id is not int
// @router /:id [put]
func (c *PermissionController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Permission{ID: id}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdatePermissionByID(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// Delete Delete
// @Title Delete
// @Description delete the Permission
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *PermissionController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeletePermission(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
