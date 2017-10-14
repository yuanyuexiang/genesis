package controllers

import (
	"encoding/json"
	"genesis/models"

	"github.com/astaxie/beego"
)

// SessionController for Session
type SessionController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *SessionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("Delete", c.Delete)
}

// Post Post
// @Title Post
// @Description create Session
// @Param	body		body 	models.Session	true		"body for Session content"
// @Success 201 {int} models.Session
// @Failure 403 body is empty
// @router / [post]
func (c *SessionController) Post() {
	var v models.AuthInfo
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if m, err := models.CreateSession(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.GetReturnData(0, "OK", m)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// Delete Delete
// @Title Delete
// @Description delete the Session
// @Param	token		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 token is empty
// @router / [delete]
func (c *SessionController) Delete() {
	token := c.Ctx.Request.Header.Get("Token")
	if token == "" {
		c.Data["json"] = models.GetReturnData(-1, "NO Token", nil)
		c.ServeJSON()
		return
	}
	if err := models.DeleteSessionByToken(token); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
