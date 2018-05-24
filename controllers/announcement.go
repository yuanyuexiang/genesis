package controllers

import (
	"encoding/json"
	"genesis/models"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
)

// AnnouncementController for Send_message
type AnnouncementController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *AnnouncementController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("PutStatus", c.PutStatus)
	c.Mapping("Delete", c.Delete)
	/*
		c.Mapping("PostAllSendNewsMessage", c.PostAllSendNewsMessage)
		c.Mapping("PostAllSendTextMessage", c.PostAllSendTextMessage)
		c.Mapping("PostAllSendVoiceMessage", c.PostAllSendVoiceMessage)
		c.Mapping("PostAllSendImageMessage", c.PostAllSendImageMessage)
		c.Mapping("PostAllAnnouncement", c.PostAllAnnouncement)
		c.Mapping("CheckAllAnnouncement", c.CheckAllAnnouncement)
		c.Mapping("DeleteAllAnnouncement", c.DeleteAllAnnouncement)
	*/
}

/*
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
*/
// Post Post
// @Title Post
// @Description create Announcement
// @Param	body		body 	models.Announcement	true		"body for Announcement content"
// @Success 201 {int} models.Announcement
// @Failure 403 body is empty
// @router / [post]
func (c *AnnouncementController) Post() {
	var v models.Announcement
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddAnnouncement(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = models.GetReturnData(0, "OK", v)
		} else {
			c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		}
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// GetOne GetOne
// @Title Get
// @Description get Announcement by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Announcement
// @Failure 403 :id is empty
// @router /:id [get]
func (c *AnnouncementController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v, err := models.GetAnnouncementByID(id)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", v)
	}
	c.ServeJSON()
}

// PutStatus PutStatus
// @Title Update
// @Description update the Article
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Article	true		"body for Article content"
// @Success 200 {object} models.Article
// @Failure 403 :id is not int
// @router /:id/status [put]
func (c *AnnouncementController) PutStatus() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	v := models.Announcement{ID: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateAnnouncementStatusByID(&v); err == nil {
			c.Data["json"] = models.GetReturnData(0, "OK", nil)
		} else {
			c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		}
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// GetAll GetAll
// @Title Get All
// @Description get Announcement
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Announcement
// @Failure 403
// @router / [get]
func (c *AnnouncementController) GetAll() {
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

	l, err := models.GetAllAnnouncement(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", l)
	}
	c.ServeJSON()
}

// Delete Delete
// @Title Delete
// @Description delete the Announcement
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *AnnouncementController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.ParseInt(idStr, 0, 64)
	if err := models.DeleteAnnouncement(id); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

/*
// PostAllSendNewsMessage PostAllSendNewsMessage
// @Title Get All
// @Success 200 {object} models.Announcements
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
// @Success 200 {object} models.Announcements
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
// @Success 200 {object} models.Announcements
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
// @Success 200 {object} models.Announcements
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
// @Success 200 {object} models.Announcements
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
// @Success 200 {object} models.Announcements
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
// @Description update the Announcements
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Announcements	true		"body for Send_message content"
// @Success 200 {object} models.Announcements
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
// @Description delete the Announcements
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
*/
