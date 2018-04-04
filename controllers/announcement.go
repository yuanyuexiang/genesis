package controllers

import (
	"encoding/json"
	"genesis/models"
	"strconv"

	"github.com/astaxie/beego"
)

// AnnouncementController for Send_message
type AnnouncementController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *AnnouncementController) URLMapping() {
	c.Mapping("PostAllSendNewsMessage", c.PostAllSendNewsMessage)
	c.Mapping("PostAllSendTextMessage", c.PostAllSendTextMessage)
	c.Mapping("PostAllSendVoiceMessage", c.PostAllSendVoiceMessage)
	c.Mapping("PostAllSendImageMessage", c.PostAllSendImageMessage)
	c.Mapping("PostAllAnnouncement", c.PostAllAnnouncement)
	c.Mapping("CheckAllAnnouncement", c.CheckAllAnnouncement)
	c.Mapping("DeleteAllAnnouncement", c.DeleteAllAnnouncement)
}

// Prepare 拦截请求
func (c *AnnouncementController) Prepare() {
	token := c.Ctx.Request.Header.Get("Token")
	err := models.CheckSessionByToken(token)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		c.ServeJSON()
		c.StopRun()
	}
}

// PostAllSendNewsMessage PostAllSendNewsMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /news [post]
func (c *AnnouncementController) PostAllSendNewsMessage() {
	var v models.AllSendNewsMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostAllSendTextMessage PostAllSendTextMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /text [post]
func (c *AnnouncementController) PostAllSendTextMessage() {
	var v models.AllSendTextMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostAllSendVoiceMessage PostAllSendVoiceMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /voice [post]
func (c *AnnouncementController) PostAllSendVoiceMessage() {
	var v models.AllSendVoiceMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostAllSendImageMessage PostAllSendImageMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /image [post]
func (c *AnnouncementController) PostAllSendImageMessage() {
	var v models.AllSendImageMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostAllAnnouncement PostAllAnnouncement
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router / [post]
func (c *AnnouncementController) PostAllAnnouncement() {
	v := map[string]interface{}{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostPreviewMessage PostPreviewMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /preview [post]
func (c *AnnouncementController) PostPreviewMessage() {
	v := map[string]interface{}{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostPreviewMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// CheckAllAnnouncement CheckAllAnnouncement
// @Title Update
// @Description update the Articles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Articles	true		"body for Send_message content"
// @Success 200 {object} models.Articles
// @Failure 403 :id is not int
// @router /:msgID/status [get]
func (c *AnnouncementController) CheckAllAnnouncement() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	if data, err := models.CheckAllSendMessage(msgID); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// DeleteAllAnnouncement DeleteAllAnnouncement
// @Title Delete
// @Description delete the Articles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:msgID/:articleIDX [delete]
func (c *AnnouncementController) DeleteAllAnnouncement() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	articleIDX, _ := strconv.ParseInt(c.Ctx.Input.Param(":articleIDX"), 0, 64)
	if data, err := models.DeleteAllSendMessage(msgID, articleIDX); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
