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

// Prepare 拦截请求
func (c *SendMessageController) Prepare() {
	token := c.Ctx.Request.Header.Get("Token")
	err := models.CheckSessionByToken(token)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, "Token Timeout", nil)
		c.ServeJSON()
		c.StopRun()
	}
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
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.GetReturnData(0, "OK", mediaInfo)
	c.ServeJSON()
}

// UploadNewsMessage UploadNewsMessage
// @Title Get
// @Description get Articles by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Articles
// @Failure 403 :id is empty
// @router /news/uplaod  [post]
func (c *SendMessageController) UploadNewsMessage() {
	var v models.Articles
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	data, err := models.UploadNewsMessage(&v)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	}
	c.ServeJSON()
}

// PostAllSendNewsMessage PostAllSendNewsMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router /news [post]
func (c *SendMessageController) PostAllSendNewsMessage() {
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
func (c *SendMessageController) PostAllSendTextMessage() {
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
func (c *SendMessageController) PostAllSendVoiceMessage() {
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
func (c *SendMessageController) PostAllSendImageMessage() {
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

// PostAllSendMessage PostAllSendMessage
// @Title Get All
// @Success 200 {object} models.Articles
// @Failure 403
// @router / [post]
func (c *SendMessageController) PostAllSendMessage() {
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
func (c *SendMessageController) PostPreviewMessage() {
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

// CheckAllSendMessage CheckAllSendMessage
// @Title Update
// @Description update the Articles
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Articles	true		"body for Send_message content"
// @Success 200 {object} models.Articles
// @Failure 403 :id is not int
// @router /:msgID/status [get]
func (c *SendMessageController) CheckAllSendMessage() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	if data, err := models.CheckAllSendMessage(msgID); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// DeleteAllSendMessage DeleteAllSendMessage
// @Title Delete
// @Description delete the Articles
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:msgID/:articleIDX [delete]
func (c *SendMessageController) DeleteAllSendMessage() {
	msgID, _ := strconv.ParseInt(c.Ctx.Input.Param(":msgID"), 0, 64)
	articleIDX, _ := strconv.ParseInt(c.Ctx.Input.Param(":articleIDX"), 0, 64)
	if data, err := models.DeleteAllSendMessage(msgID, articleIDX); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", data)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
