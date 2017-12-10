package controllers

import (
	"encoding/json"
	"genesis/models"
	"log"

	"github.com/astaxie/beego"
)

// MaterialController for Material
type MaterialController struct {
	beego.Controller
}

// URLMapping URLMapping
func (c *MaterialController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("PostFile", c.PostFile)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetMaterialCount", c.GetMaterialCount)
	c.Mapping("GetAllMaterialNewsList", c.GetAllMaterialNewsList)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
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

// Post Post
// @Title Post
// @Description create Material
// @Param	body		body 	models.MaterialArticles	true		"body for Material content"
// @Success 201 {int} models.MaterialArticles
// @Failure 403 body is empty
// @router / [post]
func (c *MaterialController) Post() {
	var v models.MaterialArticles
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if r, err := models.AddNews(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = models.GetReturnData(0, "OK", r)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// PostFile PostFile
// @Title PostFile
// @Description create File
// @Param	body		body 	models.File	true		"body for File content"
// @Success 201 {int} models.File
// @Failure 403 body is empty
// @router /image [post]
func (c *MaterialController) PostFile() {
	f, h, err := c.GetFile("uploadname")
	if err != nil {
		log.Fatal("getfile err ", err)
	}
	defer f.Close()

	filePath := "static/files/" + h.Filename
	c.SaveToFile("uploadname", filePath) // 保存位置在 static/upload, 没有文件夹要先创建
	mediaInfo, err := models.AddMaterial(filePath, "image")
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
		c.ServeJSON()
		return
	}
	c.Data["json"] = models.GetReturnData(0, "OK", mediaInfo)
	c.ServeJSON()
}

// GetOne GetOne
// @Title Get
// @Description get Material by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MaterialArticles
// @Failure 403 :id is empty
// @router /:media_id [get]
func (c *MaterialController) GetOne() {
	mediaID := c.Ctx.Input.Param(":media_id")
	v, err := models.GetMaterialByMediaID(mediaID)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", v)
	}
	c.ServeJSON()
}

// GetMaterialCount GetMaterialCount
// @Title GetMaterialCount
// @Description get Material
// @Success 200 {object} models.ReturnData
// @Failure 403
// @router /count [get]
func (c *MaterialController) GetMaterialCount() {
	l, err := models.GetMaterialcount()
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", l)
	}
	c.ServeJSON()
}

// GetAllMaterialNewsList GetAllMaterialNewsList
// @Title Get All
// @Description get Material
// @Success 200 {object} models.ReturnData
// @Failure 403
// @router / [get]
func (c *MaterialController) GetAllMaterialNewsList() {

	offset, _ := c.GetInt64("offset")
	count, _ := c.GetInt64("count")
	l, err := models.GetAllMaterialNewsList(offset, count)
	if err != nil {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	} else {
		c.Data["json"] = models.GetReturnData(0, "OK", l)
	}
	c.ServeJSON()
}

// Put Put
// @Title Update
// @Description update the Material
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MaterialArticles	true		"body for Material content"
// @Success 200 {object} models.ReturnData
// @Failure 403 :id is not int
// @router /:mediaId [put]
func (c *MaterialController) Put() {
	mediaID := c.Ctx.Input.Param(":mediaId")
	v := models.MaterialUpdate{MediaID: mediaID}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMaterialByID(&v); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}

// Delete Delete
// @Title Delete
// @Description delete the Material
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:mediaId [delete]
func (c *MaterialController) Delete() {
	mediaID := c.Ctx.Input.Param(":mediaId")
	if err := models.DeleteMaterialByMediaID(mediaID); err == nil {
		c.Data["json"] = models.GetReturnData(0, "OK", nil)
	} else {
		c.Data["json"] = models.GetReturnData(-1, err.Error(), nil)
	}
	c.ServeJSON()
}
