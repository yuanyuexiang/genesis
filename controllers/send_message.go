package controllers

import (
	"encoding/json"
	"genesis/models"
	"log"
	"strconv"

	"github.com/astaxie/beego"
)

// SendMessageController for Send_message
type SendMessageController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *SendMessageController) URLMapping() {
	c.Mapping("UploadNewsMessageImage", c.UploadNewsMessageImage)
	c.Mapping("UploadNewsMessage", c.UploadNewsMessage)
	c.Mapping("PostAllSendNewsMessage", c.PostAllSendNewsMessage)
	c.Mapping("PostAllSendTextMessage", c.PostAllSendTextMessage)
	c.Mapping("PostAllSendVoiceMessage", c.PostAllSendVoiceMessage)
	c.Mapping("PostAllSendImageMessage", c.PostAllSendImageMessage)
	c.Mapping("PostAllSendMessage", c.PostAllSendMessage)
	c.Mapping("CheckAllSendMessage", c.CheckAllSendMessage)
	c.Mapping("DeleteAllSendMessage", c.DeleteAllSendMessage)
}

// UploadNewsMessageImage  UploadNewsMessageImage
// @router /image/uplaod [post]
func (c *SendMessageController) UploadNewsMessageImage() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	filePath := "static/files/" + h.Filename
	c.SaveToFile("uploadname", filePath) // 保存位置在 static/upload, 没有文件夹要先创建
	mediaInfo, err := models.UploadNewsMessageImage(filePath)
	if err != nil {
		return
	}
	//returnData := map[string]string{"filePath": filePath + "---" + mediaInfo.Introduction}
	c.Data["json"] = mediaInfo
	c.ServeJSON()
}

// UploadNewsMessage UploadNewsMessage
// @Title Get
// @Description get Send_message by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Send_message
// @Failure 403 :id is empty
// @router /news/uplaod  [post]
func (c *SendMessageController) UploadNewsMessage() {
	var v models.Articles
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.UploadNewsMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// PostAllSendNewsMessage PostAllSendNewsMessage
// @Title Get All
// @Success 200 {object} models.Send_message
// @Failure 403
// @router /news [post]
func (c *SendMessageController) PostAllSendNewsMessage() {
	var v models.AllSendNewsMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// PostAllSendTextMessage PostAllSendTextMessage
// @Title Get All
// @Success 200 {object} models.Send_message
// @Failure 403
// @router /text [post]
func (c *SendMessageController) PostAllSendTextMessage() {
	var v models.AllSendTextMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// PostAllSendVoiceMessage PostAllSendVoiceMessage
// @Title Get All
// @Success 200 {object} models.Send_message
// @Failure 403
// @router /voice [post]
func (c *SendMessageController) PostAllSendVoiceMessage() {
	var v models.AllSendVoiceMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// PostAllSendImageMessage PostAllSendImageMessage
// @Title Get All
// @Success 200 {object} models.Send_message
// @Failure 403
// @router /image [post]
func (c *SendMessageController) PostAllSendImageMessage() {
	var v models.AllSendImageMessage
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// PostAllSendMessage PostAllSendMessage
// @Title Get All
// @Success 200 {object} models.Send_message
// @Failure 403
// @router / [post]
func (c *SendMessageController) PostAllSendMessage() {
	v := map[string]interface{}{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.PostAllSendMessage(&v)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = data
	}
	c.ServeJSON()
}

// CheckAllSendMessage CheckAllSendMessage
// @Title Update
// @Description update the Send_message
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Send_message	true		"body for Send_message content"
// @Success 200 {object} models.Send_message
// @Failure 403 :id is not int
// @router /:msgID/status [get]
func (c *SendMessageController) CheckAllSendMessage() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	if data, err := models.CheckAllSendMessage(msgID); err == nil {
		c.Data["json"] = data
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// DeleteAllSendMessage DeleteAllSendMessage
// @Title Delete
// @Description delete the Send_message
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:msgID/:articleIDX [delete]
func (c *SendMessageController) DeleteAllSendMessage() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	articleIDX, _ := strconv.ParseInt(c.Ctx.Input.Param(":articleIDX"), 0, 64)
	if data, err := models.DeleteAllSendMessage(msgID, articleIDX); err == nil {
		c.Data["json"] = data
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
