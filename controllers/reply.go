package controllers

import (
	"encoding/json"
	"errors"
	"genesis/models"
	"strings"

	"github.com/astaxie/beego"
)

// oprations for Reply
type ReplyController struct {
	beego.Controller
}

func (c *ReplyController) URLMapping() {
	c.Mapping("PostDefult", c.PostDefult)
	c.Mapping("PostKey", c.PostKey)
	c.Mapping("GetOneDefult", c.GetOneDefult)
	c.Mapping("GetOneKey", c.GetOneKey)
	c.Mapping("GetAllKey", c.GetAllKey)
	c.Mapping("PutDefult", c.PutDefult)
	c.Mapping("PutKey", c.PutKey)
	c.Mapping("DeleteDefult", c.DeleteDefult)
	c.Mapping("DeleteKey", c.DeleteKey)
}

// @Title Post
// @Description create Reply
// @Param	body		body 	models.Reply	true		"body for Reply content"
// @Success 201 {int} models.Reply
// @Failure 403 body is empty
// @router /defult [post]
func (c *ReplyController) PostDefult() {
	var v models.ReplyDefult
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddReplyDefult(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Post
// @Description create Reply
// @Param	body		body 	models.Reply	true		"body for Reply content"
// @Success 201 {int} models.Reply
// @Failure 403 body is empty
// @router /key [post]
func (c *ReplyController) PostKey() {
	var v models.ReplyKey
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if _, err := models.AddReplyKey(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = v
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Reply by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Reply
// @Failure 403 :id is empty
// @router /defult/:type [get]
func (c *ReplyController) GetOneDefult() {
	_type := c.Ctx.Input.Param(":type")
	v, err := models.GetReplyDefultByType(_type)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get
// @Description get Reply by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Reply
// @Failure 403 :id is empty
// @router /key/:key [get]
func (c *ReplyController) GetOneKey() {
	key := c.Ctx.Input.Param(":key")
	v, err := models.GetReplyKeyByKey(key)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Reply
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Reply
// @Failure 403
// @router /key [get]
func (c *ReplyController) GetAllKey() {
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

	l, err := models.GetAllReplyKey(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Reply
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Reply	true		"body for Reply content"
// @Success 200 {object} models.Reply
// @Failure 403 :id is not int
// @router /defult/:type [put]
func (c *ReplyController) PutDefult() {
	_type := c.Ctx.Input.Param(":type")
	v := models.ReplyDefult{Type: _type}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateReplyDefultByType(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Reply
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Reply	true		"body for Reply content"
// @Success 200 {object} models.Reply
// @Failure 403 :id is not int
// @router /key/:key [put]
func (c *ReplyController) PutKey() {
	key := c.Ctx.Input.Param(":key")
	v := models.ReplyKey{Key: key}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateReplyKeyByKey(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Reply
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /defult/:type [delete]
func (c *ReplyController) DeleteDefult() {
	_type := c.Ctx.Input.Param(":type")
	if err := models.DeleteReplyDefultByType(_type); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Reply
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /key/:id [delete]
func (c *ReplyController) DeleteKey() {
	key := c.Ctx.Input.Param(":key")
	if err := models.DeleteReplyKeyByKey(key); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
