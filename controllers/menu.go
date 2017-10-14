package controllers

import (
	"encoding/json"
	"genesis/models"

	"github.com/astaxie/beego"
)

// MenuController for Menu
type MenuController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *MenuController) URLMapping() {
	c.Mapping("CreateMenu", c.CreateMenu)
	c.Mapping("GetMenu", c.GetMenu)
	c.Mapping("DeleteMenu", c.DeleteMenu)
	c.Mapping("AddConditionalMenu", c.AddConditionalMenu)
	c.Mapping("DeleteConditionalMenu", c.DeleteConditionalMenu)
}

// CreateMenu CreateMenu
// @Title Post
// @Description create Menu
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 201 {int} models.Menu
// @Failure 403 body is empty
// @router / [post]
func (c *MenuController) CreateMenu() {
	v := map[string]interface{}{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.CreateMenu(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// GetMenu GetMenu
// @Title Get
// @Description get Menu by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Menu
// @Failure 403 :id is empty
// @router / [get]
func (c *MenuController) GetMenu() {
	v, err := models.GetMenu()
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", v)
	}
	c.ServeJSON()
}

// DeleteMenu DeleteMenu
// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /[delete]
func (c *MenuController) DeleteMenu() {
	if data, err := models.DeleteMenu(); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// AddConditionalMenu  AddConditionalMenu
// @Title Update
// @Description update the Menu
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Menu	true		"body for Menu content"
// @Success 200 {object} models.Menu
// @Failure 403 :id is not int
// @router /conditional [post]
func (c *MenuController) AddConditionalMenu() {
	v := map[string]interface{}{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.AddConditionalMenu(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// DeleteConditionalMenu DeleteConditionalMenu
// @Title Delete
// @Description delete the Menu
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /conditional [delete]
func (c *MenuController) DeleteConditionalMenu() {
	if data, err := models.DeleteConditionalMenu(); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
