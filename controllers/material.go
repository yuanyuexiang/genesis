package controllers

import (
	"encoding/json"
	"genesis/models"
	"log"

	"github.com/astaxie/beego"
)

// oprations for Material
type MaterialController struct {
	beego.Controller
}

func (c *MaterialController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("PostFile", c.PostFile)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetMaterialCount", c.GetMaterialCount)
	c.Mapping("GetAllMaterialNewsList", c.GetAllMaterialNewsList)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// @Title Post
// @Description create Material
// @Param	body		body 	models.Material	true		"body for Material content"
// @Success 201 {int} models.Material
// @Failure 403 body is empty
// @router / [post]
func (c *MaterialController) Post() {
	var v models.MaterialArticles
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if r, err := models.AddNews(&v); err == nil {
		c.Ctx.Output.SetStatus(201)
		c.Data["json"] = r
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

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
		return
	}
	//returnData := map[string]string{"filePath": filePath + "---" + mediaInfo.Introduction}
	c.Data["json"] = mediaInfo
	c.ServeJSON()
}

// @Title Get
// @Description get Material by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Material
// @Failure 403 :id is empty
// @router /:media_id [get]
func (c *MaterialController) GetOne() {
	mediaID := c.Ctx.Input.Param(":media_id")
	v, err := models.GetMaterialByMediaID(mediaID)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

// @Title GetMaterialCount
// @Description get Material
// @Success 200 {object} models.Material
// @Failure 403
// @router /count [get]
func (c *MaterialController) GetMaterialCount() {
	l, err := models.GetMaterialcount()
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Get All
// @Description get Material
// @Success 200 {object} models.Material
// @Failure 403
// @router / [get]
func (c *MaterialController) GetAllMaterialNewsList() {

	offset, _ := c.GetInt64("offset")
	count, _ := c.GetInt64("count")
	l, err := models.GetAllMaterialNewsList(offset, count)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

// @Title Update
// @Description update the Material
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Material	true		"body for Material content"
// @Success 200 {object} models.Material
// @Failure 403 :id is not int
// @router /:mediaId [put]
func (c *MaterialController) Put() {
	mediaID := c.Ctx.Input.Param(":mediaId")
	v := models.MaterialUpdate{MediaID: mediaID}
	json.Unmarshal(c.Ctx.Input.RequestBody, &v)
	if err := models.UpdateMaterialByID(&v); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// @Title Delete
// @Description delete the Material
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:mediaId [delete]
func (c *MaterialController) Delete() {
	mediaID := c.Ctx.Input.Param(":mediaId")
	if err := models.DeleteMaterialByMediaID(mediaID); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
