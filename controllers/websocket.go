package controllers

import (
	"fmt"
	"genesis/models"
	"log"

	"github.com/astaxie/beego"
)

// WebsocketController for Dashboard
type WebsocketController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *WebsocketController) URLMapping() {
	c.Mapping("Get", c.Get)
}

// Get Get
// @Title Get
// @Description get Dashboard by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Dashboard
// @Failure 403 :id is empty
// @router / [get]
func (c *WebsocketController) Get() {
	if v, err := c.GetInt64("userID"); err == nil {
		fmt.Println("=============================================")
		fmt.Println(v)
	}
	err := models.Upgrade(c.Ctx.ResponseWriter, c.Ctx.Request, nil)
	if err != nil {
		log.Fatal(err)
	}
}
